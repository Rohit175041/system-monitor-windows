package alerts

import (
	"fmt"

	"system-monitor-windows/internal/config"
	"system-monitor-windows/internal/model"
)

func Check(m model.Metrics) {
	// Battery alert
	if m.Battery.Percent <= config.App.BatteryAlertPercent &&
		m.Battery.State == "discharging" {

		fmt.Println("⚠️ ALERT: Low Battery",
			m.Battery.Percent, "%")
	}

	// CPU alert
	if m.CPU.Usage >= float64(config.App.CPUAlertPercent) {
		fmt.Println("⚠️ ALERT: High CPU Usage",
			m.CPU.Usage, "%")
	}
}
