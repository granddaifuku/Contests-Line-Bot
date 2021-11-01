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

type databasePersistence struct {
	Conn *sql.DB
}

func NewContestPersistence(conn *sql.DB) repository.DatabaseRepository {
	return &databasePersistence{Conn: conn}
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

func (dp *databasePersistence) InsertAtcoder(info domain.AtcoderInfo) error {
	_, err := dp.Conn.Exec("INSERT INTO atcoder(name, start_time, end_time, range) VALUES($1, $2, $3, $4)", info.Name, info.StartTime, info.EndTime, info.RatedRange)
	if err != nil {
		return xerrors.Errorf("Error when Executing Statement for Inserting AtCoder Information: %w", err)
	}

	return nil
}

func (dp *databasePersistence) InsertCodeforces(info domain.CodeforcesInfo) error {
	_, err := dp.Conn.Exec("INSERT INTO codeforces(name, start_time, end_time) VALUES($1, $2, $3)", info.Name, info.StartTime, info.EndTime)
	if err != nil {
		return xerrors.Errorf("Error when Executing Statement for Inserting Codeforces Information: %w", err)
	}

	return nil
}

func (dp *databasePersistence) InsertYukicoder(info domain.YukicoderInfo) error {
	_, err := dp.Conn.Exec("INSERT INTO yukicoder(name, start_time, end_time) VALUES($1, $2, $3)", info.Name, info.StartTime, info.EndTime)
	if err != nil {
		return xerrors.Errorf("Error when Executing Statement for Inserting Yukicoder Information: %w", err)
	}

	return nil
}

// Select all data from the spceified table
func (dp *databasePersistence) BatchGet(platform string) ([]interface{}, error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", platform)
	rows, err := dp.Conn.Query(stmt)
	if err != nil {
		return nil, xerrors.Errorf("Error when Selecting Table: %w", err)
	}

	return dp.convertRows(rows, platform)
}

// Clear Table
func (dp *databasePersistence) ClearTables() error {
	tables := []string{
		"atcoder",
		"codeforces",
		"yukicoder",
	}
	for _, table := range tables {
		stmt := fmt.Sprintf("DELETE FROM %s", table)
		_, err := dp.Conn.Exec(stmt)
		if err != nil {
			return xerrors.Errorf("Error when Deleting Table %s: %w", table, err)
		}
	}

	return nil
}

// Convert rows to the specific struct
func (dp *databasePersistence) convertRows(rows *sql.Rows, platform string) ([]interface{}, error) {
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