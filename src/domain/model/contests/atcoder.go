package domain

import (
	"strconv"
	"strings"
	"time"

	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"github.com/pkg/errors"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

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
		return *info, errors.New("Error Duration has No Timezone Suffix")
	}
	start = strings.TrimSuffix(start, tz)

	startTime, err := time.ParseInLocation(consts.TimeFormat, start, jst)
	if err != nil {
		return *info, errors.WithStack(err)
	}

	dur := strings.Split(duration, ":") // Separate duration to hours and minutes
	hours, err := strconv.Atoi(dur[0])
	if err != nil {
		return *info, errors.WithStack(err)
	}
	minutes, err := strconv.Atoi(dur[1])
	if err != nil {
		return *info, errors.WithStack(err)
	}
	endTime := startTime.Add(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)

	return AtcoderInfo{Name: name, StartTime: startTime, EndTime: endTime, RatedRange: ratedRange}, nil
}
