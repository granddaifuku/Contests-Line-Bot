package repository

import domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"

type ContestRepository interface {
	InsertAtcoder(info domain.AtcoderInfo) error
	InsertCodeforces(info domain.CodeforcesInfo) error
	InsertYukicoder(info domain.YukicoderInfo) error
	BatchGet(platform string) ([]interface{}, error)
}
