GOCMD=go
GOBUILD=$(GOCMD) build -buildmode=plugin
OUTPUT=../../../resources/handlers

.PHONY: all
all: 
	mkdir -p $(OUTPUT); $(GOBUILD) -o $(OUTPUT)/printer.so .  
    

