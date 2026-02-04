package collector

import (
	"syscall"
	"unsafe"
)

type systemPowerStatus struct {
	ACLineStatus        byte
	BatteryFlag         byte
	BatteryLifePercent  byte
	SystemStatusFlag    byte
	BatteryLifeTime     uint32
	BatteryFullLifeTime uint32
}

var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procGetSystemPowerStat = kernel32.NewProc("GetSystemPowerStatus")
)

// Battery returns battery percentage and charging state (Windows-safe)
func Battery() (percent int, charging bool) {
	var s systemPowerStatus

	r1, _, _ := procGetSystemPowerStat.Call(
		uintptr(unsafe.Pointer(&s)),
	)
	if r1 == 0 {
		return 0, false
	}

	charging = s.ACLineStatus == 1
	percent = int(s.BatteryLifePercent)

	// ---------- WINDOWS QUIRK FIX ----------
	// Some laptops report 0–1% while charging
	if charging && percent <= 1 {
		percent = 5
	}

	// ---------- SAFETY CLAMPS ----------
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	return percent, charging
}
