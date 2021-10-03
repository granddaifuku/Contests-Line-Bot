package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/granddaifuku/contest_line_bot/src/domain/repository"
	"github.com/granddaifuku/contest_line_bot/src/internal/envs"
	"golang.org/x/xerrors"
)

type contestPersistence struct {
	Conn *sql.DB
}

func NewContestPersistence(conn *sql.DB) repository.ContestRepository {
	return &contestPersistence{Conn: conn}
}

func newConn() (*sql.DB, error) {
	env, err := envs.LoadEnv()
	if err != nil {
		return nil, xerrors.Errorf("Error when Loading Envs: %w", err)
	}
	dbInfo := fmt.Sprintf("host=postgres port=%s user=%s password=%s dbname=%s", env.Port, env.UserName, env.Password, env.Name)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, xerrors.Errorf("Error when Connecting DB: %w", err)
	}

	return db, nil
}

func (cp *contestPersistence) InsertAtcoder(info domain.AtcoderInfo) error {
	_, err := cp.Conn.Exec("INSERT INTO atcoder(name, start_time, end_time, range) VALUES($1, $2, $3, $4)", info.Name, info.StartTime, info.EndTime, info.RatedRange)
	if err != nil {
		return xerrors.Errorf("Error when Executing Statement for Inserting AtCoder Information: %w", err)
	}

	return nil
}

func (cp *contestPersistence) InsertCodeforces(info domain.CodeforcesInfo) error {
	_, err := cp.Conn.Exec("INSERT INTO codeforces(name, start_time, end_time) VALUES($1, $2, $3)", info.Name, info.StartTime, info.EndTime)
	if err != nil {
		return xerrors.Errorf("Error when Executing Statement for Inserting Codeforces Information: %w", err)
	}

	return nil
}

func (cp *contestPersistence) InsertYukicoder(info domain.YukicoderInfo) error {
	_, err := cp.Conn.Exec("INSERT INTO yukicoder(name, start_time, end_time) VALUES($1, $2, $3)", info.Name, info.StartTime, info.EndTime)
	if err != nil {
		return xerrors.Errorf("Error when Executing Statement for Inserting Yukicoder Information: %w", err)
	}

	return nil
}

// Select all data from the spceified table
func (cp *contestPersistence) BatchGet(platform string) ([]interface{}, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", platform)
	rows, err := cp.Conn.Query(stmt)
	if err != nil {
		return nil, xerrors.Errorf("Error when Selecting Table: %w", err)
	}

	return cp.convertRows(rows, platform)
}

// Convert rows to the specific struct
func (cp *contestPersistence) convertRows(rows *sql.Rows, platform string) ([]interface{}, error) {
	var id int
	info := []interface{}{}
	for rows.Next() {
		switch platform {
		case "AtCoder":
			tmp := domain.AtcoderInfo{}
			err := rows.Scan(&id, &tmp.Name, &tmp.StartTime, &tmp.EndTime, &tmp.RatedRange)
			if err != nil {
				return nil, xerrors.Errorf("Error when Scanning Rows: %w", err)
			}
			info = append(info, tmp)
		case "Codeforces":
			tmp := domain.CodeforcesInfo{}
			err := rows.Scan(&id, &tmp.Name, &tmp.StartTime, &tmp.EndTime)
			if err != nil {
				return nil, xerrors.Errorf("Error when Scanning Rows: %w", err)
			}
			info = append(info, tmp)
		case "Yukicoder":
			tmp := domain.YukicoderInfo{}
			err := rows.Scan(&id, &tmp.Name, &tmp.StartTime, &tmp.EndTime)
			if err != nil {
				return nil, xerrors.Errorf("Error when Scanning Rows: %w", err)
			}
			info = append(info, tmp)
		}
	}

	return info, nil
}
