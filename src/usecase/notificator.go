package usecase

import (
	"context"

	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/granddaifuku/contest_line_bot/src/domain/service"
	"golang.org/x/xerrors"
)

type NotificatorUsecase interface {
	Reply(
		ctx context.Context,
		replyToken string,
	) error

	Broadcast(
		ctx context.Context,
	) error
}

type notificatorUsecase struct {
	ns service.NotificatorService
	nr repository.NotificatorRepository
	dr repository.DatabaseRepository
}

func NewNotificatorUsecase(
	ns service.NotificatorService,
	nr repository.NotificatorRepository,
	dr repository.DatabaseRepository,
) NotificatorUsecase {
	return &notificatorUsecase{
		ns: ns,
		nr: nr,
		dr: dr,
	}
}

func (nu *notificatorUsecase) Reply(
	ctx context.Context,
	replyToken string,
) error {
	// Get information from the database
	atcInfo, err := nu.dr.BatchGet(ctx, "AtCoder")
	if err != nil {
		return xerrors.Errorf("Error when Selecting AtCoder Database: %w", err)
	}
	cdfInfo, err := nu.dr.BatchGet(ctx, "Codeforces")
	if err != nil {
		return xerrors.Errorf("Error when Selecting Codeforces Database: %w", err)
	}
	ykcInfo, err := nu.dr.BatchGet(ctx, "Yukicoder")
	if err != nil {
		return xerrors.Errorf("Error when Selecting Yukicoder Database: %w", err)
	}

	atc := make([]domain.AtcoderInfo, len(atcInfo))
	for i := range atcInfo {
		info, ok := atcInfo[i].(domain.AtcoderInfo)
		if !ok {
			return xerrors.New("Failed to Cast AtCoder Info")
		}
		atc[i] = info
	}
	cdf := make([]domain.CodeforcesInfo, len(cdfInfo))
	for i := range cdfInfo {
		info, ok := cdfInfo[i].(domain.CodeforcesInfo)
		if !ok {
			return xerrors.New("Failed to Cast Codeforces Info")
		}
		cdf[i] = info
	}
	ykc := make([]domain.YukicoderInfo, len(ykcInfo))
	for i := range ykcInfo {
		info, ok := atcInfo[i].(domain.YukicoderInfo)
		if !ok {
			return xerrors.New("Failed to Cast Yukicoder Info")
		}
		ykc[i] = info
	}

	// Convert information to messages
	msgs, err := nu.ns.BuildMessages(ctx, atc, cdf, ykc)
	if err != nil {
		return xerrors.Errorf("Error when Calling BuildMessages Function: %w", err)
	}

	// Reply
	err = nu.nr.Reply(ctx, replyToken, msgs)
	if err != nil {
		return xerrors.Errorf("Error when Calling Reply Function: %w", err)
	}

	return nil
}

func (nu *notificatorUsecase) Broadcast(
	ctx context.Context,
) error {
	// Get information from the database
	atcInfo, err := nu.dr.BatchGet(ctx, "AtCoder")
	if err != nil {
		return xerrors.Errorf("Error when Selecting AtCoder Database: %w", err)
	}
	cdfInfo, err := nu.dr.BatchGet(ctx, "Codeforces")
	if err != nil {
		return xerrors.Errorf("Error when Selecting Codeforces Database: %w", err)
	}
	ykcInfo, err := nu.dr.BatchGet(ctx, "Yukicoder")
	if err != nil {
		return xerrors.Errorf("Error when Selecting Yukicoder Database: %w", err)
	}

	atc := make([]domain.AtcoderInfo, len(atcInfo))
	for i := range atcInfo {
		info, ok := atcInfo[i].(domain.AtcoderInfo)
		if !ok {
			return xerrors.New("Failed to Cast AtCoder Info")
		}
		atc[i] = info
	}
	cdf := make([]domain.CodeforcesInfo, len(cdfInfo))
	for i := range cdfInfo {
		info, ok := cdfInfo[i].(domain.CodeforcesInfo)
		if !ok {
			return xerrors.New("Failed to Cast Codeforces Info")
		}
		cdf[i] = info
	}
	ykc := make([]domain.YukicoderInfo, len(ykcInfo))
	for i := range ykcInfo {
		info, ok := atcInfo[i].(domain.YukicoderInfo)
		if !ok {
			return xerrors.New("Failed to Cast Yukicoder Info")
		}
		ykc[i] = info
	}

	// Convert information to messages
	msgs, err := nu.ns.BuildMessages(ctx, atc, cdf, ykc)
	if err != nil {
		return xerrors.Errorf("Error when Calling BuildMessages Function: %w", err)
	}

	// Reply
	err = nu.nr.Broadcast(ctx, msgs)
	if err != nil {
		return xerrors.Errorf("Error when Calling Broadcast Function: %w", err)
	}

	return nil
}
