package collector

import "github.com/shirou/gopsutil/v3/cpu"

func CPU() float64 {
	usage, _ := cpu.Percent(0, false)
	return usage[0]
}
