package messages

import (
	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewCodeforcesMessage(info domain.CodeforcesInfo) (*linebot.BoxComponent, error) {
	nameInfo := newMessageContestName(info.Name)
	timeInfo := newMessageContestTime(info.StartTime, info.EndTime)
	nameBox, err := newHorizontalBoxComponent(nameInfo)
	if err != nil {
		return nil, err
	}
	timeBox, err := newHorizontalBoxComponent(timeInfo)
	if err != nil {
		return nil, err
	}
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			nameBox,
			timeBox,
		},
	}, nil
}
