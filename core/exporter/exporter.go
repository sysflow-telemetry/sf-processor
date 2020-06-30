package exporter

import (
	"crypto/tls"
	"fmt"
	"os"
	"sync"

	syslog "github.com/RackSec/srslog"
	sp "github.com/sysflow-telemetry/sf-apis/go/processors"
	"github.ibm.com/sysflow/goutils/logger"
	"github.ibm.com/sysflow/sf-processor/core/policyengine/engine"
)

// Exporter defines a syslogger plugin.
type Exporter struct {
	recs    []*engine.Record
	counter int
	sysl    *syslog.Writer
	config  Config
}

// NewExporter creates a new plugin instance.
func NewExporter() sp.SFProcessor {
	return new(Exporter)
}

// Init initializes the plugin with a configuration map and cache.
func (s *Exporter) Init(conf map[string]string) error {
	var err error
	s.config = CreateConfig(conf)
	if s.config.Export == FileExport {
		os.Remove(s.config.Path)
	} else if s.config.Export == SyslogExport {
		raddr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
		if s.config.Proto == TCPTLSProto {
			// TODO: verify connection with given trust certifications
			nopTLSConfig := &tls.Config{InsecureSkipVerify: true}
			s.sysl, err = syslog.DialWithTLSConfig("tcp+tls", raddr, syslog.LOG_ALERT|syslog.LOG_DAEMON, s.config.Tag, nopTLSConfig)
		} else {
			s.sysl, err = syslog.Dial(s.config.Proto.String(), raddr, syslog.LOG_ALERT|syslog.LOG_DAEMON, s.config.Tag)
		}
	}
	return err
}

// Process implements the main interface of the plugin.
func (s *Exporter) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*engine.RecordChannel)
	record := cha.In
	defer wg.Done()
	logger.Trace.Printf("Starting Exporter in mode %s with channel capacity %d", s.config.Export.String(), cap(record))
	for {
		fc, ok := <-record
		if !ok {
			s.process()
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		s.counter++
		s.recs = append(s.recs, fc)
		if s.counter > s.config.EventBuffer {
			s.process()
			s.recs = make([]*engine.Record, 0)
			s.counter = 0
		}
	}
	logger.Trace.Println("Exiting Syslogger")
}

func (s *Exporter) process() {
	s.export(s.createEvents())
}

func (s *Exporter) createEvents() []Event {
	if s.config.ExpType == AlertType {
		return CreateOffenses(s.recs, s.config)
	}
	return CreateTelemetryRecords(s.recs, s.config)
}

func (s *Exporter) export(events []Event) {
	if s.config.Format == JSONFormat {
		s.exportAsJSON(events)
	}
}

func (s *Exporter) exportAsJSON(events []Event) {
	for _, e := range events {
		if s.config.Export == StdOutExport {
			fmt.Println(e.ToJSONStr())
		} else if s.config.Export == SyslogExport {
			s.sysl.Alert(e.ToJSONStr())
		} else if s.config.Export == FileExport {
			f, err := os.OpenFile(s.config.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logger.Error.Println("Can't open trace file:\n", err)
			}
			defer f.Close()
			if _, err := f.WriteString(e.ToJSONStr() + "\n"); err != nil {
				logger.Error.Println("Can't write to trace file:\n", err)
			}
		}
	}
}

// SetOutChan sets the output channel of the plugin.
func (s *Exporter) SetOutChan(ch interface{}) {}

// Cleanup tears down plugin resources.
func (s *Exporter) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
