package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	domain "github.com/granddaifuku/contest_line_bot/src/domain/model/contests"
	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/granddaifuku/contest_line_bot/src/internal/envs"
)

type databasePersistence struct {
	Conn *sql.DB
}

func NewDatabasePersistence(conn *sql.DB) repository.DatabaseRepository {
	if conn == nil {
		c, err := newConn()
		if err != nil {
			panic(err)
		}
		conn = c
		if err := conn.Ping(); err != nil {
			panic(err)
		}
	}
	return &databasePersistence{Conn: conn}
}

func newConn() (*sql.DB, error) {
	env, err := envs.LoadEnv()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db, err := sql.Open("postgres", env.URL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

func (dp *databasePersistence) InsertAtcoder(
	ctx context.Context,
	info domain.AtcoderInfo,
) error {
	_, err := dp.Conn.Exec("INSERT INTO atcoder(name, start_time, end_time, range) VALUES($1, $2, $3, $4)", info.Name, info.StartTime, info.EndTime, info.RatedRange)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (dp *databasePersistence) InsertCodeforces(
	ctx context.Context,
	info domain.CodeforcesInfo,
) error {
	_, err := dp.Conn.Exec("INSERT INTO codeforces(name, start_time, end_time) VALUES($1, $2, $3)", info.Name, info.StartTime, info.EndTime)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (dp *databasePersistence) InsertYukicoder(
	ctx context.Context,
	info domain.YukicoderInfo,
) error {
	_, err := dp.Conn.Exec("INSERT INTO yukicoder(name, start_time, end_time) VALUES($1, $2, $3)", info.Name, info.StartTime, info.EndTime)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Select all data from the spceified table
func (dp *databasePersistence) BatchGet(
	ctx context.Context,
	platform string,
) ([]interface{}, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", platform)
	rows, err := dp.Conn.Query(stmt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dp.convertRows(rows, platform)
}

// Clear Table
func (dp *databasePersistence) ClearTables(
	ctx context.Context,
) error {
	tables := []string{
		"atcoder",
		"codeforces",
		"yukicoder",
	}
	for _, table := range tables {
		stmt := fmt.Sprintf("DELETE FROM %s", table)
		_, err := dp.Conn.Exec(stmt)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// Convert rows to the specific struct
func (dp *databasePersistence) convertRows(
	rows *sql.Rows,
	platform string,
) ([]interface{}, error) {
	var id int
	info := []interface{}{}
	for rows.Next() {
		switch platform {
		case "AtCoder":
			tmp := domain.AtcoderInfo{}
			err := rows.Scan(&id, &tmp.Name, &tmp.StartTime, &tmp.EndTime, &tmp.RatedRange)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			info = append(info, tmp)
		case "Codeforces":
			tmp := domain.CodeforcesInfo{}
			err := rows.Scan(&id, &tmp.Name, &tmp.StartTime, &tmp.EndTime)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			info = append(info, tmp)
		case "Yukicoder":
			tmp := domain.YukicoderInfo{}
			err := rows.Scan(&id, &tmp.Name, &tmp.StartTime, &tmp.EndTime)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			info = append(info, tmp)
		}
	}

	return info, nil
}
