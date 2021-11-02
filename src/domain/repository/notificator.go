package repository

import (
	"context"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type NotificatorRepository interface {
	Broadcast(
		ctx context.Context,
		msgs []*linebot.FlexMessage,
	) error

	Reply(
		ctx context.Context,
		replyToken string,
		msgs []*linebot.FlexMessage,
	) error
}
