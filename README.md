# SysFlow Edge Processor

## Build
To build the processor, run: 
```
cd driver
go get ./...
go build .
```

## Usage 
For usage, after build, run: 
```
./driver -h
Usage: sysprocessor [-input <value>] path

Positional arguments:
  path string
        Input path

Arguments:
  -config string
        Path to pipeline configuration file (default "/usr/local/sf-processor/conf/pipeline.json")
  -cpuprofile file
        Write cpu profile to file
  -input string
        Input type {file|socket} (default "file")
  -memprofile file
        Write memory profile to file
```

## Configuration
Create a JSON file specifying the edge processing pipeline plugins and configuration settings. See template below.

```
{
   "_comment": "DO NOT EDIT THIS TEMPLATE (remove this attribute when copying)",
   "pipeline":[
     {
      "processor": "SysFlowProc",
      "handler": "Flattener",
      "in": "sysflow SysFlowChan",
      "out": "flat FlattenerChan"
     },
     {
      "processor": "PolicyEngine",
      "in": "flat FlattenerChan",
      "out": "evt EventChan",
      "policies": "file|dir path (default: /usr/local/sf-processor/conf/)",
      "mode": "alert|filter (default: alert)"
     },
     {
      "processor": "Exporter",
      "in": "evt EventChan",
      "export": "terminal|file|syslog (default: terminal)",
      "type": "alert|telemetry (default: alert)",
      "format": "json",
      "flat": "false|true (default: false)",
      "path": "output file path (default: ./export.out)",      
      "proto": "rsyslog protocol tcp|udp (default: tcp)",
      "tag": "rsyslog tag (default: sysflow)",
      "host": "rsyslog host (default: localhost)",
      "port": "ryslog port (default: 514)",
      "buffer": "event aggregation buffer (default: 0)"
     }
   ]
}

```
