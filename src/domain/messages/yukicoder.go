package messages

import (
	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
)

func NewYukicoderMessage(info domain.YukicoderInfo) (*linebot.BoxComponent, error) {
	nameInfo := newMessageContestName(info.Name)
	timeInfo := newMessageContestTime(info.StartTime, info.EndTime)
	nameBox, err := newHorizontalBoxComponent(nameInfo)
	if err != nil {
		return nil, xerrors.Errorf("Error when Building Name Box Component: %w", err)
	}
	timeBox, err := newHorizontalBoxComponent(timeInfo)
	if err != nil {
		return nil, xerrors.Errorf("Error when Building Time Box Component: %w", err)
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
