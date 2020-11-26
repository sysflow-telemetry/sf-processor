## Deploy collector and processor using docker

The following docker-compose file illustrates how to deploy the processor and the collector as containers.

```yaml
version: "3.5"
services:
  sf-processor:
    container_name: sf-processor
    image: sysflowtelemetry/sf-processor:latest
    privileged: false
    volumes:
      - socket-vol:/sock/
    environment:
      DRIVER: socket
      INPUT_PATH: /sock/sysflow.sock
      POLICYENGINE_MODE: alert
      EXPORTER_TYPE: telemetry
      EXPORTER_SOURCE: sysflow
      EXPORTER_EXPORT: syslog
      EXPORTER_HOST: <IP address of the syslog server>
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
      EXPORTER_ID: ${HOSTNAME}
      NODE_IP: <Host IP address>
      FILTER: "container.name!=sf-collector and container.name!=sf-processor" 
      INTERVAL: 300 
      SOCK_FILE: /sock/sysflow.sock
volumes:
  socket-vol:
```  

### Setting up the collector environment

The key setting in the collector portion of the file is the `FILTER` variable.  Since the collector is built atop the sysdig core, it uses the sysdig filtering mechanism described [here](https://github.com/draios/sysdig/wiki/Sysdig-User-Guide#user-content-filtering) and can support all the sysdig attributes described there in case you want to filter on specific containers, processes, operations, etc.  One of the more powerful filters is the `container.type!=host` filter, which limits collection only to container monitoring.  If you want to monitor the entire host, simply remove the `container.type` operation from the filter.

### Setting up the processor environment

As mentioned in a previous section, all custom plugin attributes can be set using the the following: `<PLUGIN NAME>_<CONFIG ATTRIBUTE NAME>` format.  Note that the docker compose file sets several attributes including `EXPORTER_TYPE`, `EXPORTER_HOST` and `EXPORTER_PORT`. Note that `EXPORTER_SOURCE` is set to the bash environment variable `${HOSTNAME}`.  `HOSTNAME` must be explicitly exported before launching docker compose in order to be picked up. 

```bash
export HOSTNAME
docker-compose up
```

The following are the default locations of the pipeline configuration and plugins directory:

- pipeline.json -  `/usr/local/sysflow/conf/pipeline.json`
- plugins dir - `/usr/local/sysflow/resources/plugins`

We can overwrite these particular files/dirs in the docker container with those on the host by setting up a virtual mounts mapping the host directories/files into the container using the volumes section of the sf-processor in the docker-compose.yaml.

```yaml
sf-processor:
    container_name: sf-processor
    image: sysflowtelemetry/sf-processor:latest
    privileged: true
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - socket-vol:/sock/
      - ./resources/pipelines/pipeline.runtimeintegrity.json:/usr/local/sysflow/conf/pipeline.json
      - ./resources/plugins:/usr/local/sysflow/resources/plugins
```

If using the policy engine, the policy folder defaults to the following location in the container:

`/usr/local/sysflow/resources/policies/`

This location can be overwritten by setting the  `POLICYENGINE_POLICIES` environment variable.

The docker container uses a default `filter.yaml` policy that outputs sysflow records in json.  You can use your own policy files from the host by mounting your policy directory into the container as follows:

```yaml
sf-processor:
    container_name: sf-processor
    image: sysflowtelemetry/sf-processor:latest
    privileged: true
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - socket-vol:/sock/
      - ./resources/policies/runtimeintegrity/:/usr/local/sysflow/resources/policies/
```