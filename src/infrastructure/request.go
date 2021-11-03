package infrastructure

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"golang.org/x/xerrors"
)

type requestPersistence struct {
	client *http.Client
}

func NewRequestPersistence(client *http.Client) repository.RequestRepository {
	if client == nil {
		client = http.DefaultClient
	}

	return &requestPersistence{client: client}
}

func (rp *requestPersistence) Get(
	ctx context.Context,
	url string,
) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, xerrors.Errorf("Error when Creating Request: %w", err)
	}
	req = req.WithContext(ctx)

	// Make http request
	res, err := rp.client.Do(req)
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

func (rp *requestPersistence) DecodeJson(
	body []byte,
	target interface{},
) error {
	if err := json.Unmarshal(body, &target); err != nil {
		return xerrors.Errorf("Error when Decoding Json: %w", err)
	}

	return nil
}
