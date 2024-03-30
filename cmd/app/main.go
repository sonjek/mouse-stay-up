package main

import (
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/tray"
)

func main() {
	mc := mouse.NewController()
	trayIcon := tray.NewTray(mc)
	trayIcon.Run()
}
