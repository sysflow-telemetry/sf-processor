# SysFlow Processor (sf-processor repo)

The SysFlow processor is a lighweight edge analytics pipeline that can process and enrich SysFlow data. The processor is written in golang, and allows users to build and configure various pipelines using a set of built-in and custom plugins and drivers. Pipeline plugins are producer-consumer objects that follow an interface and pass data to one another through pre-defined channels in a multi-threaded environment. By contrast, a driver represents a data source, which pushes data to the plugins. The processor currently supports two builtin drivers, including one that reads sysflow from a file, and another that reads streaming sysflow over a domain socket. Plugins and drivers are configured using a JSON file.  

A core built-in plugin is a policy engine that can apply logical rules to filter, alert, or semantically label sysflow records using a declarative language based on the [Falco rules syntax](https://falco.org/docs/rules/) with a few added extensions (more on this later).

Custom plugins and drivers can be implemented as dynamic libraries to tailor analytics to specific user requirements.

## Prerequisites

The processor has been tested on Ubuntu/RHEL distributions, but should work on any Linux system.

- Golang version 1.14 and make (if buiding from sources)
- Docker, docker-compose  (if building with docker)

## Build

Clone the processor repository

```bash
git clone https://github.com/sysflow-telemetry/sf-processor.git
```

Build locally, from sources

```bash
cd sf-processor
make build
```

Build with docker

```bash
cd sf-processor
make build-docker
```

## Usage

For usage information, type:

```bash
cd driver/
./sf-processor -help
```

This should yield the follwowing usage statement:

```bash
Usage: sfprocessor [[-version]|[-driver <value>] [-log <value>] [-driverdir <value>] [-plugdir <value>] path]
Positional arguments:
  path string
        Input path
Arguments:
  -config string
        Path to pipeline configuration file (default “pipeline.json”)
  -cpuprofile file
        Write cpu profile to file
  -driver string
        Driver name {file|socket|<custom>} (default “file”)
  -driverdir string
        Dynamic driver directory (default “../resources/drivers”)
  -log string
        Log level {trace|info|warn|error} (default “info”)
  -memprofile file
        Write memory profile to file
  -plugdir string
        Dynamic plugins directory (default “../resources/plugins”)
  -version
        Outputs version information
```

The four most important flags are `config`, `driverdir`, `plugdir`, and `driver`. The `config` flag points to a pipeline configuration file, which describes the entire pipeline and settings for the individual settings for the plugins. The `driverdir` and `plugdir` flags specify where any dynamic drivers and plugins shared libraries reside that should be loaded by the processor at runtime. The `driver` flag accepts a label to a pre-configured driver (either built-in or custom) that will be used as the data source to the pipeline. Currently, the pipeline only supports one driver at a time, but we anticipate handling multiple drivers in the future. There are two built-in drivers:

- _file_: loads a sysflow file reading driver that reads from `path`.  
- _socket_: the processor loads a sysflow streaming driver. The driver creates a domain socket named `path`
  and acts as a server waiting for a SysFlow collector to attach and send sysflow data.

## Pipeline Configuration

The pipeline configuration below shows how to configure a pipeline that will read a sysflow stream and push records to the policy engine, which will trigger alerts using a set of runtime policies stored in a `yaml` file.  An example pipeline with this configuration looks as follows:  

```json
{
  "pipeline":[
    {
     "processor": "sysflowreader",
     "handler": "flattener",
     "in": "sysflow sysflowchan",
     "out": "flat flattenerchan"
    },
    {
     "processor": "policyengine",
     "in": "flat flattenerchan",
     "out": "evt eventchan",
     "policies": "../resources/policies/runtimeintegrity"
    },
    {
     "processor": "exporter",
     "in": "evt eventchan",
     "export": "syslog",
     "proto": "tcp",
     "tag": "sysflow",
     "host": "localhost",
     "port": "514"
    }
  ]
}
```

> Note:  This configuration can be found in:  `sf-collector/resources/pipelines/pipeline.runtimeintegrity.json`

This pipeline specifies three built-in plugins:

- [sysflowreader](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/processor/processor.go): is a generic reader plugin that ingests sysflow from the driver, caches entities, and presents sysflow objects to a handler object (i.e., an object that implements the [handler interface](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/handler.go)) for processing. In this case, we are using the [flattener](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/flattener/flattener.go) handler, but custom handlers are possible.
- [policyengine](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/policyengine/policyengine.go): is the policy engine, which takes [flattened](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/sfgo/flatrecord.go) (row-oriented) SysFlow records as input and outputs [records](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/policyengine/engine/types.go), which represent alerts, or filtered sysflow records depending on the policy engine's _mode_ (more on this later).  
- [exporter](https://github.com/sysflow-telemetry/sf-processor/blob/master/core/exporter/exporter.go): takes records from the policy engine, and exports them to ElasticSearch, syslog, file, or terminal, in a JSON format or in Elastic Common Schema (ECS) format. Note that custom export plugins can be created to export to other serialization formats and transport protocols.

Each plugin has a set of general attributes that are present in all plugins, and a set of attributes that are custom to the specific plugins. For more details on the specific attributes in this example, see the pipeline configuration [template](https://github.com/sysflow-telemetry/sf-processor/blob/master/driver/pipeline.template.json)

The general attributes are as follows:

- _processor_ (required): the name of the processor plugin to load. Processors must implement the [SFProcessor](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go) interface; the name is the value that must be returned from the `GetName()` function as defined in the processor object.
- _handler_ (optional): the name of the handler object to be used for the processor. Handlers must implement the [SFHandler](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/handler.go) interface.
- _in_ (required): the input channel (i.e. golang channel) of objects that are passed to the plugin.
- _out_ (optional): the output channel (i.e. golang channel) for objects that are pushed out of the plugin, and into the next plugin in the pipeline sequence.

Channels are modeled as channel objects that have an `In` attribute representing some golang channel of objects. See [SFChannel](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go) for an example. The syntax for a channel in the pipeline is `[channel name] [channel type]`.  Where channel type is the label given to the channel type at plugin registration (more on this later), and channel name is a unique identifier for the current channel instance. The name and type of an output channel in one plugin must match that of the name and type of the input channel of the next plugin in the pipeline sequence.

### Export to ElasticSearch

The export target is specified via the `export` parammeter of the exporter plugin section of the pipeline configuration. Export to ElasticSearch is enabled by setting the config parameter:

```json
"export": "es"
```

The export format can be set via the exporter parameter `format` which accepts two values: `json` (default) or `ecs`. The export format can be specified independently of the exporter type. Since ES accepts any JSON-formatted input both values are permissable. For ingestion into ES we recommend to use ECS:

```json
"format": "ecs"
```

Export to ES is done via bulk ingestion. The ingestion can be controlled by some additional parameters which are read when the `es` export target is selected. Required parameters specify the ES target, index and credentials. Optional parameters control some aspects of the behavior of the bulk ingestion and may have an effect on performance. You may need to adapt their valuesfor optimal performance in your environment.

| Parameter              | Type       | Default Value | Description |
|------------------------|------------|---------------|-------------|
| `es.addresses`         | *required* | *none*   | comma-separated list of ES endpoints |
| `es.index`             | *required* | *none*   | name of the ES index to ingest into |
| `es.username`          | *required* | *none*   | ES username |
| `es.password`          | *required* | *none*   | password for the specified ES user |
| `buffer`               | *optional* | `"0"`      | Bulk size as number of records to be ingested at once. A value of `"0"` indicates record-by-record ingestion which may be highly inefficient. |
| `es.bulk.numWorkers`   | *optional* | `"0"`      | Number of workers  used in parallel. A value of `"0"` means that the exporter uses as many workers are there are cores on the machine.        |
| `es.bulk.flashBuffer`  | *optional* | `"5+e6"`   | Size (in bytes) of the flush buffer for ingestion. It should be large enough to hold one bulk (the number of records specified in `buffer`), otherwise the bulk is broken into smaller chunks. |
| `es.bulk.flushTimeout` | *optional* | `"30s"`    | Flush buffer time threshold. Valid values are golang duration strings. |

## Override plugin configuration attributes with environment variables

It is possible to override any of the custom attributes of a plugin using an environment variable. This is especially useful when operating the processor as a container, where you may have to deploy the processor to multiple nodes, and have attributes that change per node. If an environment variable is set, it overrides the setting inside the config file. The environment variables must follow the following structure:

- Environment variables must follow the naming schema `<PLUGIN NAME>_<CONFIG ATTRIBUTE NAME>`
- The plugin name inside the pipeline configuration file must be all lower case.  

For example, to set the alert mode inside the policy engine, the following environment variable is set:

```bash
export POLICYENGINE_MODE=alert
```

To set the syslog values for the exporter:

```bash
export EXPORTER_TYPE=telemetry
export EXPORTER_SOURCE=${HOSTNAME}
export EXPORTER_EXPORT=syslog
export EXPORTER_HOST=192.168.2.10
export EXPORTER_PORT=514
```

If running as a docker container, environment variables can be passed with the docker run command:

```bash
docker run
-e EXPORTER_TYPE=telemetry \
-e EXPORTER_SOURCE=${HOSTNAME} \
-e EXPORTER_EXPORT=syslog \
-e EXPORTER_HOST=192.168.2.10 \
-e EXPORTER_PORT=514
...
```
