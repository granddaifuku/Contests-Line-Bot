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
	var info []domain.AtcoderInfo
	// req, err := http.NewRequest("GET", consts.AtcoderURL, nil)
	// if err != nil {
	// 	return nil, xerrors.Errorf("Error when Creating Request: %w", err)
	// }
	// res, err := cs.client.Do(req)
	// if err != nil {
	// 	return nil, xerrors.Errorf("Error when Making Http Request: %w", err)
	// }
	// defer res.Body.Close()
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
	var info domain.YukicoderInfo
	// Call Yukicoder's future contests api
	body, err := cs.makeGetRequest(consts.YukicoderURL)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Api: %w", err)
	}
	err = cs.decodeJson(body, &info)
	if err != nil {
		return nil, xerrors.Errorf("Error when Calling Decoding Function: %w", err)
	}
	//	cs.formatYukicoderInfo(info)

	return info, nil
}

// Convert Time format from ISO to 2006-01-02 15:04:05
// func (cs *contestService) formatYukicoderInfo(info domain.YukicoderInfo) error {
// 	for _, v := range info {
// 		startTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprint(v.StartTime))
// 		if err != nil {
// 			return xerrors.Errorf("Error when Parsing Converting Start Time Format: %w", err)
// 		}
// 		endTime, err := time.Parse("2006-01-02 15:04:05", fmt.Sprint(v.EndTime))
// 		if err != nil {
// 			return xerrors.Errorf("Error when Parsing Converting End Time Format: %w", err)
// 		}
// 		v.StartTime = startTime
// 		v.EndTime = endTime
// 	}
// 	return nil
// }

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
	strings.ReplaceAll(text, "◉", "")

	// Closure to split the scraped text
	f := func(c rune) bool {
		return !unicode.IsPrint(c)
	}

	return strings.FieldsFunc(text, f)
}
