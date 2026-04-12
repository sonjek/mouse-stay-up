package main

import (
	"context"

	"github.com/sonjek/mouse-stay-up/internal/config"
	"github.com/sonjek/mouse-stay-up/internal/mouse"
	"github.com/sonjek/mouse-stay-up/internal/tray"
	"github.com/sonjek/mouse-stay-up/pkg/keyboard"
)

func main() {
	ctx := context.Background()
	conf := config.NewConfig()
	mc := mouse.NewController(conf)
	kc := keyboard.NewController()
	trayIcon := tray.NewTray(ctx, mc, kc, conf)
	trayIcon.Run()
}
