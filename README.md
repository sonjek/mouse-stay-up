# Mouse-Stay-Up

[![ci status](https://github.com/sonjek/mouse-stay-up/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sonjek/mouse-stay-up/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/sonjek/mouse-stay-up)](https://goreportcard.com/report/github.com/sonjek/mouse-stay-up) [![Contributors](https://img.shields.io/github/contributors/sonjek/mouse-stay-up)](https://github.com/sonjek/mouse-stay-up/graphs/contributors) [![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/sonjek/mouse-stay-up?include_prereleases)](https://github.com/sonjek/mouse-stay-up/releases) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/sonjek/mouse-stay-up/blob/master/LICENSE)


This lightweight application is designed to prevent your computer from entering sleep mode by periodically moving the cursor when it detects periods of inactivity.
Additionally, the program allows you to disable the keyboard programmatically (MacOS only).

## Installation from source

### Prepare Build Environment (macOS)

For macOS systems, the installation of additional tools and development libraries is not needed.

### Prepare Build Environment (Linux/BSD)

You may require a proxy app which can convert the new DBus calls to the old format.
More info in [systray](https://github.com/fyne-io/systray?tab=readme-ov-file#linuxbsd) library info.
For Debian-based systems with GTK based trays, use [this](https://gist.github.com/archisman-panigrahi/cd571ddea1aa2c5e2b4fa7bcbee7d5df) script to install **snixembed**.


### Build the Application


1. Verify that you have `Go 1.24+` installed
   ```sh
   $ go version
   ```

   If `go` is not installed, follow instructions on the [Go website](https://golang.org/doc/install).

2. Clone this repository
   ```sh
   $ git clone https://github.com/sonjek/mouse-stay-up
   $ cd mouse-stay-up
   ```

3. Build
    ```sh
    $ make build
    ```

    The binary file is built and ready to run:
    ```
    $ ./bin/mouse-stay-up
    ```

4. You can install the binary file to your OS.

   Installs to `/usr/local/bin/` by default:
    ```
    $ make install
    or
    $ sudo make install
    ```

   Install to a different location:
    ```
    $ make install INSTALL_PATH=/tmp
    ```

All available makefile actions:
```sh
% make
Usage:  make COMMAND

Commands:
  build         Build application
  clean         Remove binary file from local bin directory
  install       Install binary file from local bin directory to /usr/local/bin/
  uninstall     Remove binary file from /usr/local/bin/
  start         Build and start application
  test          Run unit tests
  check-go      Ensure that Go is installed
  tidy          Removes unused dependencies and adds missing ones
  update-deps   Update go dependencies
  get-deps      Download application dependencies
  lint          Run golangci-lint to lint go files
  lint-fix      Run golangci-lint to lint go files and fix issues
  lint-fmt      Run golangci-lint fmt to show code format issues
  format        Run gofumpt tool to fix code format issues
  deadcode      Run deadcode tool for find unreachable functions
  audit         Quality checks
  help          Display this help
```
