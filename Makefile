ROOT_DIR:=$(shell git rev-parse --show-toplevel)
#GLDFLAGS=-X $(REPO)/pkg/version.versionFromGit=$(VERSION_OVERRIDE) -X $(REPO)/pkg/version.commitFromGit=$(HASH) -X $(REPO)/pkg/version.buildDate=$(BUILD_DATE)

all: build
.PHONY: all

$(shell mkdir -p bin)

build: bin/dev-installer

bin/dev-installer: $(GOFILES)
	@echo Building $@
	@go build -o $(ROOT_DIR)/$@ ./cmd/dev-installer/

test-unit:
	@go test -v ./...

verify:
	@go vet $(shell go list ./... | grep -v /vendor/)

verify-deps:
	@echo starting vendor tests
	@go mod tidy
	@go mod vendor
	@go mod verify

clean:
	rm -rf $(ROOT_DIR)/bin

.PHONY: build clean verify-deps
