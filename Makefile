# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN=sfprocessor
OUTPUT=$(BIN)
#SRC=github.ibm.com/sysflow/sf-processor/driver
SRC=./driver

build: deps
	cd $(SRC) && $(GOBUILD) -o $(OUTPUT) -v

deps:
	cd $(SRC) && $(GOGET) ./...

test:
	cd $(SRC) && $(GOTEST) -v ./...

clean:
	cd $(SRC) && $(GOCLEAN)
	rm -f $(SRC)/$(BIN)

install: build
	mkdir -p /usr/local/sf-processor/bin /usr/local/sf-processor/conf /usr/local/sf-processor/policies
	cp ./driver/sfprocessor /usr/local/sf-processor/bin/sfprocessor
	cp ./driver/pipeline.json /usr/local/sf-processor/conf/pipeline.json
	cp ./policies/* /usr/local/sf-processor/policies/

docker-build: build
	sudo docker build -t sf-processor:latest --target=runtime -f Dockerfile.processor .

pull:
	git pull origin master

up: 
	sudo docker-compose -f docker-compose.processor.yml up

