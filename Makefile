BINDIR := $(CURDIR)/bin
BINNAME := mouse-stay-up
TARGET_BIN := $(BINDIR)/$(BINNAME)
INSTALL_PATH := /usr/local/bin


## help: Display this help
.PHONY: help
help: Makefile
	@echo "Usage:  make COMMAND"
	@echo
	@echo "Commands:"
	@sed -n 's/^##//p' $< | column -ts ':' |  sed -e 's/^/ /'

## get-deps: Download application dependencies
.PHONY: get-deps
get-deps:
	@command -v go &> /dev/null || (echo "Please install GoLang" && false)
	go mod download


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
