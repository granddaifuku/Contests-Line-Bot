package infrastructure

import (
	"context"
	"net/http"
	"strings"

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
) repository.NotificatorRepository {
	env, _ := envs.LoadEnv()
	opts := []linebot.ClientOption{}
	if client != nil {
		opts = append(opts, linebot.WithHTTPClient(client))
	}
	bot, err := linebot.New(env.ChannelSecret, env.ChannelToken, opts...)
	if err != nil {
		panic(err)
	}

	return &notificatorPersistence{client: bot}
}

func (np *notificatorPersistence) ExtractTokens(
	ctx context.Context,
	req *http.Request,
) ([]string, error) {
	events, err := np.client.ParseRequest(req)
	tokens := make([]string, 0)

	for _, event := range events {
		switch msg := event.Message.(type) {
		case *linebot.TextMessage:
			// Extract iff the message text contains the specified word
			if strings.Contains(msg.Text, "コンテスト") {
				tokens = append(tokens, event.Source.UserID)
			}
		}
	}
	if err != nil {
		return nil, xerrors.Errorf("Error when Parsing Request: %w", err)
	}

	return tokens, nil
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
	to string,
	msgs []*linebot.FlexMessage,
) error {
	for _, msg := range msgs {
		_, err := np.client.PushMessage(to, msg).WithContext(ctx).Do()
		if err != nil {
			return xerrors.Errorf("Error when Replying Messages: %w", err)
		}
	}

	return nil
}
