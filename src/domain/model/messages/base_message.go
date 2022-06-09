package messages

import (
	"time"

	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/pkg/errors"
)

func NewMessage(contents []*linebot.BoxComponent, platform string) *linebot.FlexMessage {
	body := make([]linebot.FlexComponent, 0)
	body = append(body, newTextPlatform(platform))
	body = append(body, newFlexComponent(contents)...)

	return linebot.NewFlexMessage(
		"Hello!",
		&linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Contents: body,
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
	texts[1] = newTextContentComponent(name, consts.ContentColor)

	return texts
}

func newMessageContestTime(startTime, endTime time.Time) []linebot.TextComponent {
	texts := make([]linebot.TextComponent, 2)
	contestTime := startTime.Format(consts.TimeFormat) + " - " + endTime.Format(consts.TimeFormat)
	texts[0] = newTextComponent("Time", consts.TitleColor)
	texts[1] = newTextContentComponent(contestTime, consts.ContentColor)

	return texts
}

func newMessageContestRange(ratedRange string) []linebot.TextComponent {
	texts := make([]linebot.TextComponent, 2)
	texts[0] = newTextComponent("Range", consts.TitleColor)
	texts[1] = newTextContentComponent(ratedRange, consts.ContentColor)

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
	flex := 1
	return linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  text,
		Size:  linebot.FlexTextSizeTypeSm,
		Color: color,
		Flex:  &flex,
	}
}

func newTextContentComponent(text string, color string) linebot.TextComponent {
	flex := 5
	return linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  text,
		Size:  linebot.FlexTextSizeTypeSm,
		Color: color,
		Flex:  &flex,
		Wrap:  true,
	}
}

func newTextPlatform(text string) linebot.FlexComponent {
	return &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   text,
		Size:   linebot.FlexTextSizeTypeXl,
		Weight: linebot.FlexTextWeightTypeBold,
	}
}

func newHorizontalBoxComponent(texts []linebot.TextComponent) (*linebot.BoxComponent, error) {
	if len(texts) != 2 {
		return nil, errors.New("The length of texts should be 2")
	}
	return &linebot.BoxComponent{
		Type:    linebot.FlexComponentTypeBox,
		Layout:  linebot.FlexBoxLayoutTypeBaseline,
		Margin:  linebot.FlexComponentMarginTypeLg,
		Spacing: linebot.FlexComponentSpacingTypeSm,
		Contents: []linebot.FlexComponent{
			&texts[0],
			&texts[1],
		},
	}, nil
}
