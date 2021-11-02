package domain

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewCodeforcesInfo(t *testing.T) {
	api := CodeforcesApiResult{Name: "Test Contests", Phase: "Finished", DurationSeconds: 3723, StartTimeSeconds: 1630454462}
	expected := CodeforcesInfo{Name: "Test Contests", StartTime: time.Date(2021, time.September, 1, 0, 1, 2, 0, time.UTC).In(jst), EndTime: time.Date(2021, time.September, 1, 1, 3, 5, 0, time.UTC).In(jst)}

	actual := NewCodeforcesInfo(&api)
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("NewCodeforcesInfo() Failed (-expected +actual):\n%s", diff)
	}
}
