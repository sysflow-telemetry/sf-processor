package main

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.ibm.com/sysflow/goutils/logger"
)

const (
	pluginName string = "example"
)

// Plugin exports a symbol for this plugin.
var Plugin Example

// Example defines an example plugin.
type Example struct{}

// NewExample creates a new plugin instance.
func NewExample() plugins.SFProcessor {
	return new(Example)
}

// GetName returns the plugin name.
func (s *Example) GetName() string {
	return pluginName
}

// Init initializes the plugin with a configuration map.
func (s *Example) Init(conf map[string]string) error {
	return nil
}

// Register registers plugin to plugin cache.
func (s *Example) Register(pc plugins.SFPluginCache) {
	pc.AddProcessor(pluginName, NewExample)
}

// Process implements the main interface of the plugin.
func (s *Example) Process(ch interface{}, wg *sync.WaitGroup) {
	cha := ch.(*plugins.FlatChannel)
	record := cha.In
	logger.Trace.Println("Example channel capacity:", cap(record))
	defer wg.Done()
	logger.Trace.Println("Starting Example")
	for {
		fc, ok := <-record
		if !ok {
			logger.Trace.Println("Channel closed. Shutting down.")
			break
		}
		logger.Info.Println(fc)
	}
	logger.Trace.Println("Exiting Example")
}

// SetOutChan sets the output channel of the plugin.
func (s *Example) SetOutChan(ch interface{}) {}

// Cleanup tears down plugin resources.
func (s *Example) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}
