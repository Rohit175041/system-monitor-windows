package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	// Logger
	LogIntervalSec int    `json:"log_interval_sec"`
	JSONLogFile    string `json:"json_log_file"`
	CSVLogFile     string `json:"csv_log_file"`

	// Log rotation
	LogMaxSizeMB  int `json:"log_max_size_mb"`
	LogMaxBackups int `json:"log_max_backups"`

	// Tray
	TrayRefreshSec int `json:"tray_refresh_sec"`

	// Alerts
	BatteryAlertPercent int `json:"battery_alert_percent"`
	CPUAlertPercent     int `json:"cpu_alert_percent"`

	// Server
	HTTPPort int `json:"http_port"`
}

var App Config

func Load() error {
	data, err := os.ReadFile("configs/config_windows.json")
	if err != nil {
		return err
	}

	// Load JSON into struct
	if err := json.Unmarshal(data, &App); err != nil {
		return err
	}

	// ---- Defaults (safety) ----
	if App.LogIntervalSec <= 0 {
		App.LogIntervalSec = 1
	}
	if App.TrayRefreshSec <= 0 {
		App.TrayRefreshSec = 2
	}
	if App.JSONLogFile == "" {
		App.JSONLogFile = "logs/system_metrics.log"
	}
	if App.CSVLogFile == "" {
		App.CSVLogFile = "logs/system_metrics.csv"
	}

	// Log rotation defaults
	if App.LogMaxSizeMB <= 0 {
		App.LogMaxSizeMB = 500
	}
	if App.LogMaxBackups <= 0 {
		App.LogMaxBackups = 5
	}

	if App.HTTPPort == 0 {
		App.HTTPPort = 8080
	}

	return nil
}
