package infrastructure

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/go-cmp/cmp"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/stretchr/testify/assert"
)

var jst = time.FixedZone("Azia/Tokyo", 9*60*60)

func TestNewConn(t *testing.T) {
	// The case failed to load envs
	os.Unsetenv("DB_USER_NAME")
	_, err := newConn()
	assert.Error(t, err)

	os.Setenv("DB_USER_NAME", "postgres")

	// The case success
	conn, err := newConn()
	assert.Nil(t, err)
	assert.NotNil(t, conn)
}

func TestInsertAtcoder(t *testing.T) {
	tableName := "atcoder"
	conn := dbConn()
	defer conn.Close()

	// Assume no data are inserted
	assert.Equal(t, 0, countRows(conn, tableName))

	arg := domain.AtcoderInfo{
		Name:       "AtCoder Beginner Contest 999",
		StartTime:  time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
		EndTime:    time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
		RatedRange: " ~ 1999",
	}

	cp := &contestPersistence{Conn: conn}
	err := cp.InsertAtcoder(arg)

	// Assume there's no error
	assert.Nil(t, err)

	// Assume there is one row
	rows := selectRows(conn, tableName)
	defer rows.Close()
	got, err := cp.convertRows(rows, "AtCoder")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 1, len(got))
	if diff := cmp.Diff(got[0], arg); diff != "" {
		t.Errorf("contestPersistence.InsertAtcoder() returned invalid results (-got +want):\n %s", diff)
	}

	resetTable(conn)
}

func TestInsertCodeforces(t *testing.T) {
	tableName := "codeforces"
	conn := dbConn()
	defer conn.Close()

	// Assume no data are inserted
	assert.Equal(t, 0, countRows(conn, tableName))

	arg := domain.CodeforcesInfo{
		Name:      "Codeforces Contest 999",
		StartTime: time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
		EndTime:   time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
	}

	cp := &contestPersistence{Conn: conn}
	err := cp.InsertCodeforces(arg)

	// Assume there's no error
	assert.Nil(t, err)

	// Assume there is one row
	rows := selectRows(conn, tableName)
	defer rows.Close()
	got, err := cp.convertRows(rows, "Codeforces")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 1, len(got))
	if diff := cmp.Diff(got[0], arg); diff != "" {
		t.Errorf("contestPersistence.InsertCodeforces() returned invalid results (-got +want):\n %s", diff)
	}

	resetTable(conn)
}

func TestInsertYukicoder(t *testing.T) {
	tableName := "yukicoder"
	conn := dbConn()
	defer conn.Close()

	// Assume no data are inserted
	assert.Equal(t, 0, countRows(conn, tableName))

	arg := domain.YukicoderInfo{
		Name:      "Codeforces Contest 999",
		StartTime: time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
		EndTime:   time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
	}

	cp := &contestPersistence{Conn: conn}
	err := cp.InsertYukicoder(arg)

	// Assume there's no error
	assert.Nil(t, err)

	// Assume there is one row
	rows := selectRows(conn, tableName)
	defer rows.Close()
	got, err := cp.convertRows(rows, "Yukicoder")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 1, len(got))
	if diff := cmp.Diff(got[0], arg); diff != "" {
		t.Errorf("contestPersistence.InsertYukicoder() returned invalid results (-got +want):\n %s", diff)
	}

	resetTable(conn)
}

func TestBatchGet(t *testing.T) {
	tests := []struct {
		name           string
		arg            string
		preDataset     string
		wantAtcoder    []domain.AtcoderInfo
		wantCodeforces []domain.CodeforcesInfo
		wantYukicoder  []domain.YukicoderInfo
	}{
		{
			name:       "AtCoder",
			arg:        "AtCoder",
			preDataset: "../tests/atcoder_predataset.sql",
			wantAtcoder: []domain.AtcoderInfo{
				{
					Name:       "AtCoder Beginner Contest 999",
					StartTime:  time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
					EndTime:    time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
					RatedRange: " ~ 1999",
				},
				{
					Name:       "AtCoder Regular Contest 999",
					StartTime:  time.Date(2021, 12, 31, 21, 0, 0, 0, jst),
					EndTime:    time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					RatedRange: " - ",
				},
			},
		},
		{
			name:       "Codeforces",
			arg:        "Codeforces",
			preDataset: "../tests/codeforces_predataset.sql",
			wantCodeforces: []domain.CodeforcesInfo{
				{
					Name:      "Codeforces 999",
					StartTime: time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
					EndTime:   time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
				},
			},
		},
		{
			name:       "Yukicoder",
			arg:        "Yukicoder",
			preDataset: "../tests/yukicoder_predataset.sql",
			wantYukicoder: []domain.YukicoderInfo{
				{
					Name:      "Yukicoder Contest 999",
					StartTime: time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
					EndTime:   time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := dbConn()
			execFile(conn, tt.preDataset)

			var want []interface{}
			switch tt.arg {
			case "AtCoder":
				for _, w := range tt.wantAtcoder {
					want = append(want, w)
				}
			case "Codeforces":
				for _, w := range tt.wantCodeforces {
					want = append(want, w)
				}
			case "Yukicoder":
				for _, w := range tt.wantYukicoder {
					want = append(want, w)
				}
			}

			cp := &contestPersistence{Conn: conn}
			got, err := cp.BatchGet(tt.arg)
			assert.Nil(t, err)

			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("contestsPersistence.BatchGet() returned invalid results (-got +want):\n %s", diff)
			}

			resetTable(conn)
		})
	}
}

