GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
MODELS_DIR=models
HANDLERS_DIR=handlers
DB_DIR=db
HELPERS_DIR=helpers

all:
	$(GOBUILD) -i $(HELPERS_DIR) \
	$(MODELS_DIR) \
	$(HANDLERS_DIR) \
	$(DB_DIR)
	$(GOBUILD) -o ciudad-gourmet

.PHONY: test open install

test:
	$(GOTEST) $(MODELS_DIR)
	$(GOTEST) $(HANDLERS_DIR)
	$(GOTEST) $(HELPERS_DIR)

open:
	$(shell sudo setcap cap_net_bind_service=+ep `pwd`/start)

install:
	$(MAKE) -C $(MODELS_DIR) install
	$(MAKE) -C $(HANDLERS_DIR) install
	$(MAKE) -C $(DB_DIR) install
