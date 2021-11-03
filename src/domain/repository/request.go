package repository

import "context"

type RequestRepository interface {
	Get(
		ctx context.Context,
		url string,
	) ([]byte, error)

	DecodeJson(
		body []byte,
		target interface{},
	) error
}
