GOCMD=go
GOBUILD=$(GOCMD) build -buildmode=plugin -tags exclude_graphdriver_btrfs
OUTPUT=../../../resources/actions
INSTALL=/usr/local/sysflow/resources/actions

.PHONY: all
all:
	mkdir -p $(OUTPUT); $(GOBUILD) -o $(OUTPUT)/now.so .  

.PHONY: install
install: all
	mkdir -p $(INSTALL)
	cp $(OUTPUT)/now.so $(INSTALL)
    

