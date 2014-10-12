GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install


all:
	$(GOBUILD) -i -v -p=$(GOMAXPROCS) -race ./...
	$(GOBUILD) -o ciudad-gourmet

.PHONY: test open

test:
	./ciudad-gourmet -seed
	$(GOTEST) -parallel=$(GOMAXPROCS) -v ./...
	./ciudad-gourmet -restore
open:
	$(shell sudo setcap cap_net_bind_service=+ep `pwd`/start)
