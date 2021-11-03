package interfaces

import (
	"context"
	"log"
	"net/http"

	"github.com/granddaifuku/contest_line_bot/src/usecase"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type NotificatorHandler interface {
	HandleReplyMessage(
		w http.ResponseWriter,
		req *http.Request,
	)
}

type notificatorHandler struct {
	nu usecase.NotificatorUsecase
}

func NewNotificatorHandler(nu usecase.NotificatorUsecase) NotificatorHandler {
	return &notificatorHandler{
		nu: nu,
	}
}

func (nh *notificatorHandler) HandleReplyMessage(w http.ResponseWriter, req *http.Request) {
	tokens, err := nh.nu.Extract(context.Background(), req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	err = nh.nu.Reply(context.Background(), tokens)
	if err != nil {
		log.Print(err)
	}
}
