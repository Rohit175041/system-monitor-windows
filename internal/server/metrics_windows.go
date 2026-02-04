package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"system-monitor-windows/internal/config"
	"system-monitor-windows/internal/logger"
)

func Start() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logger.Current)
	})

	addr := fmt.Sprintf(":%d", config.App.HTTPPort)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
