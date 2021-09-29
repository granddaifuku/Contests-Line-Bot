package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	FetchYukicoderInfo() (domain.YukicoderInfo, error)
}

type contestService struct {
	client *http.Client
}

func NewContestService(client *http.Client) ContestService {
	if client != nil {
		return &contestService{
			client: client,
		}
	}
	return &contestService{
		client: new(http.Client),
	}
}

func (cs *contestService) FetchAtcoderInfo() ([]domain.AtcoderInfo, error) {
	info := make([]domain.AtcoderInfo, 0)
	body, err := cs.makeGetRequest(consts.AtcoderURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}

	// Scrape AtCoder's contests inforamtion
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, xerrors.Errorf("Error when Reading Document: %w", err)
	}

	// Dive into the upcoming contests inforamtion
	scraped := doc.Find("div#contest-table-upcoming > div.panel > div.table-responsive > table.table > tbody > tr").Text()

	splited := cs.arrangeAtcoderInfo(scraped)
	// Devide the slice every 4 elements.
	for i := 0; i < len(splited); i += 4 {
		startTime := splited[i]
		name := splited[i+1]
		duration := splited[i+2]
		ratedRange := splited[i+3]
		at, err := domain.NewAtCoderInfo(name, startTime, duration, ratedRange)
		if err != nil {
			return nil, xerrors.Errorf("Error when Building AtCoder Info: %w", err)
		}
		info = append(info, at)
	}

	return info, nil
}

func (cs *contestService) FetchCodeforcesInfo() ([]domain.CodeforcesInfo, error) {
	api := domain.CodeforcesApi{}
	info := make([]domain.CodeforcesInfo, 0)
	// Call Codeforces' contests information api
	body, err := cs.makeGetRequest(consts.CodeforcesURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}
	err = cs.decodeJson(body, &api)
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

	return info, nil
}

func (cs *contestService) FetchYukicoderInfo() (domain.YukicoderInfo, error) {
	info := make(domain.YukicoderInfo, 0)
	// Call Yukicoder's future contests api
	body, err := cs.makeGetRequest(consts.YukicoderURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}
	err = cs.decodeJson(body, &info)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Decoding Function: %w", err)
	}

	return info, nil
}

func (cs *contestService) makeGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, xerrors.Errorf("Error when Creating Request: %w", err)
	}
	// Make http request
	res, err := cs.client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("Error when Making Http Request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("Http Status is not OK. Code: %v", res.StatusCode)

	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, xerrors.Errorf("Error when Reading Response Body: %w", err)
	}

	return body, nil
}

func (cs *contestService) decodeJson(body []byte, target interface{}) error {
	if err := json.Unmarshal(body, &target); err != nil {
		return xerrors.Errorf("Error when Decoding Json: %w", err)
	}

	return nil
}

// Arrange Scraped AtCoder Information
func (cs *contestService) arrangeAtcoderInfo(text string) []string {
	// Remove unnecessary elements
	text = strings.ReplaceAll(text, "◉", "")

	// Closure to split the scraped text
	f := func(c rune) bool {
		return !unicode.IsPrint(c)
	}

	return strings.FieldsFunc(text, f)
}
