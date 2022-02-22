# Action Plugins

User-defined actions can be plugged to SysFlow's Policy Engine rule declarations to perform additional processing on matched records.

## Interface

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

## Pre-requisites

* Go 1.17 (if building locally, without the plugin builder)

## Build

The `now` action is a pluggable action that creates a tag containing the current time in nanosecond precision.

First, in the root of sf-processor, build the processor and the action plugin. Note, this plugin's shared object is generated in `resources/actions/now.so`.

```bash
make build && make -C plugins/actions/example
```

Then, run:

```bash
cd driver && ./sfprocessor -log=quiet -config=../plugins/actions/example/pipeline.actions.json ../resources/traces/tcp.sf
```

## Plugin builder

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
