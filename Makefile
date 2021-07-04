VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
VERSION_BUILD ?= 0
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)

ORG := github.com
OWNER := kairen
PROJECT_NAME := elasticsearch-toolbox
REPO_PATH ?= $(ORG)/$(OWNER)/$(PROJECT_NAME)

GOOS ?= $(shell go env GOOS)

$(shell mkdir -p ./out)

.PHONY: all 
all: build 

.PHONY: build
build: out/elasticsearch-toolbox

.PHONY: out/elasticsearch-toolbox
out/elasticsearch-toolbox:
	GOOS=$(GOOS) CGO_ENABLED=0 go build \
	 -ldflags="-s -w -X $(REPO_PATH)/pkg/version.version=$(VERSION)" \
	 -a -o $@ cmd/main.go

.PHONY: image-build
image-build: 
	@docker build -t quay.io/$(OWNER)/$(PROJECT_NAME):$(VERSION) .

.PHONY: image-push 
image-push:
	@docker push quay.io/$(OWNER)/$(PROJECT_NAME):$(VERSION)