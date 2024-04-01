# Mouse-Stay-Up
This is a lightweight application designed to keep your computer awake by moving the cursor when it detects periods of inactivity.


## Installation from source

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

    The binary file is built and ready to run from current folder:
    ```
    $ ./bin/mouse-stay-up
    ```

4. You can install the binary file to your OS.

   Installs to `/usr/local/bin/` by default:
    ```
    $ make install
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
  help        Display this help
  get-deps    Download application dependencies
  build       Build application
  clean       Remove binary file from local bin directory
  install     Install binary file from local bin directory to /usr/local/bin/
  uninstall   Remove binary file from /usr/local/bin/
  start       Build and start application
```
