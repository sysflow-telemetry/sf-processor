# Handler Plugins

User-defined handler modules can be plugged to the built-in SysFlow `processor` plugin to implement custom data processing and analytic pipelines.

## How do handler plugin work?

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

## How to run this example?

The `printer` handler is a pluggable handler that logs select SysFlow records to the standard output. This plugin doesn't define any output channels, so it acts as a plugin sink (last plugin in a pipeline).

To run this example, in the root of sf-processor, build the processor and the handler plugin. Note, this plugin's shared object is generated in `resources/handlers/printer.so`.

```bash
make build && make -C plugins/handlers/printer
```

Then, run:

```bash
cd driver && ./sfprocessor -log=info -config=../plugins/handlers/printer/pipeline.printer.json ../resources/traces/tcp.sf
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
