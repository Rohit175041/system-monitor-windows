package collector

import "github.com/shirou/gopsutil/v3/disk"

func Disk() float64 {
	d, _ := disk.Usage("C:")
	return d.UsedPercent
}
