#!/bin/bash

# compile avro IDL to avro Schema
java -jar avro-tools-1.10.2.jar idl2schemata ./avdl/Event.avdl avsc/ 

# golang stub generation
# to install gogen-avro: go get github.com/actgardner/gogen-avro/v7/cmd/..
gogen-avro --containers=true --package=event event ./avsc/Event.avsc
