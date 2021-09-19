package domain

import "time"

type AtcoderInfo struct {
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	RatedRange string
}

func NewAtCoderInfo(
	name string,
	startTime time.Time,
	endTime time.Time,
	ratedRange string,
) AtcoderInfo {
	return AtcoderInfo{Name: name, StartTime: startTime, EndTime: endTime, RatedRange: ratedRange}
}
