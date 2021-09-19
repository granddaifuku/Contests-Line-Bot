package service

import (
	"encoding/json"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/granddaifuku/contest_line_bot/src/internal/consts"
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

func (cs contestService) CallGetApi(url string, body *interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// Make http request
	res, err := cs.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil
	}
	// Decode json to struct
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return err
	}

	return nil
}

func (cs contestService) FetchAtcoderInfo() ([]domain.AtcoderInfo, error) {
	var info []domain.AtcoderInfo
	req, err := http.NewRequest("GET", consts.AtcoderURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// Scrape AtCoder's contests inforamtion
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	// Dive into the upcoming contests inforamtion
	doc.Find("div#contest-table-upcoming > div.panel > div.table-responsive > table.table > tbody > tr").Each(func(_ int, s *goquery.Selection) {

	})

	return info, nil
}

func (cs contestService) FetchCodeforcesInfo() ([]domain.CodeforcesInfo, error) {
	var info []domain.CodeforcesInfo
	// Call Codeforces' contests information api
	req, err := http.NewRequest("GET", consts.CodeforcesURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return info, nil
}

func (cs contestService) FetchYukicoderInfo() ([]domain.YukicoderInfo, error) {
	var info []domain.YukicoderInfo
	// Call Yukicoder's future contests api
	req, err := http.NewRequest("GET", consts.YukicoderURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return info, nil
}
