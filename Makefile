
## help: Display help
.PHONY: help
help: Makefile
	@echo "Usage:  make COMMAND"
	@echo
	@echo "Commands:"
	@sed -n 's/^##//p' $< | column -ts ':' |  sed -e 's/^/ /'

## get-deps: Download application dependencies
.PHONY: get-deps
get-deps:
	go mod download


## build: Build application
.PHONY: build
build: get-deps
	go build -ldflags="-s -w" -trimpath -o 'bin/mouse-stay-up' ./cmd/app

## start: Build and start application
.PHONY: start
start: get-deps
	go run ./cmd/app
