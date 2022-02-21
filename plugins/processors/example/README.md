# Processor Plugins

User-defined plugins can be plugged and extend the SysFlow `processor` pipeline. These are the most generic type of plugins, from which all built-in processor plugins are build. Check the `core` package for examples. We have built-in processor plugins for flattening the telemetry stream, implementing a policy engine, and creating event exporters.

## How do data processing handlers work?

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

When loading a pipeline, sf-processor performes a series of health checks before the pipeline is enabled. If these health checks fail, the processor terminates. To enable health checks on cutom plugins, implement the `Test` function defined in the interface below.

```go
// SFTestableProcessor defines a testable SysFlow processor interface.
type SFTestableProcessor interface {
  SFProcessor
  Test() (bool, error)
}
```

## How to run this example?

The `example` plugin is a custom plugin that illustrates how to implement a minimal plugin that reads the records from the input channel and logs them to the standard output.

To run this example, in the root of sf-processor, build the processor and the example plugin. Note, this plugin's shared object is generated in `resources/plugins/example.so`.

```bash
make build && make -C plugins/processors/example
```

Then, run:

```bash
cd driver && ./sfprocessor -log=info -config=../plugins/processors/example/pipeline.example.json ../resources/traces/tcp.sf
```

## Building the plugin for release

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
