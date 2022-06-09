package service

import (
	"context"

	contests "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	messages "github.com/granddaifuku/contest_line_bot/src/domain/model/messages"
	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type NotificatorService interface {
	BuildMessages(
		ctx context.Context,
		atc []contests.AtcoderInfo,
		cdf []contests.CodeforcesInfo,
		ykc []contests.YukicoderInfo,
	) ([]*linebot.FlexMessage, error)
}

type notificatorService struct{}

func NewNotificatorService() NotificatorService {
	return &notificatorService{}
}

func (ns *notificatorService) BuildMessages(
	ctx context.Context,
	atc []contests.AtcoderInfo,
	cdf []contests.CodeforcesInfo,
	ykc []contests.YukicoderInfo,
) ([]*linebot.FlexMessage, error) {
	atcMsgs := make([]*linebot.BoxComponent, len(atc))
	cdfMsgs := make([]*linebot.BoxComponent, len(cdf))
	ykcMsgs := make([]*linebot.BoxComponent, len(ykc))

	for i, info := range atc {
		mes, err := messages.NewAtcoderMessage(info)
		if err != nil {
			return nil, err
		}
		atcMsgs[i] = mes
	}

	for i, info := range cdf {
		mes, err := messages.NewCodeforcesMessage(info)
		if err != nil {
			return nil, err
		}
		cdfMsgs[i] = mes
	}

	for i, info := range ykc {
		mes, err := messages.NewYukicoderMessage(info)
		if err != nil {
			return nil, err
		}
		ykcMsgs[i] = mes
	}
	msgs := make([]*linebot.FlexMessage, consts.NumContests)
	msgs[0] = messages.NewMessage(atcMsgs, "AtCoder")
	msgs[1] = messages.NewMessage(cdfMsgs, "Codeforces")
	msgs[2] = messages.NewMessage(ykcMsgs, "Yukicoder")

	return msgs, nil
}
