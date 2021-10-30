package messages

import (
	"time"

	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
)

func NewMessage(contents []*linebot.BoxComponent) *linebot.FlexMessage {
	return linebot.NewFlexMessage(
		"Hello!",
		&linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Contents: newFlexComponent(contents),
			},
		})
}

func newMessageSeparator() *linebot.SeparatorComponent {
	return &linebot.SeparatorComponent{
		Type: linebot.FlexComponentTypeSeparator,
	}
}

func newMessageContestName(name string) []linebot.TextComponent {
	texts := make([]linebot.TextComponent, 2)
	texts[0] = newTextComponent("Name", consts.TitleColor)
	texts[1] = newTextComponent(name, consts.ContentColor)

	return texts
}

func newMessageContestTime(startTime, endTime time.Time) []linebot.TextComponent {
	texts := make([]linebot.TextComponent, 2)
	contestTime := startTime.Format(consts.TimeFormat) + " - " + endTime.Format(consts.TimeFormat)
	texts[0] = newTextComponent("Time", consts.TitleColor)
	texts[1] = newTextComponent(contestTime, consts.ContentColor)

	return texts
}

func newMessageContestRange(ratedRange string) []linebot.TextComponent {
	texts := make([]linebot.TextComponent, 2)
	texts[0] = newTextComponent("Range", consts.TitleColor)
	texts[1] = newTextComponent(ratedRange, consts.ContentColor)

	return texts
}

func newFlexComponent(boxes []*linebot.BoxComponent) []linebot.FlexComponent {
	cpnts := make([]linebot.FlexComponent, 0)
	for _, box := range boxes {
		cpnts = append(cpnts, box)
		cpnts = append(cpnts, newMessageSeparator())
	}

	return cpnts
}

func newTextComponent(text string, color string) linebot.TextComponent {
	return linebot.TextComponent{
		Type:       linebot.FlexComponentTypeText,
		Text:       text,
		AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
		Size:       linebot.FlexTextSizeTypeSm,
		Color:      color,
	}
}

func newHorizontalBoxComponent(texts []linebot.TextComponent) (*linebot.BoxComponent, error) {
	if len(texts) != 2 {
		return nil, xerrors.New("The length of texts should be 2")
	}
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&texts[0],
			&texts[1],
		},
	}, nil
}
