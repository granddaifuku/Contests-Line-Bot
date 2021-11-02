package service

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
)

func TestBuildMessages(t *testing.T) {
	type args struct {
		atc []domain.AtcoderInfo
		cdf []domain.CodeforcesInfo
		ykc []domain.YukicoderInfo
	}
	tests := []struct {
		name string
		args args
		want []*linebot.FlexMessage
	}{
		{
			name: "Success",
			args: args{
				atc: []domain.AtcoderInfo{},
				cdf: []domain.CodeforcesInfo{},
				ykc: []domain.YukicoderInfo{},
			},
			want: []*linebot.FlexMessage{
				{
					AltText: "Hello!",
					Contents: &linebot.BubbleContainer{
						Type: "bubble",
						Body: &linebot.BoxComponent{
							Type:     linebot.FlexComponentTypeBox,
							Layout:   linebot.FlexBoxLayoutTypeVertical,
							Contents: []linebot.FlexComponent{},
						},
					},
				},
				{
					AltText: "Hello!",
					Contents: &linebot.BubbleContainer{
						Type: "bubble",
						Body: &linebot.BoxComponent{
							Type:     linebot.FlexComponentTypeBox,
							Layout:   linebot.FlexBoxLayoutTypeVertical,
							Contents: []linebot.FlexComponent{},
						},
					},
				},
				{
					AltText: "Hello!",
					Contents: &linebot.BubbleContainer{
						Type: "bubble",
						Body: &linebot.BoxComponent{
							Type:     linebot.FlexComponentTypeBox,
							Layout:   linebot.FlexBoxLayoutTypeVertical,
							Contents: []linebot.FlexComponent{},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &notificatorService{}
			got, err := ns.BuildMessages(context.Background(), tt.args.atc, tt.args.cdf, tt.args.ykc)
			assert.Nil(t, err)

			opt := cmpopts.IgnoreUnexported(linebot.FlexMessage{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("notificatorService.BuildMessage() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}
