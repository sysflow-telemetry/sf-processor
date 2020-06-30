# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build -tags exclude_graphdriver_btrfs
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -tags exclude_graphdriver_btrfs
BIN=sfprocessor
OUTPUT=$(BIN)
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
	mkdir -p /usr/local/sf-processor/bin && mkdir -p /usr/local/sf-processor/conf
	cp ./driver/sfprocessor /usr/local/sf-processor/bin/sfprocessor
	cp ./driver/pipeline.json /usr/local/sf-processor/conf/pipeline.json
	cp ./resources/policies/distribution/* /usr/local/sf-processor/conf/

docker-build: build
	sudo docker build -t sf-processor --target=runtime -f Dockerfile .

pull:
	git pull origin master

up: 
	sudo docker-compose -f docker-compose.yml up

