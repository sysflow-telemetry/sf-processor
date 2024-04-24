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
PATH = $(shell printenv PATH):/usr/local/go/bin
BACKEND_TAG ?= flatrecord
GOCMD = go
GOBUILD = $(GOCMD) build -trimpath -tags "exclude_graphdriver_btrfs ${BACKEND_TAG}"
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test -tags "exclude_graphdriver_btrfs ${BACKEND_TAG}"
GOGET = $(GOCMD) get -tags "exclude_graphdriver_btrfs ${BACKEND_TAG}"
BIN = sfprocessor
OUTPUT = $(BIN)
SRC = ./driver
PACKDIR = ./scripts/cpack
INSTALL_PATH = /usr/local/sysflow

.PHONY: build
build: version deps
	cd $(SRC) && $(GOBUILD) -o $(OUTPUT) -v

.PHONY: package
package: 
	docker run --rm --user $(id -u):$(id -g) --group-add users --entrypoint=/bin/bash \
		-v $(shell pwd)/scripts:$(INSTALL_PATH)/scripts \
		-v $(shell pwd)/resources:$(INSTALL_PATH)/resources \
		-v $(shell pwd)/LICENSE.md:$(INSTALL_PATH)/LICENSE.md \
		-v $(shell pwd)/README.md:$(INSTALL_PATH)/README.md \
		sysflowtelemetry/sf-processor:${SYSFLOW_VERSION} -- $(INSTALL_PATH)/scripts/cpack/prepackage.sh
	cd scripts/cpack && export SYSFLOW_VERSION=$(SYSFLOW_VERSION) && cpack --config ./CPackConfig.cmake

.PHONY: deps
deps:
	cd $(SRC) && $(GOGET)  ./...

.PHONY: version
version:
	cp $(SRC)/manifest/manifest.go.in $(SRC)/manifest/manifest.go
	sed -ibak -e "s/SYSFLOW_VERSION/$(SYSFLOW_VERSION)/" -e "s/\"JSON_SCHEMA_VERSION\"/$(SYSFLOW_JSON_SCHEMA_VERSION)/" -e "s/BUILD_NUMBER/$(SYSFLOW_BUILD_NUMBER)/" -e "s/ECS_VERSION/$(SYSFLOW_ECS_VERSION)/" $(SRC)/manifest/manifest.go
	rm -f $(SRC)/manifest/manifest.gobak

.PHONY: test
test:
	cd $(SRC) && $(GOTEST) ./...

.PHONY: clean
clean:
	cd $(SRC) && $(GOCLEAN)
	rm -f $(SRC)/$(BIN)
	rm -f $(SRC)/manifest/manifest.go
	cd $(PACKDIR) && ./clean.sh

.PHONY: install
install: build
	mkdir -p /usr/local/sysflow/bin /usr/local/sysflow/conf /usr/local/sysflow/resources/policies
	cp ./driver/sfprocessor /usr/local/sysflow/bin/sfprocessor
	cp ./resources/pipelines/pipeline.distribution.json /usr/local/sysflow/conf/pipeline.json
	cp ./resources/policies/distribution/* /usr/local/sysflow/resources/policies/

.PHONY: docker-build
docker-build: docker-plugin-builder	docker-processor

.PHONY: docker-processor
docker-processor:
	( DOCKER_BUILDKIT=1 docker build --cache-from=sysflowtelemetry/plugin-builder:${SYSFLOW_VERSION} -t sysflowtelemetry/sf-processor:${SYSFLOW_VERSION} --build-arg BACKEND_TAG=$(BACKEND_TAG) --build-arg UBI_VER=$(UBI_VERSION) --target=runtime -f Dockerfile . )

.PHONY: docker-plugin-builder
docker-plugin-builder:
	( DOCKER_BUILDKIT=1 docker build -t sysflowtelemetry/plugin-builder:${SYSFLOW_VERSION} --build-arg BACKEND_TAG=$(BACKEND_TAG) --build-arg UBI_VER=$(UBI_VERSION) --target=base -f Dockerfile . )

.PHONY: pull
pull:
	git pull origin master

.PHONY: up
up: 
	sudo docker-compose -f docker-compose.yml up

.PHONY: plugins
plugins:
	@for dir in `find plugins -type d`; do \
	    if [ -f $${dir}/Makefile ]; then \
	        $(MAKE) -C $${dir}; \
            fi; \
	done

