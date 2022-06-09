package messages

import (
	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// Return the flex message consists of three box components
// 1. Contest Name
// 2. Contest Time
// 3. Contest Rated Range
func NewAtcoderMessage(info domain.AtcoderInfo) (*linebot.BoxComponent, error) {
	nameInfo := newMessageContestName(info.Name)
	timeInfo := newMessageContestTime(info.StartTime, info.EndTime)
	rangeInfo := newMessageContestRange(info.RatedRange)
	nameBox, err := newHorizontalBoxComponent(nameInfo)
	if err != nil {
		return nil, err
	}
	timeBox, err := newHorizontalBoxComponent(timeInfo)
	if err != nil {
		return nil, err
	}
	rangeBox, err := newHorizontalBoxComponent(rangeInfo)
	if err != nil {
		return nil, err
	}
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			nameBox,
			timeBox,
			rangeBox,
		},
	}, nil
}
