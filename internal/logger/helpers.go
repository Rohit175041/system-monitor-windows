package logger

import "strconv"

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}
