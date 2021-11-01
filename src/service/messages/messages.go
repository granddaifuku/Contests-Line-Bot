package messages

import (
	"context"

	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/granddaifuku/contest_line_bot/src/domain/messages"
	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
)

type MessageService interface {
	BuildMessages(
		ctx context.Context,
		atc []domain.AtcoderInfo,
		cdf []domain.CodeforcesInfo,
		ykc []domain.YukicoderInfo,
	) ([]*linebot.FlexMessage, error)
}

type messageService struct{}

func NewMessageService() MessageService {
	return &messageService{}
}

func (ms *messageService) BuildMessages(
	ctx context.Context,
	atc []domain.AtcoderInfo,
	cdf []domain.CodeforcesInfo,
	ykc []domain.YukicoderInfo,
) ([]*linebot.FlexMessage, error) {
	atcMsgs := make([]*linebot.BoxComponent, len(atc))
	cdfMsgs := make([]*linebot.BoxComponent, len(cdf))
	ykcMsgs := make([]*linebot.BoxComponent, len(ykc))

	for i, info := range atc {
		mes, err := messages.NewAtcoderMessage(info)
		if err != nil {
			return nil, xerrors.Errorf("Error when Creating AtCoder Message")
		}
		atcMsgs[i] = mes
	}

	for i, info := range cdf {
		mes, err := messages.NewCodeforcesMessage(info)
		if err != nil {
			return nil, xerrors.Errorf("Error when Creating Codeforces Message")
		}
		cdfMsgs[i] = mes
	}

	for i, info := range ykc {
		mes, err := messages.NewYukicoderMessage(info)
		if err != nil {
			return nil, xerrors.Errorf("Error when Creating Yukicoder Message")
		}
		ykcMsgs[i] = mes
	}
	msgs := make([]*linebot.FlexMessage, consts.NumContests)
	msgs[0] = messages.NewMessage(atcMsgs)
	msgs[1] = messages.NewMessage(cdfMsgs)
	msgs[2] = messages.NewMessage(ykcMsgs)

	return msgs, nil
}