func TestConvertRows(t *testing.T) {
	tests := []struct {
		name           string
		platform       string
		preDataset     string
		wantAtcoder    []domain.AtcoderInfo
		wantCodeforces []domain.CodeforcesInfo
		wantYukicoder  []domain.YukicoderInfo
	}{
		{
			name:       "AtCoder",
			platform:   "AtCoder",
			preDataset: "../tests/atcoder_predataset.sql",
			wantAtcoder: []domain.AtcoderInfo{
				{
					Name:       "AtCoder Beginner Contest 999",
					StartTime:  time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
					EndTime:    time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
					RatedRange: " ~ 1999",
				},
				{
					Name:       "AtCoder Regular Contest 999",
					StartTime:  time.Date(2021, 12, 31, 21, 0, 0, 0, jst),
					EndTime:    time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					RatedRange: " - ",
				},
			},
		},
		{
			name:       "Codeforces",
			platform:   "Codeforces",
			preDataset: "../tests/codeforces_predataset.sql",
			wantCodeforces: []domain.CodeforcesInfo{
				{
					Name:      "Codeforces 999",
					StartTime: time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
					EndTime:   time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
				},
			},
		},
		{
			name:       "Yukicoder",
			platform:   "Yukicoder",
			preDataset: "../tests/yukicoder_predataset.sql",
			wantYukicoder: []domain.YukicoderInfo{
				{
					Name:      "Yukicoder Contest 999",
					StartTime: time.Date(2021, 12, 30, 21, 0, 0, 0, jst),
					EndTime:   time.Date(2021, 12, 30, 22, 40, 0, 0, jst),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := dbConn()
			execFile(conn, tt.preDataset)
			rows := selectRows(conn, tt.platform)
			defer rows.Close()

			cp := &contestPersistence{}
			got, err := cp.convertRows(rows, tt.platform)
			assert.Nil(t, err)

			var want []interface{}
			switch tt.platform {
			case "AtCoder":
				for _, w := range tt.wantAtcoder {
					want = append(want, w)
				}
			case "Codeforces":
				for _, w := range tt.wantCodeforces {
					want = append(want, w)
				}
			case "Yukicoder":
				for _, w := range tt.wantYukicoder {
					want = append(want, w)
				}
			}
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("contestPersistence.convertRows() returned invalid results (-got +want):\n %s", diff)
			}

			resetTable(conn)
		})
	}
}

// -----------------
// Utility Functions
// -----------------

func dbConn() *sql.DB {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func execFile(conn *sql.DB, fileName string) {
	query, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(string(query))
}

// Reset table for the next test
func resetTable(conn *sql.DB) {
	execFile(conn, "../../config/init.sql")
}

// Select all data from the specified table
func selectRows(conn *sql.DB, table string) *sql.Rows {
	stmt := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := conn.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

// Count the number of rows
func countRows(conn *sql.DB, table string) int {
	var count int
	stmt := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	err := conn.QueryRow(stmt).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

// --------------------------
// Test For Utility Functions
// --------------------------

func TestExecFile(t *testing.T) {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM atcoder").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 0, count)

	execFile(db, "../tests/atcoder_predataset.sql")

	err = db.QueryRow("SELECT COUNT(*) FROM atcoder").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 2, count)

	// Clear Table
	_, err = db.Exec("DELETE FROM atcoder")
	if err != nil {
		log.Fatal(err)
	}
}

func TestResetTable(t *testing.T) {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Insert Data
	name := "AtCoder Beginner Contests 999"
	startTime := time.Now()
	endTime := time.Now()
	ratedRange := " ~ 1999"
	_, err = db.Query("INSERT INTO atcoder(name, start_time, end_time, range) VALUES($1, $2, $3, $4)", name, startTime, endTime, ratedRange)

	// Check the data is correctly inserted
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM atcoder").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 1, count)

	resetTable(db)

	// Check the tables are reseted
	err = db.QueryRow("SELECT COUNT(*) FROM atcoder").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 0, count)
}

func TestSelectRows(t *testing.T) {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Insert Data
	name := "AtCoder Beginner Contests 999"
	startTime := time.Now()
	endTime := time.Now()
	ratedRange := " ~ 1999"
	_, err = db.Query("INSERT INTO atcoder(name, start_time, end_time, range) VALUES($1, $2, $3, $4)", name, startTime, endTime, ratedRange)

	assert.NotNil(t, selectRows(db, "atcoder"))

	// Clear Table
	_, err = db.Exec("DELETE FROM atcoder")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCountRows(t *testing.T) {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=password dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Insert Data
	name := "AtCoder Beginner Contests 999"
	startTime := time.Now()
	endTime := time.Now()
	ratedRange := " ~ 1999"
	_, err = db.Query("INSERT INTO atcoder(name, start_time, end_time, range) VALUES($1, $2, $3, $4)", name, startTime, endTime, ratedRange)

	assert.Equal(t, 1, countRows(db, "atcoder"))

	// Clear Table
	_, err = db.Exec("DELETE FROM atcoder")
	if err != nil {
		log.Fatal(err)
	}
}
