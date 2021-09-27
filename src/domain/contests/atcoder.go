package domain

import (
	"strconv"
	"strings"
	"time"

	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"golang.org/x/xerrors"
)

var jst = time.FixedZone("Azia/Tokyo", 9*60*60)

type AtcoderInfo struct {
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	RatedRange string
}

func NewAtCoderInfo(
	name string,
	start string,
	duration string,
	ratedRange string,
) (AtcoderInfo, error) {
	info := &AtcoderInfo{}
	tz := "+0900"
	// Delete the timezone suffix
	if !strings.HasSuffix(start, tz) {
		return *info, xerrors.New("Error duration has no timezone suffix")
	}
	start = strings.TrimSuffix(start, tz)
	dur := strings.Split(duration, ":") // Separate duratino to hours and minutes
	startTime, err := time.ParseInLocation(consts.TimeLayout, start, jst)
	if err != nil {
		return *info, xerrors.Errorf("Error when Parsing Start Time: %w", err)
	}
	hours, err := strconv.Atoi(dur[0])
	if err != nil {
		return *info, xerrors.Errorf("Error when Converting String Hours to Int: %w", err)
	}
	minutes, err := strconv.Atoi(dur[1])
	if err != nil {
		return *info, xerrors.Errorf("Error when Converting String Minutes to Int: %w", err)
	}
	endTime := startTime.Add(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)

	return AtcoderInfo{Name: name, StartTime: startTime, EndTime: endTime, RatedRange: ratedRange}, nil
}
