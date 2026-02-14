package utils

import (
	"errors"
	"os/exec"
	"runtime"
)

// OpenWebPage opens the given URL in the user's default browser.
func OpenWebPage(url string) error {
	var cmd string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	default:
		return errors.New("utils: unsupported platform")
	}
	return exec.Command(cmd, url).Start() //nolint:gosec // cmd is a fixed platform command, not user input
}
