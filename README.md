# SysFlow Processor

The SysFlow Processor implements a pluggable stream-processing pipeline and contains a built-in policy engine that evaluates rules on the ingested SysFlow stream.

## Build

To build the processor locally:

```
make build
```

To build the processor container:

```
make docker-build
```

## Dynamic plugins
To build the example dynamic plugin:

```
cd plugins/example && make
```

This will install the example plugin under `resources/plugins`. To test it, use the example pipeline configuration `pipeline.example.json`.

## Usage

For usage, after build, run:

```
./sfprocessor -h
Usage: sfprocessor [[-version]|[-input <value>] [-log <value>] [-plugdir <value>] path]

Positional arguments:
  path string
        Input path

Arguments:
  -config string
        Path to pipeline configuration file (default "pipeline.json")
  -cpuprofile file
        Write cpu profile to file
  -input string
        Input type {file|socket} (default "file")
  -log string
        Log level {trace|info|warn|error} (default "info")
  -memprofile file
        Write memory profile to file
  -plugdir string
        Dynamic plugins directory (default "../resources/plugins")
  -version
        Outputs version information
```

## Configuration

Create a JSON file specifying the edge processing pipeline plugins and configuration settings.
See template below for options; driver/pipeline.json contains default values.

pipeline.json config settings can be overridden by environment variables from Dockerfile.
The convension is \<PLUGINNAME\>\_\<CONFIGKEY\>.
For example, you can override export type in the exporter plugin by doing:
$ export EXPORTER_TYPE=file

```json
{
   "_comment": "DO NOT EDIT THIS TEMPLATE (remove this attribute when copying)",
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
      "policies": "file|dir path (default: /usr/local/sf-processor/conf/)",
      "mode": "alert|filter (default: alert)"
     },
     {
      "processor": "exporter",
      "in": "evt eventchan",
      "export": "terminal|file|syslog (default: terminal)",
      "type": "alert|telemetry (default: alert)",
      "format": "json",
      "flat": "false|true (default: false)",
      "path": "output file path (default: ./export.out)",
      "proto": "rsyslog protocol tcp|udp|tcp+tls (default: tcp)",
      "tag": "rsyslog tag (default: sysflow)",
      "source": "rsyslog source hostname (default: hostname)",
      "host": "rsyslog host (default: localhost)",
      "port": "ryslog port (default: 514)",
      "buffer": "event aggregation buffer (default: 0)"
     }
   ]
}
```
