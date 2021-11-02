package usecase

import (
	"context"

	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	service "github.com/granddaifuku/contest_line_bot/src/service/contests"
	"golang.org/x/xerrors"
)

type ContestUsecase interface {
	Crawl(
		ctx context.Context,
	) error
}

type contestUsecase struct {
	cs service.ContestService
	rr repository.RequestRepository
	dr repository.DatabaseRepository
}

func NewContestUsecase(
	cs service.ContestService,
	dr repository.DatabaseRepository,
) ContestUsecase {
	return &contestUsecase{
		cs: cs,
		dr: dr,
	}
}

func (cu *contestUsecase) Crawl(
	ctx context.Context,
) error {
	// Fetch contests information
	atc, err := cu.cs.FetchAtcoderInfo(ctx)
	if err != nil {
		return xerrors.Errorf("Error when Fetching AtCoder Information: %w", err)
	}
	cdf, err := cu.cs.FetchCodeforcesInfo(ctx)
	if err != nil {
		return xerrors.Errorf("Error when Fetching Codeforces Information: %w", err)
	}
	ykc, err := cu.cs.FetchYukicoderInfo(ctx)
	if err != nil {
		return xerrors.Errorf("Error when Fetching Yukicoder Information: %w", err)
	}

	// Clear Table
	cu.dr.ClearTables(ctx)

	// Insert
	for _, info := range atc {
		err = cu.dr.InsertAtcoder(ctx, info)
		if err != nil {
			return xerrors.Errorf("Error when Inserting AtCoder Table: %w", err)
		}
	}
	for _, info := range cdf {
		err = cu.dr.InsertCodeforces(ctx, info)
		if err != nil {
			return xerrors.Errorf("Error when Inserting Codeforces Table: %w", err)
		}
	}
	for _, info := range ykc {
		err = cu.dr.InsertYukicoder(ctx, info)
		if err != nil {
			return xerrors.Errorf("Error when Inserting Yukicoder Table: %w", err)
		}
	}

	return nil
}
