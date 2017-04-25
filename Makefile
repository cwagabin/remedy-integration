GOTOOLS = github.com/mitchellh/gox
DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES = $(shell go list ./...)

all: deps format build

deps:
	@echo "--> Installing build dependencies"
	@go get -v $(GOTOOLS)
	@go get -d -v ./... $(DEPS)

test: deps
	@go test

format: deps
	@echo "--> Running go fmt"
	@go fmt $(PACKAGES)

build: deps
	@echo "--> Building pd-remedy"
	@mkdir -p bin/
	@gox -output "bin/{{.OS}}_{{.Arch}}/pd-remedy"
