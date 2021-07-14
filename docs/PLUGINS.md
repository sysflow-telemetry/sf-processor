## Write a simple processing plugin

In addition to its core plugins, the processor also supports custom plugins that can be dynamically loaded into the processor via a compiled golang shared library using the [golang plugin package](https://golang.org/pkg/plugin/). Custom plugins enables easy extension of the processor and the creation of custom pipelines tailored to specific use cases.

A dynamic plugin example is provided in [github](https://github.com/sysflow-telemetry/sf-processor/tree/master/plugins/example). The core of the plugin is building an object that implements an [SFProcessor interface](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go). Such an implementation looks as follows:

```golang
package main

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-processor/core/flattener"
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
	cha := ch.(*flattener.FlatChannel)
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
```

The object must implement the following interface:  

* `GetName()` - returns a lowercase string representing the plugin's label.  This label is important, because it identifies the plugin in the `pipeline.json` file, enabling the processor to load the plugin. In the object above, this plugin is called `example`. Note that the label must be unique.
* `Init(config map[string]string) error` - used to initialize the plugin.  The configuration map that is passed to the function stores all the configuration information defined in the plugin's definition inside `pipeline.json` (more on this later).
* `Register(pc plugins.SFPluginCache)` - this registers the plugin with the plugin cache of the processor.
    * `pc.AddProcessor(pluginName, <plugin constructor function>)` (required) - registers the plugin named `example` with the processor.  You must define a constructor function using the convention `New<PluginName>` which is used to instantiate the plugin, and returns it as an `SFProcessor` interface - see `NewExample` for an example.
    * `pc.AddChannel(channelName, <output channel constructor function>)` (optional)  - if your plugin is using a custom output channel of objects (i.e., the channel used to pass output objects from this plugin to the next in the pipeline), it should be included in this plugin.  
         * The `channelName` should be a lowercase unique label defining the channel type.  
         * The constructor function should return a golang `interface{}` representing an object that as an `In` attribute of type `chan <ObjectToBePassed>`.  We will call this object, a wrapped channel object going forward.  For example, the channel object that passes sysflow objects is called SFChannel, and is defined [here](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go)
         * For a complete example of defining an output channel, see `NewFlattenerChan` in the [flattener](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/flattener/flattener.go) as well as the `Register` function.  The `FlatChannel` is defined [here](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/handler.go)
* `Process(ch interface{}, wg *sync.WaitGroup)`  - this function is launched by the processor as a go thread and is where the main plugin processing occurs.  It takes a wrapped channel object, which acts as the input data source to the plugin (i.e., this is the channel that is configured as the input channel to the plugin in the pipeline.json).  It also takes a sync.WaitGroup object, which is used to signal to the processor when the plugin has completed running (see `defer wg.Done()` in code).  The processor must loop on the input channel, and do its analysis on each input record.  In this case, the example plugin is reading flat records and printing them to the screen. 
* `SetOutChan(ch interface{})` - sets the wrapped channel that will serve as the output channel for the plugin.  The output channel is instantiated by the processor, which is also in charge of stitching the plugins together.  If the plugin is the last one in the chain, then this function can be left empty. See the `SetOutputChan` function in the [flattener](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/flattener/flattener.go) to see how an output channel is implemented.
* `Cleanup()` - Used to cleanup any resources.  This function is called by the processor after the plugin `Process` function exits.  One of the key items to close in the `Cleanup` function is the output channel using the golang `close()` [function](https://gobyexample.com/closing-channels).  Closing the output channel enables the pipeline to be torn down gracefully and in sequence.         
* `main(){}` - this main method is not used by the plugin or processor.  It's required by golang in order to be able to compile as a shared object.

To compile the example plugin, use the provided Makefile:

```
cd plugins/example
make
```

This will build the plugin and copy it into `resources/plugins/`.

To use the new plugin, use the configuration provided in [github](https://github.com/sysflow-telemetry/sf-processor/tree/master/plugins/example), which defines the following pipeline:

```bash
{
   "pipeline":[
     {
      "processor": "sysflowreader",
      "handler": "flattener",
      "in": "sysflow sysflowchan",
      "out": "flat flattenerchan"
     },
     {
      "processor": "example",
      "in": "flat flattenerchan"
     }
   ]
}
```

This pipeline contains two plugins:
- The builtin `sysflowReader` plugin with flattener handler, which takes raw sysflow objects, and flattens them
   into arrays of integers and strings for easier processing in certain plugins like the policy engine.
- The `example` plugin, which takes the flattened output from the sysflowreader plugin, and prints it the screen.

The key item to note is that the output channel (i.e., `out`) of `sysflowreader` matches the input channel (i.e., `in`) of the example plugin. This ensures that the plugins will be properly stitched together.

To run the example pipeline:

```
cd driver
./sfprocessor -config ../plugins/example/pipeline.example.json -plugdir ../resources/plugins/  ../resources/traces/mon.1531776712.sf
```
