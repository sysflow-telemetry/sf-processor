# SysFlow Processor (sf-processor repo)

The SysFlow processor is a lighweight edge analytics pipeline that can process and enrich SysFlow data. The processor is written in golang, and allows users to build and configure various pipelines using a set of built-in and custom plugins and drivers. Pipeline plugins are producer-consumer objects that follow an interface and pass data to one another through pre-defined channels in a multi-threaded environment. By contrast, a driver represents a data source, which pushes data to the plugins. The processor currently supports two builtin drivers, including one that reads sysflow from a file, and another that reads streaming sysflow over a domain socket. Plugins and drivers are configured using a JSON file.  

A core built-in plugin is a policy engine that can apply logical rules to filter, alert, or semantically label sysflow records using a declarative language based on the [Falco rules syntax](https://falco.org/docs/rules/) with a few added extensions (more on this later).

Custom plugins and drivers can be implemented as dynamic libraries to tailor analytics to specific user requirements.

The endpoint of a pipeline configuration is an exporter plugin that sends the processed data to a target. The processor supports various types of export plugins for a variety of different targets.

## Prerequisites

The processor has been tested on Ubuntu/RHEL distributions, but should work on any Linux system.

- Golang version 1.14+ and make (if buiding from sources)
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
make docker-build
```

## Usage

For usage information, type:

```bash
cd driver/
./sfprocessor -help
```

This should yield the following usage statement:

```bash
Usage: sfprocessor [[-version]|[-driver <value>] [-log <value>] [-driverdir <value>] [-plugdir <value>] path]
Positional arguments:
  path string
        Input path
Arguments:
  -config string
        Path to pipeline configuration file (default "pipeline.json")
  -cpuprofile file
        Write cpu profile to file
  -driver string
        Driver name {file|socket|<custom>} (default "file")
  -driverdir string
        Dynamic driver directory (default "../resources/drivers")
  -log string
        Log level {trace|info|warn|error} (default "info")
  -memprofile file
        Write memory profile to file
  -plugdir string
        Dynamic plugins directory (default "../resources/plugins")
  -test
        Test pipeline configuration
  -traceprofile file
        Write trace profile to file
  -version
        Output version information
```

The four most important flags are `config`, `driverdir`, `plugdir`, and `driver`. The `config` flag points to a pipeline configuration file, which describes the entire pipeline and settings for the individual settings for the plugins. The `driverdir` and `plugdir` flags specify where any dynamic drivers and plugins shared libraries reside that should be loaded by the processor at runtime. The `driver` flag accepts a label to a pre-configured driver (either built-in or custom) that will be used as the data source to the pipeline. Currently, the pipeline only supports one driver at a time, but we anticipate handling multiple drivers in the future. There are two built-in drivers:

- _file_: loads a sysflow file reading driver that reads from `path`.  
- _socket_: the processor loads a sysflow streaming driver. The driver creates a domain socket named `path`
  and acts as a server waiting for a SysFlow collector to attach and send sysflow data.
