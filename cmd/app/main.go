package main

import (
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/tray"
)

const (
	GitRepo string = "https://github.com/sonjek/mouse-stay-up"
)

func main() {
	mc := mouse.NewController(GitRepo)
	trayIcon := tray.NewTray(mc)
	trayIcon.Run()
}
