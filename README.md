# Mouse-Stay-Up

[![ci status](https://github.com/sonjek/mouse-stay-up/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sonjek/mouse-stay-up/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/sonjek/mouse-stay-up)](https://goreportcard.com/report/github.com/sonjek/mouse-stay-up) [![Contributors](https://img.shields.io/github/contributors/sonjek/mouse-stay-up)](https://github.com/sonjek/mouse-stay-up/graphs/contributors) ![Go](https://img.shields.io/github/go-mod/go-version/sonjek/mouse-stay-up?label=go) [![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/sonjek/mouse-stay-up?include_prereleases)](https://github.com/sonjek/mouse-stay-up/releases) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/sonjek/mouse-stay-up/blob/master/LICENSE)

This lightweight application is designed to prevent your computer from entering sleep mode by periodically moving the cursor when it detects periods of inactivity.
Additionally, the program allows you to disable the keyboard programmatically (MacOS only).

## Installation from source

### Prepare Build Environment (macOS)

For macOS systems, the installation of additional tools and development libraries is not needed.

### Prepare Build Environment (Linux/BSD)

You may require a proxy app which can convert the new DBus calls to the old format.
More info in [systray](https://github.com/fyne-io/systray?tab=readme-ov-file#linuxbsd) library info.
For Debian-based systems with GTK based trays, use [this](https://gist.github.com/archisman-panigrahi/cd571ddea1aa2c5e2b4fa7bcbee7d5df) script to install **snixembed**.


### Install

Verify that you have `Go 1.26+` installed. If `go` is not installed, follow instructions on the [Go website](https://go.dev/doc/install).

#### Install via go install

```sh
go install github.com/sonjek/mouse-stay-up/cmd/mouse-stay-up@latest
```

This should download the source code and compile the executable into `$GOPATH/bin/mouse-stay-up`.

Make sure `$GOPATH/bin` is in your `$PATH` so the shell can discover this application.

For example, my `~/.profile` contains this:

```sh
export GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"
```

Alternatively, you can run this application using the full path `/Users/<USERNAME>/go/bin/mouse-stay-up`.

#### Build yourself

1. Clone this repository

```sh
git clone https://github.com/sonjek/mouse-stay-up && cd mouse-stay-up
```

2. Build

```sh
make build
```

The binary file is built and ready to run:

```sh
./bin/mouse-stay-up
```

3. You can install the binary file to your OS.

   Installs to `/usr/local/bin/` by default:

```sh
make install
```

or

```sh
sudo make install
```

Install to a custom location:

```sh
make install INSTALL_PATH=/tmp
```

### MacOS

#### Autostart at login

1. Create a new file in `~/Library/LaunchAgents/` called `com.github.sonjek.mouse-stay-up.plist` with the following contents:

```sh
cat <<EOF > ~/Library/LaunchAgents/com.github.sonjek.mouse-stay-up.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.sonjek.mouse-stay-up</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/mouse-stay-up</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <false/>
    <key>StandardOutPath</key>
    <string>/tmp/mouse-stay-up.stdout.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/mouse-stay-up.stderr.log</string>
</dict>
</plist>
EOF
```
or

```sh
cp deployments/com.github.sonjek.mouse-stay-up.plist ~/Library/LaunchAgents/com.github.sonjek.mouse-stay-up.plist
```

If you used `go install`, update the application path to `/Users/<USERNAME>/go/bin/mouse-stay-up`.

2. Load the launch agent:

```sh
launchctl load ~/Library/LaunchAgents/com.github.sonjek.mouse-stay-up.plist
```

#### Access rights

1. Accessibility permissions are required for the keyboard lock to function. Please grant these permissions in settings to enable this feature:
- Go to `System Settings > Privacy & Security > Accessibility`.
- Enable `mouse-stay-up`.

2. To disable background activity:
- Go to `System Settings > General > Login Items & Extensions`.
- Disable `mouse-stay-up` in app background activity section.
