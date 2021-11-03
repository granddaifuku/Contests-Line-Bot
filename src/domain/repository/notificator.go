package repository

import (
	"context"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type NotificatorRepository interface {
	// Parse the requests and extract the reply tokens if event type is text
	ExtractTokens(
		ctx context.Context,
		req *http.Request,
	) ([]string, error)

	Broadcast(
		ctx context.Context,
		msgs []*linebot.FlexMessage,
	) error

	Reply(
		ctx context.Context,
		to string,
		msgs []*linebot.FlexMessage,
	) error
}
