GOCMD=go
GOBUILD=$(GOCMD) build -buildmode=plugin -tags exclude_graphdriver_btrfs
OUTPUT=../../../resources/actions

.PHONY: all
all: 
	mkdir -p $(OUTPUT); $(GOBUILD) -o $(OUTPUT)/now.so .  
    

