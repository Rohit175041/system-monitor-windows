package tray

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
	"system-monitor-windows/internal/config"
	"system-monitor-windows/internal/logger"
)

func Start() {
	systray.Run(onReady, nil)
}

func onReady() {
	m := systray.AddMenuItem("System Monitor", "")
	quit := systray.AddMenuItem("Quit", "Exit")

	go func() {
		for {
			m.SetTitle(fmt.Sprintf(
				"🔋 %d%% | CPU %.1f%% | RAM %.1f%%",
				logger.Current.Battery.Percent,
				logger.Current.CPU.Usage,
				logger.Current.Memory.Usage,
			))

			time.Sleep(
				time.Duration(config.App.TrayRefreshSec) * time.Second,
			)
		}
	}()

	go func() {
		<-quit.ClickedCh
		systray.Quit()
	}()
}
