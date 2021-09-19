package domain

import "time"

type YukicoderInfo struct {
	Name      string    `json:"Name"`
	StartTime time.Time `json:"Date"`
	EndTime   time.Time `json:"EndDate"`
}
