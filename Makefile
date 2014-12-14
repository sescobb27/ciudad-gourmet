GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test


all:
	$(GOBUILD) -i -p=$(GOMAXPROCS)
	$(GOBUILD) -ldflags "-w" -v -p=$(GOMAXPROCS) ./...

.PHONY: test open install test-race clean rsa

test:
	./ciudad-gourmet -seed
	$(GOTEST) -parallel=$(GOMAXPROCS) -v ./...
	./ciudad-gourmet -restore

test-race:
	./ciudad-gourmet -seed
	$(GOTEST) -parallel=$(GOMAXPROCS) -race -v ./...
	./ciudad-gourmet -restore

clean:
	rm -rf ciudad-gourmet.log-*
	rm -rf handlers/ciudad-gourmet.log-*

rsa:
	openssl genrsa -out cg.rsa 4096
	openssl rsa -in cg.rsa -pubout > cg.rsa.pub


open:
	$(shell sudo setcap cap_net_bind_service=+ep `pwd`/start)
