BINDIR := $(CURDIR)/bin
BINNAME := mouse-stay-up
TARGET_BIN := $(BINDIR)/$(BINNAME)
INSTALL_PATH := /usr/local/bin

GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6.1

# -------------------------------------------------------------------------------------------------
# main
# -------------------------------------------------------------------------------------------------

all: help

## build: Build application
.PHONY: build
build: get-deps
	go build -ldflags="-s -w" -trimpath -o '$(TARGET_BIN)' ./cmd/app
	@echo "Binary file located in '$(TARGET_BIN)'"

## clean: Remove binary file from local bin directory
.PHONY: clean
clean:
	-@rm -rf $(BINDIR)

## install: Install binary file from local bin directory to /usr/local/bin/
.PHONY: install
install:
	@test -e "$(BINDIR)/$(BINNAME)" &> /dev/null || (echo "There are no executable file. Please run 'make build'" && false)
	@install "$(TARGET_BIN)" "$(INSTALL_PATH)/$(BINNAME)"
	-@$(MAKE) clean
	@echo "Binary file installed to '$(INSTALL_PATH)/$(BINNAME)'"

## uninstall: Remove binary file from /usr/local/bin/
.PHONY: uninstall
uninstall: clean
	@echo "Removing: $(INSTALL_PATH)/$(BINNAME)"
	@rm -f "$(INSTALL_PATH)/$(BINNAME)"

## start: Build and start application
.PHONY: start
start: get-deps
	go run ./cmd/app

# -------------------------------------------------------------------------------------------------
# testing
# -------------------------------------------------------------------------------------------------

## test: Run unit tests
.PHONY: test
test: check-go
	@go test -v -count=1 ./...

# -------------------------------------------------------------------------------------------------
# tools && shared
# -------------------------------------------------------------------------------------------------

## check-go: Ensure that Go is installed
.PHONY: check-go
check-go:
	@command -v go &> /dev/null || (echo "Please install GoLang" && false)

## tidy: Removes unused dependencies and adds missing ones
.PHONY: tidy
tidy: check-go
	go mod tidy

## update-deps: Update go dependencies
.PHONY: update-deps
update-deps: check-go
	go get -u ./...
	-@$(MAKE) tidy

## get-deps: Download application dependencies
.PHONY: get-deps
get-deps: check-go
	go mod download

## lint: Run golangci-lint to lint go files
.PHONY: lint
lint:
	go run $(GOLANGCI_LINT_PACKAGE) run

## lint-fix: Run golangci-lint to lint go files and fix issues
.PHONY: lint-fix
lint-fix:
	go run $(GOLANGCI_LINT_PACKAGE) run --fix

## lint-fmt: Run golangci-lint fmt to show code format issues
.PHONY: lint-fmt
lint-fmt:
	go run $(GOLANGCI_LINT_PACKAGE) fmt

## deadcode: Run deadcode tool for find unreachable functions
.PHONY: deadcode
deadcode:
	go run golang.org/x/tools/cmd/deadcode@latest -test ./...

## audit: Quality checks
.PHONY: audit
audit: check-go
	go mod verify
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## help: Display this help
.PHONY: help
help: Makefile
	@echo "Usage:  make COMMAND"
	@echo
	@echo "Commands:"
	@sed -n 's/^##//p' $< | column -ts ':' |  sed -e 's/^/ /'
