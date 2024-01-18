package processor

import (
	"fmt"
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/source/otel"
)

const (
	pluginLabel string = "otelexporter"
)

var Plugin OTELExporter

type OTELExporter struct{}

func NewOTELExporter() plugins.SFProcessor {
	return new(OTELExporter)
}

func (s *OTELExporter) GetName() string {
	return pluginLabel
}

func (s *OTELExporter) Init(conf map[string]interface{}) error {
	return nil
}

func (s *OTELExporter) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginLabel, NewOTELExporter)
}

func (s *OTELExporter) Process(ch []interface{}, wg *sync.WaitGroup) {
	for _, chi := range ch {
		cha := chi.(*plugins.Channel[*otel.ResourceLogs])
		record := cha.In
		defer wg.Done()

		for {
			fc, ok := <-record
			if !ok {
				logger.Trace.Println("Channel closed shutting down")
				break
			}
			fmt.Printf("Dealing with a record--%s\n", fc)
		}
	}
	logger.Trace.Println("Exiting otel exporter")
}

func (s *OTELExporter) SetOutChan(ch []interface{}) {}

func (s *OTELExporter) Cleanup() {}

func main() {}
