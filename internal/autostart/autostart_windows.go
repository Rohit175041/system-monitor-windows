package autostart

import (
	"os"

	"golang.org/x/sys/windows/registry"
)

func Enable() {
	exe, _ := os.Executable()
	key, _, _ := registry.CreateKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.SET_VALUE,
	)
	key.SetStringValue("SystemMonitor", exe)
	key.Close()
}
