## Docker usage

Documentation and scripts for how to deploy the SysFlow Processor with docker compose can be found in [here](https://sysflow.readthedocs.io/en/latest/docker.html). 

### Processor environment

As mentioned in a previous section, all custom plugin attributes can be set using the following: `<PLUGIN NAME>_<CONFIG ATTRIBUTE NAME>` format.  Note that the docker compose file sets several attributes including `EXPORTER_TYPE`, `EXPORTER_HOST` and `EXPORTER_PORT`.

The following are the default locations of the pipeline configuration and plugins directory:

- pipeline.json:  `/usr/local/sysflow/conf/pipeline.json`
- drivers dir: `/usr/local/sysflow/resources/drivers`
- plugins dir: `/usr/local/sysflow/resources/plugins`
- handler dir: `/usr/local/sysflow/resources/handlers`
- actions dir: `/usr/local/sysflow/resources/actions`

The default configuration can be changed by setting up a virtual mounts mapping the host directories/files into the container using the volumes section of the sf-processor in the docker-compose.yaml.

```yaml
sf-processor:
    container_name: sf-processor
    image: sysflowtelemetry/sf-processor:latest
    privileged: true
    volumes:
      ...
      - ./path/to/my.pipeline.json:/usr/local/sysflow/conf/pipeline.json      
```

The policy location can be overwritten by setting the `POLICYENGINE_POLICIES` environment variable, which can point to a policy file or a directory containing policy files (must have yaml extension).

The docker container uses a default `filter.yaml` policy that outputs SysFlow records in json. You can use your own policy files from the host by mounting your policy directory into the container as follows, in which the custom pipeline points to the mounted policies.

```yaml
sf-processor:
    container_name: sf-processor
    image: sysflowtelemetry/sf-processor:latest
    privileged: true
    volumes:  
      ...    
      - ./path/to/my.pipeline.json:/usr/local/sysflow/conf/pipeline.json
      - ./path/to/policies/:/usr/local/sysflow/resources/policies/
```
