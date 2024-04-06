package main

import (
	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/tray"
)

func main() {
	conf := config.NewConfig()
	mc := mouse.NewController(conf)
	trayIcon := tray.NewTray(mc, conf)
	trayIcon.Run()
}
