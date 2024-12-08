# Mouse-Stay-Up

[![ci status](https://github.com/sonjek/mouse-stay-up/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sonjek/mouse-stay-up/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/sonjek/mouse-stay-up)](https://goreportcard.com/report/github.com/sonjek/mouse-stay-up) [![Contributors](https://img.shields.io/github/contributors/sonjek/mouse-stay-up)](https://github.com/sonjek/mouse-stay-up/graphs/contributors) [![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/sonjek/mouse-stay-up?include_prereleases)](https://github.com/sonjek/mouse-stay-up/releases) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/sonjek/mouse-stay-up/blob/master/LICENSE)


This lightweight application is designed to prevent your computer from entering sleep mode by periodically moving the cursor when it detects periods of inactivity.
Additionally, the program allows you to disable the keyboard programmatically (MacOS only).

## Installation from source

### Prepare Build Environment (macOS)

For macOS systems, the installation of build essential tools and development libraries is not needed.

### Prepare Build Environment (Debian-Based Linux)

To build the application on Debian-based Linux systems, follow these steps to install the necessary dependencies:

1. **Install Build Essential Tools**:
   Install the essential tools required for building software on Debian-based systems:

    ```bash
    sudo apt install build-essential
    ```

2. **Install Development Libraries**:

    - **libx11-dev**:
      Install the development files for the X11 library, which is required for graphical applications:

        ```bash
        sudo apt install libx11-dev
        ```

    - **libayatana-appindicator3-dev**:
      Install the development files for the Ayatana AppIndicator library, which is used for creating application indicators:

        ```bash
        sudo apt install libayatana-appindicator3-dev
        ```


### Build the Application


1. Verify that you have `Go 1.22+` installed
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
  format        Fix code format issues
  deadcode      Run deadcode tool for find unreachable functions
  audit         Quality checks
  help          Display this help
```
