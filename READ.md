system-monitor-windows/
│
├── cmd/
│   └── agent/
│       └── main_windows.go        # App entry point & lifecycle
│
├── internal/
│   ├── collector/
│   │   ├── battery_windows.go     # Windows power API (normalized)
│   │   ├── cpu_windows.go         # CPU usage
│   │   ├── memory_windows.go      # RAM usage
│   │   ├── disk_windows.go        # Disk usage
│   │   └── network_windows.go     # Network throughput
│   │
│   ├── config/
│   │   └── config.go              # JSON config + defaults
│   │
│   ├── logger/
│   │   ├── logger_windows.go      # Metric collection & async enqueue
│   │   ├── async.go               # Buffered channels
│   │   ├── writer.go              # JSON + CSV writer goroutine
│   │   ├── rotate.go              # Size-based log rotation
│   │   ├── helpers.go             # Formatting helpers
│   │   └── stop.go                # Graceful shutdown
│   │
│   ├── tray/
│   │   └── tray_windows.go        # System tray UI
│   │
│   ├── server/
│   │   └── metrics_windows.go     # HTTP /metrics endpoint
│   │
│   ├── alerts/
│   │   └── alerts_windows.go      # Battery & CPU alerts
│   │
│   └── autostart/
│       └── autostart_windows.go   # Windows registry auto-start
│
├── configs/
│   └── config_windows.json        # Runtime configuration
│
├── assets/
│   └── icon.ico                   # Tray icon
│
├── logs/
│   ├── system_metrics.log         # Active JSON log
│   ├── system_metrics.log.1       # Rotated logs
│   ├── system_metrics.csv         # Active CSV log
│   └── system_metrics.csv.1
│
├── build/
│   └── build.ps1                  # Windows build script
│
├── go.mod
├── go.sum
└── README.md



# System Monitor (Go)

• Async logging pipeline (channel-based)
• Size-based log rotation
• HTTP metrics endpoint
• Windows system integration

Architecture:
collector → logger → channel → writer → disk

// compile command
<!--  go build -ldflags="-H=windowsgui" -o system-monitor.exe ./cmd/agent  -->


