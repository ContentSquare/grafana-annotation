# git options
GIT_COMMIT ?= $(shell git rev-parse HEAD)
GIT_TAG    ?= $(shell git tag --points-at HEAD)
DIST_TYPE  ?= snapshot
BRANCH     ?= $(shell git rev-parse --abbrev-ref HEAD)

PROJECT_NAME := grafana-annotation
PKG_ORG      := Contentsquare
PKG 		 := grafana-annotation
PKG_LIST 	 := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES 	 := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
BINARY_NAME  := grafana-annotation

M = $(shell printf "\033[34;1m▶\033[0m")

GO			 := go
GOFMT		 := gofmt
OS           := $(shell uname -s)
GOOS         ?= $(shell echo $(OS) | tr '[:upper:]' '[:lower:]')
GOARCH		 ?= amd64

BUILD_CONSTS = -X main.buildRevision=$(GIT_COMMIT) -X main.buildTag=$(GIT_TAG)
BUILD_OPTS = -ldflags="$(BUILD_CONSTS)" -gcflags="-trimpath=$(GOPATH)/src"

tidy:
	$(info =====  $@  =====)
	GO111MODULE=on go mod tidy

deps:
	$(info =====  $@  =====)
	GO111MODULE=on go mod vendor

format:
	$(info =====  $@  =====)
	$(GOFMT) -w -s $(GO_FILES)

build:
	$(info =====  $@  =====)
	GO111MODULE=on GOOS=$(GOOS) $(GO) build -tags static -o $(BINARY_NAME) $(BUILD_OPTS)

test:
	$(info =====  $@  =====)
	$(GO) test -v -race -cover -coverprofile=coverage.out ./...

fmt:
	$(info =====  gofmt =====)
	$(GOFMT) -d -e -s $(GO_FILES)

#lint:
#	$(info =====  $@  =====)
#	$(GO) vet $(PKG_LIST)
#	$(GO) list ./... | grep -Ev /vendor/ | sed -e 's/${PROJECT_NAME}/./' | xargs -L1 golint -set_exit_status
#

lint:
	$(info =====  $@  =====)
	golangci-lint run . -v

tarball: version
	$(info =====  $@  =====)
ifneq ($(GIT_TAG),)
	tar czf $(BINARY_NAME)-$(GOOS)-$(GOARCH)-$(GIT_TAG).tar.gz $(BINARY_NAME)
else
	tar czf $(BINARY_NAME)-$(GOOS)-$(GOARCH)-$(VERSION_FILE).tar.gz $(BINARY_NAME)
endif

version:
	$(info =====  $@  =====)
ifneq ($(GIT_TAG),)
	$(eval VERSION := $(GIT_TAG))
	$(eval VERSION_FILE := $(GIT_TAG))
else
	$(eval VERSION := $(subst /,-,$(BRANCH)))
	$(eval VERSION_FILE := $(GIT_COMMIT)-SNAPSHOT)
endif
	@test -n "$(VERSION)"
	$(info Building $(VERSION)/$(VERSION_FILE) on sha1 $(GIT_COMMIT))

get_version: version
	$(info $(VERSION_FILE))

.PHONY: tidy \
		deps \
		format \
		build \
		test \
		fmt \
		lint \
		tarball