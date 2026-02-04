package logger

import "system-monitor-windows/internal/model"

var (
	logCh  = make(chan model.Metrics, 256) // buffered
	stopCh = make(chan struct{})
)
