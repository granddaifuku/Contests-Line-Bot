package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
	"golang.org/x/xerrors"
)

type ContestService interface {
	FetchAtcoderInfo() ([]domain.AtcoderInfo, error)
	FetchCodeforcesInfo() ([]domain.CodeforcesInfo, error)
	FetchYukicoderInfo() ([]domain.YukicoderInfo, error)
}

type contestService struct {
	client *http.Client
}

func NewContestService() ContestService {
	client := new(http.Client)
	return &contestService{
		client: client,
	}
}

func (cs contestService) callGetApi(url string, body interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return xerrors.Errorf("Error when Creating Request: %w", err)
	}
	// Make http request
	res, err := cs.client.Do(req)
	if err != nil {
		return xerrors.Errorf("Error when Making Http Request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return xerrors.Errorf("Http Status is not OK. Code: %v", res.StatusCode)

	}
	// Decode json to struct
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return xerrors.Errorf("Error when Decoding Json: %w", err)
	}

	return nil
}

func (cs contestService) FetchAtcoderInfo() ([]domain.AtcoderInfo, error) {
	var info []domain.AtcoderInfo
	req, err := http.NewRequest("GET", consts.AtcoderURL, nil)
	if err != nil {
		return nil, xerrors.Errorf("Error when Creating Request: %w", err)
	}
	res, err := cs.client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("Error when Making Http Request: %w", err)
	}
	defer res.Body.Close()

	// Scrape AtCoder's contests inforamtion
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, xerrors.Errorf("Error when Reading Document: %w", err)
	}

	// Dive into the upcoming contests inforamtion
	scraped := doc.Find("div#contest-table-upcoming > div.panel > div.table-responsive > table.table > tbody > tr").Text()
	strings.ReplaceAll(scraped, "◉", "") // Remove unnecessary elements
	// Closure to split the scraped text
	f := func(c rune) bool {
		return !unicode.IsPrint(c)
	}
	splited := strings.FieldsFunc(scraped, f)
	for i := range splited {
		// Devide the slice every 4 elements.
		startTime := splited[i]
		name := splited[i+1]
		duration := splited[i+2]
		ratedRange := splited[i+3]
		at, err := domain.NewAtCoderInfo(name, startTime, duration, ratedRange)
		if err != nil {
			return nil, xerrors.Errorf("Error when Building AtCoder Info: %w", err)
		}
		info = append(info, at)
		i += 3
	}

	return info, nil
}

func (cs contestService) FetchCodeforcesInfo() ([]domain.CodeforcesInfo, error) {
	var info []domain.CodeforcesInfo
	// Call Codeforces' contests information api
	err := cs.callGetApi(consts.CodeforcesURL, &info)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}

	return info, nil
}

func (cs contestService) FetchYukicoderInfo() ([]domain.YukicoderInfo, error) {
	var info []domain.YukicoderInfo
	// Call Yukicoder's future contests api
	err := cs.callGetApi(consts.YukicoderURL, &info)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}

	return info, nil
}
