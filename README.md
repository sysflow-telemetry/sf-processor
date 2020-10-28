[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sysflowtelemetry/sf-processor)](https://hub.docker.com/r/sysflowtelemetry/sf-processor/builds)
[![Docker Pulls](https://img.shields.io/docker/pulls/sysflowtelemetry/sf-processor)](https://hub.docker.com/r/sysflowtelemetry/sf-processor)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/sysflow-telemetry/sf-processor)
[![Documentation Status](https://readthedocs.org/projects/sysflow/badge/?version=latest)](https://sysflow.readthedocs.io/en/latest/?badge=latest)
[![GitHub](https://img.shields.io/github/license/sysflow-telemetry/sf-processor)](https://github.com/sysflow-telemetry/sf-processor/blob/master/LICENSE.md)

# Supported tags and respective `Dockerfile` links

-	[`0.1.0`](https://github.com/sysflow-telemetry/sf-processor/blob/0.1.0/Dockerfile), [`latest`](https://github.com/sysflow-telemetry/sf-processor/blob/master/Dockerfile)

# Quick reference

-	**Documentation**:  
	[the SysFlow Documentation](https://sysflow.readthedocs.io)
  
-	**Where to get help**:  
	[the SysFlow Community Slack](https://join.slack.com/t/sysflow-telemetry/shared_invite/enQtODA5OTA3NjE0MTAzLTlkMGJlZDQzYTc3MzhjMzUwNDExNmYyNWY0NWIwODNjYmRhYWEwNGU0ZmFkNGQ2NzVmYjYxMWFjYTM1MzA5YWQ)

-	**Where to file issues**:  
	[the github issue tracker](https://github.com/sysflow-telemetry/sf-docs/issues) (include the `sf-processor` tag)

-	**Source of this description**:  
	[repo's readme](https://github.com/sysflow-telemetry/sf-processor/edit/master/README.md) ([history](https://github.com/sysflow-telemetry/sf-processor/commits/master))

# What is SysFlow?

The SysFlow Telemetry Pipeline is a framework for monitoring cloud workloads and for creating performance and security analytics. The goal of this project is to build all the plumbing required for system telemetry so that users can focus on writing and sharing analytics on a scalable, common open-source platform. The backbone of the telemetry pipeline is a new data format called SysFlow, which lifts raw system event information into an abstraction that describes process behaviors, and their relationships with containers, files, and network. This object-relational format is highly compact, yet it provides broad visibility into container clouds. We have also built several APIs that allow users to process SysFlow with their favorite toolkits. Learn more about SysFlow in the [SysFlow specification document](https://sysflow.readthedocs.io/en/latest/spec.html).

# About this image

The SysFlow processor is a lighweight edge analytics pipeline that can process and enrich SysFlow data. The processor is written in golang, and allows users to build and configure various pipelines using a set of built-in and custom plugins and drivers. Pipeline plugins are producer-consumer objects that follow an interface and pass data to one another through pre-defined channels in a multi-threaded environment. By contrast, a driver represents a data source, which pushes data to the plugins. The processor currently supports two builtin drivers, including one that reads sysflow from a file, and another that reads streaming sysflow over a domain socket. Plugins and drivers are configured using a JSON file.  

Please check [Sysflow Processor](https://sysflow.readthedocs.io/en/latest/processor.html) for documentation on deployment and configuration options.

# How to use this image

### Starting the processor
The easiest way to run the SysFlow processor is from a Docker container, with volume mounts for processor configuration. The following command shows how to run sf-processor with processor events exported to rsyslog.

```
docker run -d --privileged --name sf-collector \
	     -v /var/run/docker.sock:/host/var/run/docker.sock \
	     -v /dev:/host/dev -v /proc:/host/proc:ro \
	     -v /boot:/host/boot:ro -v /lib/modules:/host/lib/modules:ro \
             -v /usr:/host/usr:ro -v /mnt/data:/mnt/data \
             -e INTERVAL=60 \
             -e EXPORTER_ID=${HOSTNAME} \
             -e OUTPUT=/mnt/data/    \
             -e FILTER="container.name!=sf-collector and container.name!=sf-exporter" \
             --rm sysflowtelemetry/sf-collector
```
where INTERVAL denotes the time in seconds before a new trace file is generated, EXPORTER\_ID sets the exporter name, OUTPUT is the directory in which trace files are written, and FILTER is the filter expression used to filter collected events. Note: append `container.type!=host` to FILTER expression to filter host events. 

Instructions for `docker compose` and `helm` deployments are available in [here](https://sysflow.readthedocs.io/en/latest/deploy.html).


### Configuration

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


# License

View [license information](https://github.com/sysflow-telemetry/sf-exporter/blob/master/LICENSE.md) for the software contained in this image.

As with all Docker images, these likely also contain other software which may be under other licenses (such as Bash, etc from the base distribution, along with any direct or indirect dependencies of the primary software being contained).

As for any pre-built image usage, it is the image user's responsibility to ensure that any use of this image complies with any relevant licenses for all software contained within.
