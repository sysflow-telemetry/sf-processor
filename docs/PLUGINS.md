## Plugins

In addition to its core plugins, the processor also supports custom plugins that can be dynamically loaded into the processor via a compiled golang shared library using the [golang plugin package](https://golang.org/pkg/plugin/). Custom plugins enable easy extension of the processor and the creation of custom pipelines tailored to specific use cases.

The processor supports four types of plugins:

* **sources**: enable the ingestion of different telemetry sources into the processor pipeline.
* **processors**: enable the creation of custom data processing and analytic plugins to extend sf-processor pipelines.
* **handlers**: enable the creation of custom SysFlow record handling plugins.
* **actions**: enable the creation of custom action plugins to extend sf-processor's policy engine.

### Processor Plugins

User-defined plugins can be plugged and extend the sf-processor pipeline. These are the most generic type of plugins, from which all built-in processor plugins are build. Check the `core` package for examples. We have built-in processor plugins for flattening the telemetry stream, implementing a policy engine, and creating event exporters.

#### Interface

Processor plugins (or just plugins) are implemented via the golang plugin mechanism. A plugin must implement the following interface, defined in the `github.com/sysflow-telemetry/sf-apis/go/plugins` package.

```go
// SFProcessor defines the SysFlow processor interface.
type SFProcessor interface {
  Register(pc SFPluginCache)
  Init(conf map[string]interface{}) error
  Process(ch interface{}, wg *sync.WaitGroup)
  GetName() string
  SetOutChan(ch []interface{})
  Cleanup()
}
```

The `Process` function is the main function of the plugin.It's where the "main loop" of the plugin should be implemented. It receives the input channel configured in the custom plugin's block in the pipeline configuration. It also received the pepeline thread WaitGroup. Custom processing code should be implemented using this function. `Init` is called once, when the pipeline is loaded. `Cleanup` is called when the pipeline is terminated. `SetOutChannel` receives a slice with the output channels configured in the plugin's block in the pipeline configuration.

When loading a pipeline, sf-processor performs a series of health checks before the pipeline is enabled. If these health checks fail, the processor terminates. To enable health checks on custom plugins, implement the `Test` function defined in the interface below. For an example, check `core/exporter/exporter.go`.

```go
// SFTestableProcessor defines a testable SysFlow processor interface.
type SFTestableProcessor interface {
  SFProcessor
  Test() (bool, error)
}
```

#### Example

A dynamic plugin example is provided in [github](https://github.com/sysflow-telemetry/sf-processor/tree/master/plugins/processors/example). The core of the plugin is building an object that implements an [SFProcessor interface](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go). Such an implementation looks as follows:

```golang
package main

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
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
func (s *Example) Init(conf map[string]interface{}) error {
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
		if fc.Ints[sfgo.SYSFLOW_IDX][sfgo.SF_REC_TYPE] == sfgo.PROC_EVT {
			logger.Info.Printf("Process Event: %s, %d", fc.Strs[sfgo.SYSFLOW_IDX][sfgo.PROC_EXE_STR], fc.Ints[sfgo.SYSFLOW_IDX][sfgo.EV_PROC_TID_INT])
		}
	}
	logger.Trace.Println("Exiting Example")
}

// SetOutChan sets the output channel of the plugin.
func (s *Example) SetOutChan(ch []interface{}) {}

// Cleanup tears down plugin resources.
func (s *Example) Cleanup() {}

// This function is not run when module is used as a plugin.
func main() {}

```

The custom plugin must implement the following interface:  

* `GetName()` - returns a lowercase string representing the plugin's label.  This label is important, because it identifies the plugin in the `pipeline.json` file, enabling the processor to load the plugin. In the object above, this plugin is called `example`. Note that the label must be unique.
* `Init(conf map[string]interface{}) error` - used to initialize the plugin.  The configuration map that is passed to the function stores all the configuration information defined in the plugin's definition inside `pipeline.json` (more on this later).
* `Register(pc plugins.SFPluginCache)` - this registers the plugin with the plugin cache of the processor.
    * `pc.AddProcessor(pluginName, <plugin constructor function>)` (required) - registers the plugin named `example` with the processor.  You must define a constructor function using the convention `New<PluginName>` which is used to instantiate the plugin, and returns it as an `SFProcessor` interface - see `NewExample` for an example.
    * `pc.AddChannel(channelName, <output channel constructor function>)` (optional)  - if your plugin is using a custom output channel of objects (i.e., the channel used to pass output objects from this plugin to the next in the pipeline), it should be included in this plugin.  
         * The `channelName` should be a lowercase unique label defining the channel type.  
         * The constructor function should return a golang `interface{}` representing an object that as an `In` attribute of type `chan <ObjectToBePassed>`.  We will call this object, a wrapped channel object going forward.  For example, the channel object that passes sysflow objects is called SFChannel, and is defined [here](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go)
         * For a complete example of defining an output channel, see `NewFlattenerChan` in the [flattener](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/flattener/flattener.go) as well as the `Register` function.  The `FlatChannel` is defined [here](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/handler.go)
* `Process(ch interface{}, wg *sync.WaitGroup)`  - this function is launched by the processor as a go thread and is where the main plugin processing occurs.  It takes a wrapped channel object, which acts as the input data source to the plugin (i.e., this is the channel that is configured as the input channel to the plugin in the pipeline.json).  It also takes a sync.WaitGroup object, which is used to signal to the processor when the plugin has completed running (see `defer wg.Done()` in code).  The processor must loop on the input channel, and do its analysis on each input record.  In this case, the example plugin is reading flat records and printing them to the screen. 
* `SetOutChan(ch []interface{})` - sets the wrapped channels that will serve as the output channels for the plugin.  The output channels are instantiated by the processor, which is also in charge of stitching the plugins together.  If the plugin is the last one in the chain, then this function can be left empty. See the `SetOutputChan` function in the [flattener](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/flattener/flattener.go) to see how an output channel is implemented.
* `Cleanup()` - Used to cleanup any resources.  This function is called by the processor after the plugin `Process` function exits.  One of the key items to close in the `Cleanup` function is the output channel using the golang `close()` [function](https://gobyexample.com/closing-channels).  Closing the output channel enables the pipeline to be torn down gracefully and in sequence.
* `main(){}` - this main method is not used by the plugin or processor.  It's required by golang in order to be able to compile as a shared object.

To compile the example plugin, use the provided Makefile:

```bash
make -C plugins/processors/example
```

This will build the plugin and copy it into `resources/plugins/`.

To use the new plugin, use the configuration provided in [github](https://github.com/sysflow-telemetry/sf-processor/tree/master/plugins/processors/example), which defines the following pipeline:

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
* The builtin `sysflowReader` plugin with flattener handler, which takes raw sysflow objects, and flattens them
   into arrays of integers and strings for easier processing in certain plugins like the policy engine.
* The `example` plugin, which takes the flattened output from the sysflowreader plugin, and prints it the screen.

The key item to note is that the output channel (i.e., `out`) of `sysflowreader` matches the input channel (i.e., `in`) of the example plugin. This ensures that the plugins will be properly stitched together.

#### Build

The `example` plugin is a custom plugin that illustrates how to implement a minimal plugin that reads the records from the input channel and logs them to the standard output.

To run this example, in the root of sf-processor, build the processor and the example plugin. Note, this plugin's shared object is generated in `resources/plugins/example.so`.

```bash
make build && make -C plugins/processors/example
```

Then, run:

```bash
cd driver && ./sfprocessor -log=info -config=../plugins/processors/example/pipeline.example.json ../resources/traces/tcp.sf
```

#### Plugin builder

To build the plugin for release, Go requires the code to be compiled with the exact package versions that the SysFlow processor was compiled with. The easiest way to achieve this is to use the pre-built `plugin-builder` Docker image in your build. This option also works for building plugins for deployment with the SysFlow binary packages.

Below is an example of how this can be achieved. Set $TAG to a SysFlow release (>=0.4.0), `edge`, or `dev`.

First, build the plugin:

```bash
docker run --rm \
    -v $(pwd)/plugins:/go/src/github.com/sysflow-telemetry/sf-processor/plugins \
    -v $(pwd)/resources:/go/src/github.com/sysflow-telemetry/sf-processor/resources \
    sysflowtelemetry/plugin-builder:$TAG \
    make -C /go/src/github.com/sysflow-telemetry/sf-processor/plugins/processors/example
```

To test it, run the pre-built processor with the example configuration and trace.

```bash
docker run --rm \
    -v $(pwd)/plugins:/usr/local/sysflow/plugins \
    -v $(pwd)/resources:/usr/local/sysflow/resources \
    -w /usr/local/sysflow/bin \
    --entrypoint=/usr/local/sysflow/bin/sfprocessor \
    sysflowtelemetry/sf-processor:$TAG \
    -log=info -config=../plugins/processors/example/pipeline.example.json ../resources/traces/tcp.sf
```

The output on the above pre-recorded trace should look like this:

```plain
[Health] 2022/02/21 12:55:19 pipeline.go:246: Health checks: passed
[Info] 2022/02/21 12:55:19 main.go:147: Successfully loaded pipeline configuration
[Info] 2022/02/21 12:55:19 pipeline.go:170: Starting the processing pipeline
[Info] 2022/02/21 12:55:19 example.go:75: Process Event: ./server, 13823
[Info] 2022/02/21 12:55:19 example.go:75: Process Event: ./client, 13824
[Info] 2022/02/21 12:55:19 example.go:75: Process Event: ./client, 13824
[Info] 2022/02/21 12:55:19 example.go:75: Process Event: ./server, 13823
```

### Handler Plugins

User-defined handler modules can be plugged to the built-in SysFlow `processor` plugin to implement custom data processing and analytic pipelines.

#### Interface

Handlers are implemented via the golang plugin mechanism. A handler must implement the following interface, defined in the `github.com/sysflow-telemetry/sf-apis/go/plugins` package.

```go
// SFHandler defines the SysFlow handler interface.
type SFHandler interface {
  RegisterChannel(pc SFPluginCache)
  RegisterHandler(hc SFHandlerCache)
  Init(conf map[string]interface{}) error
  IsEntityEnabled() bool
  HandleHeader(sf *CtxSysFlow, hdr *sfgo.SFHeader) error
  HandleContainer(sf *CtxSysFlow, cont *sfgo.Container) error
  HandleProcess(sf *CtxSysFlow, proc *sfgo.Process) error
  HandleFile(sf *CtxSysFlow, file *sfgo.File) error
  HandleNetFlow(sf *CtxSysFlow, nf *sfgo.NetworkFlow) error
  HandleNetEvt(sf *CtxSysFlow, ne *sfgo.NetworkEvent) error
  HandleFileFlow(sf *CtxSysFlow, ff *sfgo.FileFlow) error
  HandleFileEvt(sf *CtxSysFlow, fe *sfgo.FileEvent) error
  HandleProcFlow(sf *CtxSysFlow, pf *sfgo.ProcessFlow) error
  HandleProcEvt(sf *CtxSysFlow, pe *sfgo.ProcessEvent) error
  SetOutChan(ch []interface{})
  Cleanup()
}
```

Each `Handle*` function receives the current SysFlow record being processed along with its corresponding parsed record type. Custom processing code should be implemented using these functions.

#### Build

The `printer` handler is a pluggable handler that logs select SysFlow records to the standard output. This plugin doesn't define any output channels, so it acts as a plugin sink (last plugin in a pipeline).

To run this example, in the root of sf-processor, build the processor and the handler plugin. Note, this plugin's shared object is generated in `resources/handlers/printer.so`.

```bash
make build && make -C plugins/handlers/printer
```

Then, run:

```bash
cd driver && ./sfprocessor -log=info -config=../plugins/handlers/printer/pipeline.printer.json ../resources/traces/tcp.sf
```

#### Plugin builder

To build the plugin for release, Go requires the code to be compiled with the exact package versions that the SysFlow processor was compiled with. The easiest way to achieve this is to use the pre-built `plugin-builder` Docker image in your build. This option also works for building plugins for deployment with the SysFlow binary packages.

Below is an example of how this can be achieved. Set $TAG to a SysFlow release (>=0.4.0), `edge`, or `dev`.

First, build the plugin:

```bash
docker run --rm \
    -v $(pwd)/plugins:/go/src/github.com/sysflow-telemetry/sf-processor/plugins \
    -v $(pwd)/resources:/go/src/github.com/sysflow-telemetry/sf-processor/resources \
    sysflowtelemetry/plugin-builder:$TAG \
    make -C /go/src/github.com/sysflow-telemetry/sf-processor/plugins/handlers/printer
```

To test it, run the pre-built processor with the example configuration and trace.

```bash
docker run --rm \
    -v $(pwd)/plugins:/usr/local/sysflow/plugins \
    -v $(pwd)/resources:/usr/local/sysflow/resources \
    -w /usr/local/sysflow/bin \
    --entrypoint=/usr/local/sysflow/bin/sfprocessor \
    sysflowtelemetry/sf-processor:$TAG \
    -log=info -config=../plugins/handlers/printer/pipeline.printer.json ../resources/traces/tcp.sf
```

The output on the above pre-recorded trace should look like this:

```plain
[Info] 2022/02/21 15:39:58 printer.go:118: ProcEvt ./server, 13823
[Info] 2022/02/21 15:39:58 printer.go:100: FileFlow ./server, 3
[Info] 2022/02/21 15:39:58 printer.go:100: FileFlow ./server, 3
[Info] 2022/02/21 15:39:58 printer.go:118: ProcEvt ./client, 13824
[Info] 2022/02/21 15:39:58 printer.go:100: FileFlow ./client, 3
[Info] 2022/02/21 15:39:58 printer.go:100: FileFlow ./client, 3
[Info] 2022/02/21 15:39:58 printer.go:94: NetworkFlow ./client, 8080
[Info] 2022/02/21 15:39:58 printer.go:118: ProcEvt ./client, 13824
[Info] 2022/02/21 15:39:58 printer.go:94: NetworkFlow ./server, 8080
[Info] 2022/02/21 15:39:58 printer.go:118: ProcEvt ./server, 13823
```

### Action Plugins

User-defined actions can be plugged to SysFlow's Policy Engine rule declarations to perform additional processing on matched records.

#### Interface

Actions are implemented via the golang plugin mechanism. An action must implement the following interface, defined in the `github.com/sysflow-telemetry/sf-processor/core/policyengine/engine` package.

```go
// Prototype of an action function
type ActionFunc func(r *Record) error

// Action interface for user-defined actions
type Action interface {
        GetName() string
        GetFunc() ActionFunc
}
```

Actions have a name and an action function. Within a single policy engine instance, action names must be unique. User-defined actions cannot re-declare built-in actions. Reusing names of user-defined actions overwrites previously registered actions.

The action function receives the current record as an argument and thus has access to all record attributes. The action result can be stored in the record context via the context modifier methods. 

#### Build

The `now` action is a pluggable action that creates a tag containing the current time in nanosecond precision.

First, in the root of sf-processor, build the processor and the action plugin. Note, this plugin's shared object is generated in `resources/actions/now.so`.

```bash
make build && make -C plugins/actions/example
```

Then, run:

```bash
cd driver && ./sfprocessor -log=quiet -config=../plugins/actions/example/pipeline.actions.json ../resources/traces/tcp.sf
```

#### Plugin builder

To build the plugin for release, Go requires the code to be compiled with the exact package versions that the SysFlow processor was compiled with. The easiest way to achieve this is to use the pre-built `plugin-builder` Docker image in your build. This option also works for building plugins for deployment with the SysFlow binary packages.

Below is an example of how this can be achieved. Set $TAG to a SysFlow release (>=0.4.0), `edge`, or `dev`.

First, build the plugin:

```bash
docker run --rm \
    -v $(pwd)/plugins:/go/src/github.com/sysflow-telemetry/sf-processor/plugins \
    -v $(pwd)/resources:/go/src/github.com/sysflow-telemetry/sf-processor/resources \
    sysflowtelemetry/plugin-builder:$TAG \
    make -C /go/src/github.com/sysflow-telemetry/sf-processor/plugins/actions/example
```

To test it, run the pre-built processor with the example configuration and trace.

```bash
docker run --rm \
    -v $(pwd)/plugins:/usr/local/sysflow/plugins \
    -v $(pwd)/resources:/usr/local/sysflow/resources \
    -w /usr/local/sysflow/bin \
    --entrypoint=/usr/local/sysflow/bin/sfprocessor \
    sysflowtelemetry/sf-processor:$TAG \
    -log=quiet -config=../plugins/actions/example/pipeline.actions.json ../resources/traces/tcp.sf
```

In the output, observe that all records matching the policy speficied in `pipeline.actions.json` are tagged by action `now` with the tag `now_in_nanos`. For example:

```plain
{
  "version": 4,
  "endts": 0,
  "opflags": [
    "EXEC"
  ],
  ...
  "policies": [
    {
      "id": "Action example",
      "desc": "user-defined action example",
      "priority": 0
    }
  ],
  "tags": [
    "now_in_nanos:1645409122055957900"
  ]
}
```
