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

build: 
	cd $(SRC) && $(GOBUILD) -o $(OUTPUT) -v

deps:
	cd $(SRC) && $(GOGET) ./..

test:
	cd $(SRC) && $(GOTEST) -v ./...

clean:
	cd $(SRC) && $(GOCLEAN)
	rm -f $(SRC)/$(BIN)

docker-build: build
	sudo docker build -t sf-processor:latest -f Dockerfile.processor .

pull:
	git pull origin master

up: 
	sudo docker-compose -f docker-compose.processor.yml up

