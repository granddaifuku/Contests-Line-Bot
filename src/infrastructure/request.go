package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/pkg/errors"
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
		return nil, errors.WithStack(err)
	}
	req = req.WithContext(ctx)

	// Make http request
	res, err := rp.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Http Status is not OK. Code: %v", res.StatusCode))

	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}

func (rp *requestPersistence) DecodeJson(
	body []byte,
	target interface{},
) error {
	if err := json.Unmarshal(body, &target); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
