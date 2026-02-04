package collector

import "github.com/shirou/gopsutil/v3/mem"

func Memory() float64 {
	m, _ := mem.VirtualMemory()
	return m.UsedPercent
}
