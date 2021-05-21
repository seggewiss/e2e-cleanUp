.DEFAULT_GOAL := help

build:		## Build binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o proxy -v cmd/proxy/main.go
fmt:		## Format source code
	go fmt ./...
help:
	@echo
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target> env=<environment> \033[36m\033[0m\n\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)