package logger

import (
	"math"
	"syscall"
	"time"
	"unsafe"

	"system-monitor-windows/internal/alerts"
	"system-monitor-windows/internal/collector"
	"system-monitor-windows/internal/config"
	"system-monitor-windows/internal/model"
)

var Current model.Metrics

// -------- Windows-safe hostname --------
func windowsHostname() string {
	var size uint32 = 256
	buf := make([]uint16, size)

	if err := syscall.GetComputerName(
		(*uint16)(unsafe.Pointer(&buf[0])),
		&size,
	); err != nil {
		return "UNKNOWN"
	}
	return syscall.UTF16ToString(buf[:size])
}

func Start() {
	host := windowsHostname()

	ticker := time.NewTicker(
		time.Duration(config.App.LogIntervalSec) * time.Second,
	)
	defer ticker.Stop()

	for range ticker.C {
		battery, charging := collector.Battery()
		downKB, upKB := collector.Network()

		state := "discharging"
		if charging {
			state = "charging"
		}

		m := model.Metrics{
			Timestamp: time.Now().Format(time.RFC3339),
			Level:     "INFO",
			App:       "system-monitor",
			Host:      host,
		}

		m.Battery.Percent = battery
		m.Battery.State = state
		m.CPU.Usage = math.Round(collector.CPU()*100) / 100
		m.Memory.Usage = math.Round(collector.Memory())
		m.Disk.Usage = math.Round(collector.Disk()*100) / 100
		m.Network.DownKB = downKB
		m.Network.UpKB = upKB
		m.Network.TotalKB = downKB + upKB

		Current = m
		alerts.Check(m)

		// -------- ASYNC PUSH (NON-BLOCKING) --------
		select {
		case logCh <- m:
		default:
			// channel full → drop (intentional)
		}
	}
}
