package main

import (
	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/tray"
)

func main() {
	config := config.NewConfig()
	mc := mouse.NewController(config)
	trayIcon := tray.NewTray(mc, config)
	trayIcon.Run()
}
