package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/granddaifuku/contest_line_bot/src/domain/service"
	"github.com/granddaifuku/contest_line_bot/src/infrastructure"
	"github.com/granddaifuku/contest_line_bot/src/interfaces"
	"github.com/granddaifuku/contest_line_bot/src/usecase"
)

var (
	crawl = flag.Bool("c", false, "Run crawler")
)

func main() {
	flag.Parse()
	// Crawling Option
	if *crawl {
		// Persistence
		dp := infrastructure.NewDatabasePersistence(nil)
		rp := infrastructure.NewRequestPersistence(nil)

		// Service
		cs := service.NewCrawlerService(rp)

		// Usecase
		cu := usecase.NewCrawlerUsecase(cs, dp)

		// Interfaces
		ch := interfaces.NewCrawlerHandler(cu)

		ch.Crawl()

		return
	}

	// Persistence
	dp := infrastructure.NewDatabasePersistence(nil)
	np := infrastructure.NewNotificatorPersistence(nil)

	// Service
	ns := service.NewNotificatorService()

	// Usecase
	nu := usecase.NewNotificatorUsecase(ns, np, dp)

	// Interfaces
	nh := interfaces.NewNotificatorHandler(nu)

	// Define routes
	http.HandleFunc("/callback", nh.HandleReplyMessage)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
