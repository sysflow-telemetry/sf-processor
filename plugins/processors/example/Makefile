GOCMD=go
GOBUILD=$(GOCMD) build -buildmode=plugin
OUTPUT=../../../resources/plugins

.PHONY: all
all: 
	mkdir -p $(OUTPUT); $(GOBUILD) -o $(OUTPUT)/example.so .  
    

