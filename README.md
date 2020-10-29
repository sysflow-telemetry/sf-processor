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
The easiest way to run the SysFlow Processor is using [docker-compose](https://github.com/sysflow-telemetry/sf-processor/edit/master/docker-compose.yml). The following compose file shows how to run sf-processor with processor events exported to rsyslog.

````yaml
version: "3.5"
services:
  sf-processor:
    container_name: sf-processor
    image: sysflowtelemetry/sf-processor:latest
    privileged: true
    volumes:
      - socket-vol:/sock/
    environment:
      FILE_PATH: /sock/sysflow.sock
      POLICYENGINE_MODE: alert
      EXPORTER_TYPE: telemetry
      EXPORTER_SOURCE: ${HOSTNAME}
      EXPORTER_EXPORT: terminal
      EXPORTER_HOST: localhost
      EXPORTER_PORT: 514
  sf-collector:
    container_name: sf-collector
    image: sysflowtelemetry/sf-collector:latest
    depends_on:
      - "sf-processor"
    privileged: true
    volumes:
      - /var/run/docker.sock:/host/var/run/docker.sock 
      - /dev:/host/dev 
      - /proc:/host/proc:ro 
      - /boot:/host/boot:ro 
      - /lib/modules:/host/lib/modules:ro 
      - /usr:/host/usr:ro
      - /mnt/data:/mnt/data
      - socket-vol:/sock/
      - ./resources/traces:/tests/traces
    environment:
      EXPORTER_ID: local
      NODE_IP: "127.0.0.1"
      FILTER: "container.name!=sf-collector and container.name!=sf-processor" 
      INTERVAL: 300 
      SOCK_FILE: /sock/sysflow.sock
volumes:
  socket-vol:
```

Instructions for `docker compose` and `helm` deployments are available in [here](https://sysflow.readthedocs.io/en/latest/deploy.html).

### Configuration

Create a JSON file specifying the edge processing pipeline plugins and configuration settings.
See [template](https://github.com/sysflow-telemetry/sf-processor/blob/master/driver/pipeline.template.json) for available options. The config settings can also be overridden by setting environment variables following the convension \<PLUGINNAME\>\_\<CONFIGKEY\>. For example, you can override _export_ in the exporter plugin by specifying ```-E EXPORTER_TYPE=file``` when running the container.

# License

View [license information](https://github.com/sysflow-telemetry/sf-exporter/blob/master/LICENSE.md) for the software contained in this image.

As with all Docker images, these likely also contain other software which may be under other licenses (such as Bash, etc from the base distribution, along with any direct or indirect dependencies of the primary software being contained).

As for any pre-built image usage, it is the image user's responsibility to ensure that any use of this image complies with any relevant licenses for all software contained within.
