## Configuration

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

> **NOTE:**  This configuration can be found in:  `sf-collector/resources/pipelines/pipeline.syslog.json`

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

Channels are modelled as channel objects that have an `In` attribute representing some golang channel of objects. See [SFChannel](https://github.com/sysflow-telemetry/sf-apis/blob/master/go/plugins/processor.go) for an example. The syntax for a channel in the pipeline is `[channel name] [channel type]`.  Where channel type is the label given to the channel type at plugin registration (more on this later), and channel name is a unique identifier for the current channel instance. The name and type of an output channel in one plugin must match that of the name and type of the input channel of the next plugin in the pipeline sequence.

> **NOTE:** A plugin has exacly one input channel but it may specify more than one output channels. This allows pipeline definitions that fan out data to more than one receiver plugin similar to a Unix `tee` command. While there must be always one SysFlow reader acting as the entry point of a pipeline, a pipeline configuration may specify policy engines passing data to different exporters or a SysFlow reader passing data to different policy engines. Generally, pipelines form a tree rather being a linear structure.

### Policy engine configuration

The policy engine (`"processor": "policyengine"`) plugin is driven by a set of rules. These rules are specified in a YAML file which adopts the same syntax as the rules of the [Falco](https://falco.org/docs/rules) project. A policy engine plugin specification may have the following attributes:

- _policies_ (required for `alert` mode`): The path to the YAML rules specification file. More information on rules can be found in the [Policies](POLICIES.md) section.
- _mode_ (optional): The mode of the policy engine. Allowed values are:
  - `alert` (default): the policy engine generates rule-based alerts; `alert` is a blocking mode that drops all records that do not match any given rule. If no mode is specified, the policy engine runs in `alert` mode by default.
  - `enrich` for enriching records with additional context from the rule. In contrast to `alert`, this is a non-blocking mode which applies tagging and action enrichments to matching records as defined in the policy file. Non-matching records are passed on "as is".
- _monitor_ (optional): Specifies if changes to the policy file(s) should be monitored and updated in the policy engine.
  - `none` (default): no monitor is used.
  - `local`: the processor will monitor for changes in the policies path and update its rule set if changes are detected.
- _monitor.interval_ (optional): The interval in seconds for updating policies, if a monitor is used. (default: 30 seconds).
- _concurrency_ (optional); The number of concurrent threads for record processing. (default: 5).
- _actiondir_ (optional): The path of the directory containing the shared object files for user-defined action plugins. See the section on [User-defined Actions](POLICIES.md#user-defined-actions) for more information.

> **NOTE:** Prior to release 0.4.0, the _mode_ attribute accepted different values with different semantics. To preserve the behavior of older releases:
> - For old `alert` behavior, use `enrich` mode.
> - For old `filter` behavior, use `enrich` mode and a policy file with filter rules only.
> - For old `bypass` behavior, use `enrich` and drop the _policies_ key from the configuration.

### Exporter configuration

An exporter (`"processor": "exporter"`) plugin consists of two modules, an encoder for converting the data to a suitable format, and a transport module for sending the data to the target. Encoders target specific, i.e. for a particular export target a particular set of encoders may be used. In the exporter configuration the transport module is specified via the _export_ parameter (required). The encoder is selected via the _format_ parameter (optional). The default format is `json`.

The following table lists the currently supported exporter modules and the corresponding encoders. Additional encoders and transport modules can be implemented if need arises. If you plan to [contribute](../CONTIRBUTING.md) or want to get involved in the discussion please join the SysFlow community.

| Transport module (_export_) | Target                     | Encoders (_format_) |
|-----------------------------|----------------------------|---------------------|
| `terminal`                  | console                    | `json`, `ecs`       |
| `file`                      | local file                 | `json`, `ecs`       |
| `es`                        | ElasticSearch service      | `ecs`               |
| `syslog`                    | syslog service             | `json`, `ecs`       |
| `findings`                  | IBM Findings API           | `occurence`         |
| `null`                      |                            |                     |

Some of these combinations require additional configuration as described in the following sections. `null` is used for debugging the processor and doesn't export any data.

#### File

If _export_ is set to `file`, an additional parameter _file.path_ allows the specification of the target file.

#### Syslog

If the _export_ parameter is set to `syslog`, output to syslog is enabled and the following addtional parameters are used:

- _syslog.proto_ (optional): The protocol used for communicating with the syslog server. Allows values are `tcp`, `tls` and `udp`. Default is `tcp`.
- _syslog.tag_ (optional): The tag used for each Sysflow record in syslog. Default is `SysFlow`.
- _syslog.source_ (optional): If set adds a hostname to the syslog header.
- _syslog.host_ (optional): The hostname of the sysflow server. Default is `localhost`.
- _syslog.port_ (optional): The port of the syslow server. Default is `514`.

#### ElasticSearch

Export to ElasticSearch is enabled by setting the config parameter _export_ to `es`. The only supported _format_ for export to ElasticSearch is `ecs`.

Data export is done via bulk ingestion. The ingestion can be controlled by some additional parameters which are read when the `es` export target is selected. Required parameters specify the ES target, index and credentials. Optional parameters control some aspects of the behavior of the bulk ingestion and may have an effect on performance. You may need to adapt their valuesfor optimal performance in your environment.

- _es.addresses_ (required): A comma-separated list of ES endpoints.
- _es.index_ (required): The name of the ES index to ingest into.
- _es.username_  (required): The ES username.
- _es.password_  (required): The password for the specified ES user.
- _buffer_ (optional) The bulk size as the number of records to be ingested at once. Default is `0` but value of `0` indicates record-by-record ingestion which may be highly inefficient.
- _es.bulk.numWorkers_ (optional): The number of ingestion workers used in parallel. Default is `0` which means that the exporter uses as many workers as there are cores in the machine.
- _es.bulk.flashBuffer_ (optional): The size in bytes of the flush buffer for ingestion. It should be large enough to hold one bulk (the number of records specified in _buffer_), otherwise the bulk is broken into smaller chunks. Default is `5e+6`.
- _es.bulk.flushTimeout_ (optional): The flush buffer time threshold. Valid values are golang duration strings. Default is `30s`.

The Elastic exporter does not make any assumption on the existence or configuration of the index specified in _es.index_. If the index does not exist, Elastic will automatically create it and apply a default dynamic mapping. It may be beneficial to use an explicit mapping for the ECS data generated by the Elastic exporter. For convinience we provide an [explicit mapping](resources/mappings/ecs_mapping.json) for creating a new tailored index in Elastic. For more information refer to the [Elastic Mapping](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping.html) reference.

<!--
#### IBM Findings

Export to IBM Findings API allows adding custom findings to the IBM Cloud Security & Compliance Center (SCC). The mode is enabled via setting the configuration parameter _export_ to `findings`. The _format_ parameter must be set to `occurence` in this case. For export to IBM Findings, the following parameters are used:

- _findings.apikey_ (required): The API key used for the Advisor service instance.
- _findings.url_ (required): The URL of the Advisor service instance.
- _findings.accountid_ (required): The account ID used for the Advisor service instance.
- _findings.provider_ (required): Unique ID of the note provider
- _findings.region_ (required): The cloud region of Advisor service instance.
- _findings.sqlqueryurl_ (required):
- _findings.sqlquerycrn_ (required):
- _findings.s3region_ (required):
- _findings.s3bucket_ (required):
- _findings.path_ (required):
- _findings.pool.capacity_ (optional): The capacity of the findings pool, Default is `250`.
- _findings.pool.maxage_ (woptional): The maximum age of the security findings in the pool in minutes. Default is `1440`.

For more information about inserting custom findings into IBM SCC, refer to [Custom Findings](https://cloud.ibm.com/docs/security-advisor?topic=security-advisor-setup_custom) section of IBM Cloud Security Advisor.
-->

### Environment variables

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

### Rate limiter configuration (experimental)

The `flattener` handler has a built-in time decay filter that can be enabled to reduce even rates in the processor. The filter uses a time-decay bloom filter based on a semantic hashing of records. This means that the filter should only forward one record matching a semantic hash per time decay period. The semantic hash takes into consideration process, flow and event attributes. To enable rate limiting, modify the `sysflowreader` processor as follows:

```json
{
     "processor": "sysflowreader",
     "handler": "flattener",
     "in": "sysflow sysflowchan",
     "out": "flat flattenerchan",
     "filter.enabled": "on|off (default: off)",
     "filter.maxage": "time decay in minutes (default: 24H)"
}
```
