GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install


all:
	$(GOBUILD) -i -v -p=$(GOMAXPROCS) -race ./...
	$(GOBUILD) -o ciudad-gourmet

.PHONY: test open

test:
	$(GOTEST) -parallel=$(GOMAXPROCS) -v ./...

open:
	$(shell sudo setcap cap_net_bind_service=+ep `pwd`/start)
