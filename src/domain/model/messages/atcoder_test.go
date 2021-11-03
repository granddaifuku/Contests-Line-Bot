package messages

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
)

func TestNewAtcoderMessage(t *testing.T) {
	flex1 := 1
	flex5 := 5
	jst := time.FixedZone("Azia/Tokyo", 9*60*60)
	arg := domain.AtcoderInfo{
		Name:       "AtCoder Beginner Contest 999",
		StartTime:  time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
		EndTime:    time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
		RatedRange: " ~ 1999",
	}
	want := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:    linebot.FlexComponentTypeBox,
				Layout:  linebot.FlexBoxLayoutTypeBaseline,
				Spacing: linebot.FlexComponentSpacingType(linebot.FlexSpacerSizeTypeSm),
				Margin:  linebot.FlexComponentMarginTypeLg,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  "Name",
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#aaaaaa",
						Flex:  &flex1,
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  "AtCoder Beginner Contest 999",
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#666666",
						Flex:  &flex5,
						Wrap:  true,
					},
				},
			},

			&linebot.BoxComponent{
				Type:    linebot.FlexComponentTypeBox,
				Layout:  linebot.FlexBoxLayoutTypeBaseline,
				Spacing: linebot.FlexComponentSpacingType(linebot.FlexSpacerSizeTypeSm),
				Margin:  linebot.FlexComponentMarginTypeLg,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  "Time",
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#aaaaaa",
						Flex:  &flex1,
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  "2021-12-30 21:00:00 - 2021-12-30 22:40:00",
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#666666",
						Flex:  &flex5,
						Wrap:  true,
					},
				},
			},
			&linebot.BoxComponent{
				Type:    linebot.FlexComponentTypeBox,
				Layout:  linebot.FlexBoxLayoutTypeBaseline,
				Spacing: linebot.FlexComponentSpacingType(linebot.FlexSpacerSizeTypeSm),
				Margin:  linebot.FlexComponentMarginTypeLg,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  "Range",
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#aaaaaa",
						Flex:  &flex1,
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  " ~ 1999",
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#666666",
						Flex:  &flex5,
						Wrap:  true,
					},
				},
			},
		},
	}

	got, err := NewAtcoderMessage(arg)
	assert.Nil(t, err)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("NewAtcoderMessage() returned invalid results (-got +want):\n %s", diff)
	}
}
