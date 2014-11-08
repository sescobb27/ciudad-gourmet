GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install


all: install
	$(GOBUILD) -v -p=$(GOMAXPROCS) -race ./...
	$(GOBUILD) -ldflags "-w" -o ciudad-gourmet

.PHONY: test open install test-race

test:
	./ciudad-gourmet -seed
	$(GOTEST) -parallel=$(GOMAXPROCS) -v ./...
	./ciudad-gourmet -restore

install:
	$(GOBUILD) -i -p=$(GOMAXPROCS)

test-race:
	./ciudad-gourmet -seed
	$(GOTEST) -parallel=$(GOMAXPROCS) -race -v ./...
	./ciudad-gourmet -restore

open:
	$(shell sudo setcap cap_net_bind_service=+ep `pwd`/start)
