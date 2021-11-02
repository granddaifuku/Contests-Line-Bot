package repository

import (
	"context"

	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
)

type DatabaseRepository interface {
	InsertAtcoder(
		ctx context.Context,
		info domain.AtcoderInfo,
	) error

	InsertCodeforces(
		ctx context.Context,
		info domain.CodeforcesInfo,
	) error

	InsertYukicoder(
		ctx context.Context,
		info domain.YukicoderInfo,
	) error

	BatchGet(
		ctx context.Context,
		platform string,
	) ([]interface{}, error)

	ClearTables(
		ctx context.Context,
	) error
}
