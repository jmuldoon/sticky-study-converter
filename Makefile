SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 0.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS_BY_PATH = $(shell go list ./... | grep -v /vendor/)
PKGS = $(shell glide novendor)

.PHONY: all build clean install uninstall fmt imports check test coverage run help coverageall
.PHONY: tools varcheck structcheck aligncheck deadcode errcheck checkall testverbose deps doc
.PHONY: depsupdate

all: version tools deps checkall install

$(TARGET): deps $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    clean              Remove binaries, artifacts and releases.'
	@echo '    doc                Start Go documentation server on port 8080.'
	@echo '    tools              Install tools needed by the project.'
	@echo '    deps               Glide download and install build time dependencies.'
	@echo '    depsupdate         Glide update the dependencies.'
	@echo '    check              Runs go fmt, lint, vet, imports.'
	@echo '    checkall          	Runs go fmt, lint, vet, imports, deadcode, varcheck, errcheck, structcheck, aligncheck.'
	@echo '    test               Run unit tests, and check.'
	@echo '    testverbose        Run unit tests in verbose mode, and check.'
	@echo '    coverage           Report code tests coverage, and check.'
	@echo '    coverageall        Report code tests coverage, and run checkall.'
	@echo '    build              Build project for current platform.'
	@echo '    fmt                Run go fmt.'
	@echo '    errcheck           Finds unchecked errors in a go programs.'
	@echo '    varcheck           Finds unused global variables and constants.'
	@echo '    structcheck        Find unused struct fields.'
	@echo '    aligncheck         Find inefficiently packed structs.'
	@echo '    deadcode           Find unused declarations.'
	@echo '    imports            Run goimports to remove/add (un)necessary imports.'
	@echo '    install            Run go install'
	@echo '    uninstall          Force removes the artifact'
	@echo '    version 	          Checks the go version'
	@echo ''
	@echo 'Targets run by default are: checkall, install'
	@echo ''

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

doc:
	godoc -http=:8080 -index

install: deps
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

deps:
	@glide install

depsupdate:
	@glide up

tools:
	@go get github.com/Masterminds/glide
	@go get golang.org/x/tools/cmd/goimports
	@go get github.com/golang/lint/golint
	@go get github.com/kisielk/errcheck
	@go get github.com/opennota/check/cmd/aligncheck
	@go get github.com/opennota/check/cmd/structcheck
	@go get github.com/opennota/check/cmd/varcheck
	@go get github.com/remyoudompheng/go-misc/deadcode

errcheck:
	@errcheck $(PKGS)

varcheck:
	@varcheck $(PKGS)

structcheck:
	@structcheck $(PKGS)

aligncheck:
	@aligncheck $(PKGS)

deadcode:
	@deadcode -test $(PKGS_BY_PATH)

imports:
	@goimports -l -w .

check: imports
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $(PKGS); do golint $${d}; done
	@go tool vet ${SRC}

checkall: imports errcheck varcheck structcheck aligncheck deadcode
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $(PKGS); do golint $${d}; done
	@go tool vet ${SRC}

test: check
	@go test $(PKGS)

testverbose: check
	@go test -v $(PKGS)

coverage: check
	@go test -cover $(PKGS)

coverageall: checkall
	@go test -cover $(PKGS)

version:
	@go version