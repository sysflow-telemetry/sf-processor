package windows

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.ibm.com/sysflow/goutils/logger"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/winlogbeat/checkpoint"
	"github.com/elastic/beats/v7/winlogbeat/eventlog"
	"github.ibm.com/sysflow/sf-processor/core/flattener"
	"github.ibm.com/sysflow/sf-processor/driver/driver"
	"github.ibm.com/sysflow/sf-processor/driver/pipeline"
	"github.ibm.com/sysflow/sf-processor/driver/windows/sysmon"
)

const (
	cName        = "name"
	cAPI         = "api"
	cChannelSize = 100000
)

// WinEvtDriver represents a Windows Event Data Source
type WinEvtDriver struct {
	pipeline *pipeline.Pipeline
}

// NewWinEvtDriver creates a new windows event driver
func NewWinEvtDriver() driver.Driver {
	return &WinEvtDriver{}
}

// Init initializes the driver.
func (w *WinEvtDriver) Init(pipeline *pipeline.Pipeline) error {
	w.pipeline = pipeline
	w.pipeline.AddChannel("flattenerchan", new(flattener.EFRChannel))
	return nil
}

// Run processes windows event logs and creates and exports them into the pipeline
func (w *WinEvtDriver) Run(path string, running *bool) error {
	conf := make(map[string]interface{})
	channel := w.pipeline.GetRootChannel()
	efrChannel := channel.(*flattener.EFRChannel)
	sm := sysmon.NewSMProcessor(efrChannel)
	conf[cAPI] = ""
	conf[cName] = sm.GetProvider()
	cfg, err := common.NewConfigFrom(conf)
	if err != nil {
		logger.Error.Println("winevtlog provider config error:", err)
		return err
	}

	eventLog, err := eventlog.New(cfg)
	if err != nil {
		logger.Error.Println("Failed to create new event log error:", err)
		return err
	}
	cp := checkpoint.EventLogState{}
	err = eventLog.Open(cp)
	if err != nil {
		fmt.Printf("failed to open windows event log: %v", err)
		return err
	}
	logger.Trace.Printf("Windows Event Log '%s' opened successfully", eventLog.Name())
	// setup closing the API if either the run function is signaled asynchronously
	// to shut down or when returning after io.EOF
	/*cancelCtx, cancelFn := ctxtool.WithFunc(ctxtool.FromCanceller(ctx.Cancelation), func() {
		if err := api.Close(); err != nil {
			log.Errorf("Error while closing Windows Eventlog Access: %v", err)
		}
	})
	defer cancelFn()
	*/
	// read loop
	for *running {
		records, err := eventLog.Read()
		switch err {
		case nil:
			break
		case io.EOF:
			logger.Trace.Printf("End of Winlog event stream reached: %v", err)
			return nil
		default:
			// only log error if we are not shutting down
			/*	if cancelCtx.Err() != nil {
				return nil
			}*/

			logger.Error.Printf("Error occured while reading from Windows Event Log '%v': %v", eventLog.Name(), err)
			return errors.New("Unknown Error occurred while looping through windows evt")
		}

		if len(records) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		sm.Process(records)
	}
	close(efrChannel.In)
	w.pipeline.Wait()
	return nil
}
