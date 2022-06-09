package service

import (
	"bytes"
	"context"
	"sort"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"golang.org/x/xerrors"
)

type CrawlerService interface {
	FetchAtcoderInfo(
		ctx context.Context,
	) ([]domain.AtcoderInfo, error)
	FetchCodeforcesInfo(
		ctx context.Context,
	) ([]domain.CodeforcesInfo, error)
	FetchYukicoderInfo(
		ctx context.Context,
	) ([]domain.YukicoderInfo, error)
}

type crawlerService struct {
	rr repository.RequestRepository
}

func NewCrawlerService(rr repository.RequestRepository) CrawlerService {
	return &crawlerService{
		rr: rr,
	}
}

func (cs *crawlerService) FetchAtcoderInfo(ctx context.Context) ([]domain.AtcoderInfo, error) {
	info := make([]domain.AtcoderInfo, 0)
	body, err := cs.rr.Get(ctx, consts.AtcoderURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}

	// Scrape AtCoder's contests information
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, xerrors.Errorf("Error when Reading Document: %w", err)
	}

	// Dive into the upcoming contests inforamtion
	scraped := doc.Find("div#contest-table-upcoming > div.panel > div.table-responsive > table.table > tbody > tr").Text()

	splited := cs.arrangeAtcoderInfo(scraped)
	// Devide the slice every 5 elements.
	for i := 0; i < len(splited); i += 5 {
		startTime := splited[i]
		name := splited[i+1] + " " + splited[i+2] // category + name
		duration := splited[i+3]
		ratedRange := splited[i+4]
		at, err := domain.NewAtCoderInfo(name, startTime, duration, ratedRange)
		if err != nil {
			return nil, xerrors.Errorf("Error when Building AtCoder Info: %w", err)
		}
		info = append(info, at)
	}

	return info, nil
}

func (cs *crawlerService) FetchCodeforcesInfo(ctx context.Context) ([]domain.CodeforcesInfo, error) {
	api := domain.CodeforcesApi{}
	info := make([]domain.CodeforcesInfo, 0)
	// Call Codeforces' contests information api
	body, err := cs.rr.Get(ctx, consts.CodeforcesURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}
	err = cs.rr.DecodeJson(body, &api)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Decoding Function: %w", err)
	}
	for _, res := range api.Result {
		// Only want future contests
		if res.Phase != "BEFORE" {
			continue
		}
		cf := domain.NewCodeforcesInfo(&res)
		info = append(info, cf)
	}
	// Make Codeforces information ascending order
	sort.SliceStable(info, func(i, j int) bool { return info[i].StartTime.Before(info[j].StartTime) })

	return info, nil
}

func (cs *crawlerService) FetchYukicoderInfo(ctx context.Context) ([]domain.YukicoderInfo, error) {
	info := make([]domain.YukicoderInfo, 0)
	// Call Yukicoder's future contests api
	body, err := cs.rr.Get(ctx, consts.YukicoderURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}
	err = cs.rr.DecodeJson(body, &info)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Decoding Function: %w", err)
	}

	return info, nil
}

// Arrange Scraped AtCoder Information
func (cs *crawlerService) arrangeAtcoderInfo(text string) []string {
	// Remove unnecessary elements
	text = strings.ReplaceAll(text, "◉", "")

	// Closure to split the scraped text
	f := func(c rune) bool {
		return !unicode.IsPrint(c)
	}

	return strings.FieldsFunc(text, f)
}
