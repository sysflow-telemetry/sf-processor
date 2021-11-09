GOCMD=go
GOBUILD=$(GOCMD) build -buildmode=plugin
OUTPUT=../../../resources/actions

.PHONY: all
all: 
	mkdir -p $(OUTPUT); $(GOBUILD) -o $(OUTPUT)/now.so .  
    

