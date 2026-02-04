package collector

import "github.com/shirou/gopsutil/v3/net"

var lastSent, lastRecv uint64

func Network() (downKB, upKB uint64) {
	io, _ := net.IOCounters(false)
	downKB = (io[0].BytesRecv - lastRecv) / 1024
	upKB = (io[0].BytesSent - lastSent) / 1024
	lastRecv = io[0].BytesRecv
	lastSent = io[0].BytesSent
	return
}
