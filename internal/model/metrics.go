package model

type Metrics struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	App       string `json:"app"`
	Host      string `json:"host"`

	Battery struct {
		Percent int    `json:"percent"`
		State   string `json:"state"`
	} `json:"battery"`

	CPU struct {
		Usage float64 `json:"usage_percent"`
	} `json:"cpu"`

	Memory struct {
		Usage float64 `json:"usage_percent"`
	} `json:"memory"`

	Disk struct {
		Usage float64 `json:"usage_percent"`
	} `json:"disk"`

	Network struct {
		DownKB  uint64 `json:"down_kbps"`
		UpKB    uint64 `json:"up_kbps"`
		TotalKB uint64 `json:"total_kbps"`
	} `json:"network"`
}
