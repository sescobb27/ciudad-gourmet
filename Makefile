GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
MODELS_DIR=src/models
HANDLERS_DIR=src/handlers
DB_DIR=src/db

all: model handler database
	${GOBUILD}

model:
	$(MAKE) -C $(MODELS_DIR)

handler:
	$(MAKE) -C $(HANDLERS_DIR)

database:
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
