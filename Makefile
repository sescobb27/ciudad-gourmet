GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
MODELS_DIR=models
HANDLERS_DIR=handlers
DB_DIR=db

all: model handler db
	${GOBUILD}

model:
	$(MAKE) -C $(MODELS_DIR)

handler:
	$(MAKE) -C $(HANDLERS_DIR)

db:
	$(MAKE) -C $(DB_DIR)

.PHONY: test open install

test:
	$(MAKE) -C $(MODELS_DIR) test
	$(MAKE) -C $(HANDLERS_DIR) test
	$(MAKE) -C $(DB_DIR) test

open:
	$(shell sudo setcap cap_net_bind_service=+ep `pwd`/start)

install:
	$(MAKE) -C $(MODELS_DIR) install
	$(MAKE) -C $(HANDLERS_DIR) install
	$(MAKE) -C $(DB_DIR) install
