package infrastructure

import (
	"context"
	"net/http"

	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/granddaifuku/contest_line_bot/src/internal/envs"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"golang.org/x/xerrors"
)

type notificatorPersistence struct {
	client *linebot.Client
}

func NewNotificatorPersistence(
	client *http.Client,
) (repository.NotificatorRepository, error) {
	env, _ := envs.LoadEnv()
	opts := []linebot.ClientOption{}
	if client != nil {
		opts = append(opts, linebot.WithHTTPClient(client))
	}
	bot, err := linebot.New(env.ChannelSecret, env.ChannelToken, opts...)
	if err != nil {
		return nil, xerrors.Errorf("Error when Creating Client: %w", err)
	}

	return &notificatorPersistence{client: bot}, nil
}

func (np *notificatorPersistence) Broadcast(
	ctx context.Context,
	msgs []*linebot.FlexMessage,
) error {
	for _, msg := range msgs {
		_, err := np.client.BroadcastMessage(msg).WithContext(ctx).Do()
		if err != nil {
			return xerrors.Errorf("Error when Broadcasting Messages: %w", err)
		}
	}

	return nil
}

func (np *notificatorPersistence) Reply(
	ctx context.Context,
	replyToken string,
	msgs []*linebot.FlexMessage,
) error {
	for _, msg := range msgs {
		_, err := np.client.ReplyMessage(replyToken, msg).WithContext(ctx).Do()
		if err != nil {
			return xerrors.Errorf("Error when Replying Messages: %w", err)
		}
	}

	return nil
}
