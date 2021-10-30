package repository

import (
	"context"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessageRepository interface {
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
