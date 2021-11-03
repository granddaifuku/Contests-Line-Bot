package messages

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	type args struct {
		boxes    []*linebot.BoxComponent
		platform string
	}
	tests := []struct {
		name string
		args args
		want *linebot.FlexMessage
	}{

		{
			name: "Some args",
			args: args{
				boxes: []*linebot.BoxComponent{
					{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.BoxComponent{
								Type:   linebot.FlexComponentTypeBox,
								Layout: linebot.FlexBoxLayoutTypeBaseline,
								Contents: []linebot.FlexComponent{
									&linebot.TextComponent{
										Type:       linebot.FlexComponentTypeText,
										Text:       "Title",
										AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
										Size:       linebot.FlexTextSizeTypeSm,
										Color:      "#aaaaaa",
									},
									&linebot.TextComponent{
										Type:       linebot.FlexComponentTypeText,
										Text:       "Content",
										AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
										Size:       linebot.FlexTextSizeTypeSm,
										Color:      "#666666",
									},
								},
							},
						},
					},
				},
				platform: "platform",
			},
			want: &linebot.FlexMessage{
				AltText: "Hello!",
				Contents: &linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   "platform",
								Size:   linebot.FlexTextSizeTypeXl,
								Weight: linebot.FlexTextWeightTypeBold,
							},
							&linebot.BoxComponent{
								Type:   linebot.FlexComponentTypeBox,
								Layout: linebot.FlexBoxLayoutTypeVertical,
								Contents: []linebot.FlexComponent{
									&linebot.BoxComponent{
										Type:   linebot.FlexComponentTypeBox,
										Layout: linebot.FlexBoxLayoutTypeBaseline,
										Contents: []linebot.FlexComponent{
											&linebot.TextComponent{
												Type:       linebot.FlexComponentTypeText,
												Text:       "Title",
												AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
												Size:       linebot.FlexTextSizeTypeSm,
												Color:      "#aaaaaa",
											},
											&linebot.TextComponent{
												Type:       linebot.FlexComponentTypeText,
												Text:       "Content",
												AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
												Size:       linebot.FlexTextSizeTypeSm,
												Color:      "#666666",
											},
										},
									},
								},
							},
							&linebot.SeparatorComponent{
								Type: linebot.FlexComponentTypeSeparator,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMessage(tt.args.boxes, tt.args.platform)
			opts := cmpopts.IgnoreUnexported(linebot.FlexMessage{})
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Errorf("NewMessage() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}

func TestNewMessaegSeparator(t *testing.T) {
	want := &linebot.SeparatorComponent{
		Type: linebot.FlexComponentTypeSeparator,
	}
	got := newMessageSeparator()

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("NewMessageSeparator() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewMessageContestName(t *testing.T) {
	flex1 := 1
	flex5 := 5
	want := []linebot.TextComponent{
		{
			Type:  linebot.FlexComponentTypeText,
			Text:  "Name",
			Size:  linebot.FlexTextSizeTypeSm,
			Color: "#aaaaaa",
			Flex:  &flex1,
		},
		{
			Type:  linebot.FlexComponentTypeText,
			Text:  "Test",
			Size:  linebot.FlexTextSizeTypeSm,
			Color: "#666666",
			Flex:  &flex5,
			Wrap:  true,
		},
	}
	got := newMessageContestName("Test")

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newMessageContestName() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewMessageContestTime(t *testing.T) {
	flex1 := 1
	flex5 := 5
	jst := time.FixedZone("Azia/Tokyo", 9*60*60)
	want := []linebot.TextComponent{
		{
			Type:  linebot.FlexComponentTypeText,
			Text:  "Time",
			Size:  linebot.FlexTextSizeTypeSm,
			Color: "#aaaaaa",
			Flex:  &flex1,
		},
		{
			Type:  linebot.FlexComponentTypeText,
			Text:  "2021-09-30 21:00:00 - 2021-09-30 22:40:00",
			Size:  linebot.FlexTextSizeTypeSm,
			Color: "#666666",
			Flex:  &flex5,
			Wrap:  true,
		},
	}
	got := newMessageContestTime(time.Date(2021, 9, 30, 21, 0, 0, 0, jst), time.Date(2021, 9, 30, 22, 40, 0, 0, jst))

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newMessageContestTime() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewMessageContestRange(t *testing.T) {
	flex1 := 1
	flex5 := 5
	want := []linebot.TextComponent{
		{
			Type:  linebot.FlexComponentTypeText,
			Text:  "Range",
			Size:  linebot.FlexTextSizeTypeSm,
			Color: "#aaaaaa",
			Flex:  &flex1,
		},
		{
			Type:  linebot.FlexComponentTypeText,
			Text:  " ~ 1999",
			Size:  linebot.FlexTextSizeTypeSm,
			Color: "#666666",
			Flex:  &flex5,
			Wrap:  true,
		},
	}
	got := newMessageContestRange(" ~ 1999")

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newMessageContestRange() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewFlexComponent(t *testing.T) {
	args := []*linebot.BoxComponent{
		{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeHorizontal,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:       linebot.FlexComponentTypeText,
							Text:       "Title",
							AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
							Size:       linebot.FlexTextSizeTypeSm,
							Color:      "#aaaaaa",
						},
						&linebot.TextComponent{
							Type:       linebot.FlexComponentTypeText,
							Text:       "Content",
							AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
							Size:       linebot.FlexTextSizeTypeSm,
							Color:      "#666666",
						},
					},
				},
			},
		},
	}
	want := []linebot.FlexComponent{
		&linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeHorizontal,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:       linebot.FlexComponentTypeText,
							Text:       "Title",
							AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
							Size:       linebot.FlexTextSizeTypeSm,
							Color:      "#aaaaaa",
						},
						&linebot.TextComponent{
							Type:       linebot.FlexComponentTypeText,
							Text:       "Content",
							AdjustMode: linebot.FlexComponentAdjustModeTypeShrinkToFit,
							Size:       linebot.FlexTextSizeTypeSm,
							Color:      "#666666",
						},
					},
				},
			},
		},
		&linebot.SeparatorComponent{
			Type: linebot.FlexComponentTypeSeparator,
		},
	}
	got := newFlexComponent(args)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newFlexComponent() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewTextComponent(t *testing.T) {
	flex := 1
	want := linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  "Test",
		Size:  linebot.FlexTextSizeTypeSm,
		Color: "#666666",
		Flex:  &flex,
	}
	got := newTextComponent("Test", "#666666")

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("newTextComponent() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestNewTextPlatform(t *testing.T) {
	tests := []struct {
		name string
		args string
		want linebot.FlexComponent
	}{
		{
			name: "Success",
			args: "AtCoder",
			want: &linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   "AtCoder",
				Size:   linebot.FlexTextSizeTypeXl,
				Weight: linebot.FlexTextWeightTypeBold,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newTextPlatform(tt.args)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("newTextPlatform() returned invalid results (-got +want):\n%s", diff)
			}
		})
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
				Type:    linebot.FlexComponentTypeBox,
				Layout:  linebot.FlexBoxLayoutTypeBaseline,
				Margin:  linebot.FlexComponentMarginTypeLg,
				Spacing: linebot.FlexComponentSpacingTypeSm,
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
