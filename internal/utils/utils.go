package utils

import (
	"os/exec"
	"runtime"
)

func OpenWebPage(url string) {
	switch runtime.GOOS {
	case "darwin":
		exec.Command("open", url).Start()
	case "linux":
		exec.Command("xdg-open", url).Start()
	default:
		panic("unsupported platform")
	}
}
