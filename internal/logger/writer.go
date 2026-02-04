// package logger

// import (
// 	"bufio"
// 	"encoding/csv"
// 	"encoding/json"
// 	"os"
// 	"strconv"

// 	"system-monitor-windows/internal/config"
// )

// func StartWriter() {
// 	// ---------- JSON ----------
// 	jsonFile, err := os.OpenFile(
// 		config.App.JSONLogFile,
// 		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
// 		0644,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	jsonWriter := bufio.NewWriterSize(jsonFile, 32*1024)

// 	// ---------- CSV ----------
// 	csvFile, err := os.OpenFile(
// 		config.App.CSVLogFile,
// 		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
// 		0644,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	csvBuf := bufio.NewWriterSize(csvFile, 32*1024)
// 	csvWriter := csv.NewWriter(csvBuf)

// 	// CSV header once
// 	if stat, _ := csvFile.Stat(); stat.Size() == 0 {
// 		csvWriter.Write([]string{
// 			"timestamp",
// 			"host",
// 			"battery_percent",
// 			"battery_state",
// 			"cpu_usage_percent",
// 			"memory_usage_percent",
// 			"disk_usage_percent",
// 			"net_down_kbps",
// 			"net_up_kbps",
// 			"net_total_kbps",
// 		})
// 		csvWriter.Flush()
// 		csvBuf.Flush()
// 	}

// 	go func() {
// 		defer func() {
// 			jsonWriter.Flush()
// 			csvWriter.Flush()
// 			csvBuf.Flush()
// 			jsonFile.Sync()
// 			csvFile.Sync()
// 			jsonFile.Close()
// 			csvFile.Close()
// 		}()

// 		for {
// 			select {
// 			case m := <-logCh:
// 				// JSON
// 				if b, err := json.Marshal(m); err == nil {
// 					jsonWriter.Write(b)
// 					jsonWriter.WriteByte('\n')
// 				}

// 				// CSV
// 				csvWriter.Write([]string{
// 					m.Timestamp,
// 					m.Host,
// 					strconv.Itoa(m.Battery.Percent),
// 					m.Battery.State,
// 					formatFloat(m.CPU.Usage),
// 					formatFloat(m.Memory.Usage),
// 					formatFloat(m.Disk.Usage),
// 					strconv.FormatUint(m.Network.DownKB, 10),
// 					strconv.FormatUint(m.Network.UpKB, 10),
// 					strconv.FormatUint(m.Network.TotalKB, 10),
// 				})

// 				jsonWriter.Flush()
// 				csvWriter.Flush()
// 				csvBuf.Flush()

// 			case <-stopCh:
// 				return
// 			}
// 		}
// 	}()
// }


package logger

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"system-monitor-windows/internal/config"
)

func StartWriter() {
	// ---------- Ensure log directories exist ----------
	if err := os.MkdirAll(filepath.Dir(config.App.JSONLogFile), 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Dir(config.App.CSVLogFile), 0755); err != nil {
		panic(err)
	}

	// ---------- JSON ----------
	jsonFile, err := os.OpenFile(
		config.App.JSONLogFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	jsonWriter := bufio.NewWriterSize(jsonFile, 32*1024)

	// ---------- CSV ----------
	csvFile, err := os.OpenFile(
		config.App.CSVLogFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	csvBuf := bufio.NewWriterSize(csvFile, 32*1024)
	csvWriter := csv.NewWriter(csvBuf)

	// ---------- CSV header (once) ----------
	if stat, _ := csvFile.Stat(); stat.Size() == 0 {
		_ = csvWriter.Write([]string{
			"timestamp",
			"host",
			"battery_percent",
			"battery_state",
			"cpu_usage_percent",
			"memory_usage_percent",
			"disk_usage_percent",
			"net_down_kbps",
			"net_up_kbps",
			"net_total_kbps",
		})
		csvWriter.Flush()
		csvBuf.Flush()
	}

	go func() {
		defer func() {
			jsonWriter.Flush()
			csvWriter.Flush()
			csvBuf.Flush()
			jsonFile.Sync()
			csvFile.Sync()
			jsonFile.Close()
			csvFile.Close()
		}()

		for {
			select {
			case m := <-logCh:
				// JSON
				if b, err := json.Marshal(m); err == nil {
					jsonWriter.Write(b)
					jsonWriter.WriteByte('\n')
				}

				// CSV
				_ = csvWriter.Write([]string{
					m.Timestamp,
					m.Host,
					strconv.Itoa(m.Battery.Percent),
					m.Battery.State,
					formatFloat(m.CPU.Usage),
					formatFloat(m.Memory.Usage),
					formatFloat(m.Disk.Usage),
					strconv.FormatUint(m.Network.DownKB, 10),
					strconv.FormatUint(m.Network.UpKB, 10),
					strconv.FormatUint(m.Network.TotalKB, 10),
				})

				jsonWriter.Flush()
				csvWriter.Flush()
				csvBuf.Flush()

			case <-stopCh:
				return
			}
		}
	}()
}
