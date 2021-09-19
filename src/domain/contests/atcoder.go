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
	ratedRangeBegin int,
	ratedRangeEnd int,
) AtcoderInfo {
	if ratedRangeBegin == -1 {

	}
}
