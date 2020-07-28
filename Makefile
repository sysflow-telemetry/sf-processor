#
# Copyright (C) 2020 IBM Corporation.
#
# Authors:
# Frederico Araujo <frederico.araujo@ibm.com>
# Teryl Taylor <terylt@ibm.com>
#

# Build environment configuration
include ./makefile.manifest.inc

# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build -tags exclude_graphdriver_btrfs
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -tags exclude_graphdriver_btrfs
BIN=sfprocessor
OUTPUT=$(BIN)
SRC=./driver

.PHONY: deps
build: deps
	cd $(SRC) && $(GOBUILD) -o $(OUTPUT) -v

.PHONY: build
deps:
	cd $(SRC) && $(GOGET) ./...

.PHONY: version
version:
	cp $(SRC)/manifest.go.in $(SRC)/manifest.go
	sed -ibak -e "s/SYSFLOW_VERSION/$(SYSFLOW_VERSION)/" -e "s/\"JSON_SCHEMA_VERSION\"/$(SYSFLOW_JSON_SCHEMA_VERSION)/" -e "s/BUILD_NUMBER/$(SYSFLOW_BUILD_NUMBER)/" $(SRC)/manifest.go
	rm $(SRC)/manifest.goback

.PHONY: test
test:
	cd $(SRC) && $(GOTEST) -v ./...

.PHONY: clean
clean:
	cd $(SRC) && $(GOCLEAN)
	rm -f $(SRC)/$(BIN)

.PHONY: build
install: build
	mkdir -p /usr/local/sf-processor/bin /usr/local/sf-processor/conf /usr/local/sf-processor/policies
	cp ./driver/sfprocessor /usr/local/sf-processor/bin/sfprocessor
	cp ./driver/pipeline.json /usr/local/sf-processor/conf/pipeline.json
	cp ./resources/policies/distribution/* /usr/local/sf-processor/policies/

.PHONY: docker-build
docker-build: build
	sudo docker build -t sf-processor --target=runtime -f Dockerfile .

.PHONY: pull
pull:
	git pull origin master

.PHONY: up
up: 
	sudo docker-compose -f docker-compose.yml up

