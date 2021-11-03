package interfaces

import (
	"context"
	"log"

	"github.com/granddaifuku/contest_line_bot/src/usecase"
)

type CrawlerHandler interface {
	Crawl()
}

type crawlerHandler struct {
	cu usecase.CrawlerUsecase
}

func NewCrawlerHandler(cu usecase.CrawlerUsecase) CrawlerHandler {
	return &crawlerHandler{
		cu: cu,
	}
}

func (ch *crawlerHandler) Crawl() {
	err := ch.cu.Crawl(context.Background())

	if err != nil {
		log.Print(err)
	}
}
