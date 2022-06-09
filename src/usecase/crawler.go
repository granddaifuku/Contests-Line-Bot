package usecase

import (
	"context"

	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/granddaifuku/contest_line_bot/src/domain/service"
)

type CrawlerUsecase interface {
	Crawl(
		ctx context.Context,
	) error
}

type crawlerUsecase struct {
	cs service.CrawlerService
	dr repository.DatabaseRepository
}

func NewCrawlerUsecase(
	cs service.CrawlerService,
	dr repository.DatabaseRepository,
) CrawlerUsecase {
	return &crawlerUsecase{
		cs: cs,
		dr: dr,
	}
}

// Crawl the contest platforms and gather the future contests information
func (cu *crawlerUsecase) Crawl(
	ctx context.Context,
) error {
	// Fetch contests information
	atc, err := cu.cs.FetchAtcoderInfo(ctx)
	if err != nil {
		return err
	}
	cdf, err := cu.cs.FetchCodeforcesInfo(ctx)
	if err != nil {
		return err
	}
	ykc, err := cu.cs.FetchYukicoderInfo(ctx)
	if err != nil {
		return err
	}

	// Clear Table
	err = cu.dr.ClearTables(ctx)
	if err != nil {
		return err
	}

	// Insert
	for _, info := range atc {
		err = cu.dr.InsertAtcoder(ctx, info)
		if err != nil {
			return err
		}
	}
	for _, info := range cdf {
		err = cu.dr.InsertCodeforces(ctx, info)
		if err != nil {
			return err
		}
	}
	for _, info := range ykc {
		err = cu.dr.InsertYukicoder(ctx, info)
		if err != nil {
			return err
		}
	}

	return nil
}
