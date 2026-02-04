package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"system-monitor-windows/internal/config"
	"system-monitor-windows/internal/logger"
	"system-monitor-windows/internal/server"
	"system-monitor-windows/internal/tray"
)

func main() {
	// Load JSON config
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	log.Println("System Monitor starting...")

	// 🔹 Start async log writer FIRST
	logger.StartWriter()

	// 🔹 Start background services
	go logger.Start()
	go server.Start()

	// 🔹 Handle OS shutdown (Ctrl+C, taskkill)
	go handleShutdown()

	// 🔹 Start tray (BLOCKING)
	tray.Start()

	// Tray exited → clean shutdown
	logger.Stop()
	log.Println("System Monitor stopped")
}

// Graceful shutdown handler
func handleShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	logger.Stop()
	os.Exit(0)
}
