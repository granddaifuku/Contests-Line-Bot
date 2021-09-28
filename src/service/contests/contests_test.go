package service

import (
	"net/http"
	"sort"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/v2/recorder"
	"github.com/google/go-cmp/cmp"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
)

var jst = time.FixedZone("Azia/Tokyo", 9*60*60)

func TestMakeGetRequest(t *testing.T) {
	tests := []struct {
		name        string
		arg         string
		fixturePath string
		want        []byte
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := recorder.New("../../../fixtures/service/contests/" + tt.fixturePath)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Stop()
			cs := &contestService{client: &http.Client{Transport: r}}
			got, err := cs.makeGetRequest(tt.arg)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("contestService.makeGetRequest() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}

func TestFetchCodeforcesInfo(t *testing.T) {
	r, err := recorder.New("../../../fixtures/service/contests/fetch_codeforces_info")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	want := []domain.CodeforcesInfo{
		{
			Name:      "Codeforces Round #744 (Div. 3)",
			StartTime: time.Date(2021, 9, 28, 23, 35, 0, 0, jst),
			EndTime:   time.Date(2021, 9, 29, 1, 50, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #745 (Div. 1)",
			StartTime: time.Date(2021, 9, 30, 19, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 9, 30, 21, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #745 (Div. 2)",
			StartTime: time.Date(2021, 9, 30, 19, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 9, 30, 21, 5, 0, 0, jst),
		},
		{
			Name:      "ICPC WF Moscow Invitational Contest - Online Mirror (Unrated, ICPC Rules, Teams Preferred)",
			StartTime: time.Date(2021, 10, 1, 22, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 2, 3, 5, 0, 0, jst),
		},
		{
			Name:      "Kotlin Heroes: Practice 8",
			StartTime: time.Date(2021, 10, 1, 22, 35, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 8, 22, 35, 0, 0, jst),
		},
		{
			Name:      "COMPFEST 13 - Finals Online Mirror (Unrated, ICPC Rules, Teams Preferred)",
			StartTime: time.Date(2021, 10, 2, 22, 35, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 3, 3, 35, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #746 (Div. 2)",
			StartTime: time.Date(2021, 10, 3, 23, 35, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 4, 1, 35, 0, 0, jst),
		},
		{
			Name:      "Kotlin Heroes: Episode 8",
			StartTime: time.Date(2021, 10, 7, 23, 35, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 8, 2, 5, 0, 0, jst),
		},
		{
			Name:      "2021 ICPC Communication Routing Challenge: Marathon",
			StartTime: time.Date(2021, 10, 9, 9, 0, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 13, 9, 0, 0, 0, jst),
		},
		{
			Name:      "Technocup 2022 - Elimination Round 1",
			StartTime: time.Date(2021, 10, 17, 20, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 17, 22, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 1, based on Technocup 2022 Elimination Round 1)",
			StartTime: time.Date(2021, 10, 17, 20, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 17, 22, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 2, based on Technocup 2022 Elimination Round 1)",
			StartTime: time.Date(2021, 10, 17, 20, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 17, 22, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 2)",
			StartTime: time.Date(2021, 10, 24, 19, 35, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 24, 21, 35, 0, 0, jst),
		},
		{
			Name:      "Technocup 2022 - Elimination Round 2",
			StartTime: time.Date(2021, 11, 14, 15, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 11, 14, 17, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 1, based on Technocup 2022 Elimination Round 2)",
			StartTime: time.Date(2021, 11, 14, 15, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 11, 14, 17, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 2, based on Technocup 2022 Elimination Round 2)",
			StartTime: time.Date(2021, 11, 14, 15, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 11, 14, 17, 5, 0, 0, jst),
		},
		{
			Name:      "Technocup 2022 - Elimination Round 3",
			StartTime: time.Date(2021, 12, 13, 0, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 12, 13, 2, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 1, based on Technocup 2022 Elimination Round 3)",
			StartTime: time.Date(2021, 12, 13, 0, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 12, 13, 2, 5, 0, 0, jst),
		},
		{
			Name:      "Codeforces Round #TBA (Div. 2, based on Technocup 2022 Elimination Round 3)",
			StartTime: time.Date(2021, 12, 13, 0, 5, 0, 0, jst),
			EndTime:   time.Date(2021, 12, 13, 2, 5, 0, 0, jst),
		},
	}
	cs := &contestService{client: &http.Client{Transport: r}}
	got, err := cs.FetchCodeforcesInfo()
	sort.SliceStable(want, func(i, j int) bool { return want[i].Name < want[j].Name })
	sort.SliceStable(got, func(i, j int) bool { return got[i].Name < got[j].Name })
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("contestService.FetchCodeforcesInfo() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestFetchYukicoderInfo(t *testing.T) {
	r, err := recorder.New("../../../fixtures/service/contests/fetch_yukicoder_info")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	want := domain.YukicoderInfo{
		{
			Name:      "yukicoder contest",
			StartTime: time.Date(2021, 10, 1, 21, 20, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 1, 23, 20, 0, 0, jst),
		},
		{
			Name:      "yukicoder contest",
			StartTime: time.Date(2021, 10, 8, 21, 20, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 8, 23, 20, 0, 0, jst),
		},
		{
			Name:      "yukicoder contest",
			StartTime: time.Date(2021, 10, 15, 21, 20, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 15, 23, 20, 0, 0, jst),
		},
		{
			Name:      "yukicoder contest",
			StartTime: time.Date(2021, 10, 22, 21, 20, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 22, 23, 20, 0, 0, jst),
		},
		{
			Name:      "yukicoder contest",
			StartTime: time.Date(2021, 10, 29, 21, 20, 0, 0, jst),
			EndTime:   time.Date(2021, 10, 29, 23, 20, 0, 0, jst),
		},
		{
			Name:      "yukicoder contest （オムニバス 6問程度問題募集中）",
			StartTime: time.Date(2021, 11, 5, 21, 20, 0, 0, jst),
			EndTime:   time.Date(2021, 11, 5, 23, 20, 0, 0, jst),
		},
	}
	cs := &contestService{client: &http.Client{Transport: r}}
	got, err := cs.FetchYukicoderInfo()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("contestService.FetchYukicoderInfo() returned invalid results (-got +want):\n %s", diff)
	}
}
