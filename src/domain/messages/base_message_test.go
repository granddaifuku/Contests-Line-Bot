package messages

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	tests := []struct {
		name string
		args []linebot.FlexComponent
		want *linebot.BubbleContainer
	}{
		{
			name: "No args",
			args: nil,
			want: &linebot.BubbleContainer{
				Type: linebot.FlexContainerTypeBubble,
				Body: &linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: nil,
				},
			},
		},
		{
			name: "Some args",
			args: []linebot.FlexComponent{
				&linebot.ImageComponent{
					Type:     linebot.FlexComponentTypeImage,
					URL:      "https://example.com/flex/images/image.jpg",
					Animated: true,
				},
				&linebot.SeparatorComponent{
					Type: linebot.FlexComponentTypeSeparator,
				},
				&linebot.TextComponent{
					Type:       linebot.FlexComponentTypeText,
					Text:       "Text in the box",
					AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
				},
				&linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: []linebot.FlexComponent{},
					Width:    "30px",
					Height:   "30px",
					Background: &linebot.BoxBackground{
						Type:           linebot.FlexBoxBackgroundTypeLinearGradient,
						Angle:          "0deg",
						StartColor:     "#ff0000",
						EndColor:       "#00ff00",
						CenterColor:    "#0000ff",
						CenterPosition: "10%",
					},
				},
			},
			want: &linebot.BubbleContainer{
				Type: linebot.FlexContainerTypeBubble,
				Body: &linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeVertical,
					Contents: []linebot.FlexComponent{
						&linebot.ImageComponent{
							Type:     linebot.FlexComponentTypeImage,
							URL:      "https://example.com/flex/images/image.jpg",
							Animated: true,
						},
						&linebot.SeparatorComponent{
							Type: linebot.FlexComponentTypeSeparator,
						},
						&linebot.TextComponent{
							Type:       linebot.FlexComponentTypeText,
							Text:       "Text in the box",
							AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
						},
						&linebot.BoxComponent{
							Type:     linebot.FlexComponentTypeBox,
							Layout:   linebot.FlexBoxLayoutTypeVertical,
							Contents: []linebot.FlexComponent{},
							Width:    "30px",
							Height:   "30px",
							Background: &linebot.BoxBackground{
								Type:           linebot.FlexBoxBackgroundTypeLinearGradient,
								Angle:          "0deg",
								StartColor:     "#ff0000",
								EndColor:       "#00ff00",
								CenterColor:    "#0000ff",
								CenterPosition: "10%",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMessage(tt.args)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMessage() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}

func TestNewMessaegSeparator(t *testing.T) {
	want := &linebot.SeparatorComponent{
		Type: linebot.FlexComponentTypeSeparator,
	}
	got := NewMessageSeparator()

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("NewMessageSeparator() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewMessageContestName(t *testing.T) {
	want := []linebot.TextComponent{
		{
			Type:       linebot.FlexComponentTypeText,
			Text:       "Name",
			AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
			Size:       linebot.FlexTextSizeTypeSm,
			Color:      "#aaaaaa",
		},
		{
			Type:       linebot.FlexComponentTypeText,
			Text:       "Test",
			AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
			Size:       linebot.FlexTextSizeTypeSm,
			Color:      "#666666",
		},
	}
	got := newMessageContestName("Test")

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newMessageContestName() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewMessageContestTime(t *testing.T) {
	jst := time.FixedZone("Azia/Tokyo", 9*60*60)
	want := []linebot.TextComponent{
		{
			Type:       linebot.FlexComponentTypeText,
			Text:       "Time",
			AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
			Size:       linebot.FlexTextSizeTypeSm,
			Color:      "#aaaaaa",
		},
		{
			Type:       linebot.FlexComponentTypeText,
			Text:       "2021-09-30 21:00:00 - 2021-09-30 22:40:00",
			AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
			Size:       linebot.FlexTextSizeTypeSm,
			Color:      "#666666",
		},
	}
	got := newMessageContestTime(time.Date(2021, 9, 30, 21, 0, 0, 0, jst), time.Date(2021, 9, 30, 22, 40, 0, 0, jst))

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newMessageContestTime() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewMessageContestRange(t *testing.T) {
	want := []linebot.TextComponent{
		{
			Type:       linebot.FlexComponentTypeText,
			Text:       "Range",
			AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
			Size:       linebot.FlexTextSizeTypeSm,
			Color:      "#aaaaaa",
		},
		{
			Type:       linebot.FlexComponentTypeText,
			Text:       " ~ 1999",
			AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
			Size:       linebot.FlexTextSizeTypeSm,
			Color:      "#666666",
		},
	}
	got := newMessageContestRange(" ~ 1999")

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newMessageContestRange() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewTextComponent(t *testing.T) {
	want := linebot.TextComponent{
		Type:       linebot.FlexComponentTypeText,
		Text:       "Test",
		AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
		Size:       linebot.FlexTextSizeTypeSm,
		Color:      "#666666",
	}
	got := newTextComponent("Test", "#666666")

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newTextComponent() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewHorizontalBoxComponent(t *testing.T) {
	tests := []struct {
		name    string
		args    []linebot.TextComponent
		want    *linebot.BoxComponent
		wantErr bool
	}{
		{
			name: "Success",
			args: []linebot.TextComponent{
				{
					Text: "Test Title",
				},
				{
					Text: "Test Content",
				},
			},
			want: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text: "Test Title",
					},
					&linebot.TextComponent{
						Text: "Test Content",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Invalid length of the arg",
			args:    []linebot.TextComponent{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newHorizontalBoxComponent(tt.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("newHorizontalBoxComponent() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}
