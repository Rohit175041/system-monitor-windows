package logger

func Stop() {
	close(stopCh)
}
