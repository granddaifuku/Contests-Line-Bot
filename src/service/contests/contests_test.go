package service

import (
	"net/http"
	"sort"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/v2/recorder"
	"github.com/google/go-cmp/cmp"
	domain "github.com/granddaifuku/contest_line_bot/src/domain/contests"
	"github.com/stretchr/testify/assert"
)

var jst = time.FixedZone("Azia/Tokyo", 9*60*60)

func TestFetchAtcoderInfo(t *testing.T) {
	r, err := recorder.New("../../../fixtures/service/contests/fetch_atcoder_info")
	assert.Nil(t, err)
	defer r.Stop()
	want := []domain.AtcoderInfo{
		{
			Name:       "AtCoder Beginner Contest 221",
			StartTime:  time.Date(2021, 10, 2, 21, 0, 0, 0, jst),
			EndTime:    time.Date(2021, 10, 2, 22, 40, 0, 0, jst),
			RatedRange: " ~ 1999",
		},
		{
			Name:       "エクサウィザーズプログラミングコンテスト2021（AtCoder Beginner Contest 222）",
			StartTime:  time.Date(2021, 10, 9, 21, 0, 0, 0, jst),
			EndTime:    time.Date(2021, 10, 9, 22, 40, 0, 0, jst),
			RatedRange: " ~ 1999",
		},
		{
			Name:       "デジタルの日特別イベント「HACK TO THE FUTURE for Youth+」",
			StartTime:  time.Date(2021, 10, 10, 13, 30, 0, 0, jst),
			EndTime:    time.Date(2021, 10, 10, 17, 30, 0, 0, jst),
			RatedRange: "-",
		},
		{
			Name:       "デジタルの日特別イベント「HACK TO THE FUTURE for Youth+」 open",
			StartTime:  time.Date(2021, 10, 10, 13, 30, 0, 0, jst),
			EndTime:    time.Date(2021, 10, 10, 17, 30, 0, 0, jst),
			RatedRange: "-",
		},
		{
			Name:       "大和証券プログラミングコンテスト2021（AtCoder Regular Contest 128）",
			StartTime:  time.Date(2021, 10, 16, 21, 0, 0, 0, jst),
			EndTime:    time.Date(2021, 10, 16, 23, 0, 0, 0, jst),
			RatedRange: " ~ 2799",
		},
		{
			Name:       "AtCoder Grand Contest 055",
			StartTime:  time.Date(2021, 10, 31, 21, 0, 0, 0, jst),
			EndTime:    time.Date(2021, 10, 31, 23, 30, 0, 0, jst),
			RatedRange: "1200 ~ ",
		},
	}
	cs := &contestService{client: &http.Client{Transport: r}}
	got, err := cs.FetchAtcoderInfo()
	assert.Nil(t, err)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("contestService.FetchAtcoderInfo() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestFetchCodeforcesInfo(t *testing.T) {
	r, err := recorder.New("../../../fixtures/service/contests/fetch_codeforces_info")
	assert.Nil(t, err)
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
	assert.Nil(t, err)
	sort.SliceStable(want, func(i, j int) bool { return want[i].Name < want[j].Name })
	sort.SliceStable(got, func(i, j int) bool { return got[i].Name < got[j].Name })

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("contestService.FetchCodeforcesInfo() returned invalid results (-got +want):\n %s", diff)
	}
}

func TestFetchYukicoderInfo(t *testing.T) {
	r, err := recorder.New("../../../fixtures/service/contests/fetch_yukicoder_info")
	assert.Nil(t, err)
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
	assert.Nil(t, err)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("contestService.FetchYukicoderInfo() returned invalid results (-got +want):\n %s", diff)
	}
}

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
			assert.Nil(t, err)
			defer r.Stop()
			cs := &contestService{client: &http.Client{Transport: r}}
			got, err := cs.makeGetRequest(tt.arg)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("contestService.makeGetRequest() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}

func TestDecodeJson(t *testing.T) {
	type Cat struct {
		Name     string   `json:"name"`
		Age      int      `json:"age"`
		Siblings []string `json:"siblings"`
	}
	type args struct {
		body   []byte
		target Cat
	}
	tests := []struct {
		name    string
		args    args
		want    Cat
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				body:   []byte(`{"name":"Haru","age":1,"siblings":["Hime"]}`),
				target: Cat{},
			},
			want: Cat{
				Name: "Haru",
				Age:  1,
				Siblings: []string{
					"Hime",
				},
			},
			wantErr: false,
		},
		{
			name: "Cannot unmarshal json data",
			args: args{
				body:   []byte(`{"name""Haru","age":1,"siblings":["Hime"]}`),
				target: Cat{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &contestService{}
			err := cs.decodeJson(tt.args.body, &tt.args.target)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			if diff := cmp.Diff(tt.args.target, tt.want); diff != "" {
				t.Errorf("contestService.decodeJson() do not work properly (-got, +want):\n %s", diff)
			}
		})
	}
}

func TestArrangeAtcoderInfo(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{
			name: "Success",
			arg: `2021-10-02 21:00:00+0900
				
					◉
					AtCoder Beginner Contest 221
				
				01:40
				 ~ 1999`,
			want: []string{
				"2021-10-02 21:00:00+0900",
				"AtCoder Beginner Contest 221",
				"01:40",
				" ~ 1999",
			},
		},
		{
			name: "No \"◉\" in the text",
			arg: `
			2021-10-02 21:00:00+0900
			
			AtCoder Beginner Contest 221
				
			01:40
				 ~ 1999`,
			want: []string{
				"2021-10-02 21:00:00+0900",
				"AtCoder Beginner Contest 221",
				"01:40",
				" ~ 1999",
			},
		},
		{
			name: "No delimiters in the text",
			arg:  "2021-10-02 21:00:00+0900◉AtCoder Beginner Contest 22101:40~ 1999",
			want: []string{
				"2021-10-02 21:00:00+0900AtCoder Beginner Contest 22101:40~ 1999",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &contestService{}
			got := cs.arrangeAtcoderInfo(tt.arg)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("contestsService.arrangeAtcoderInfo() returned invalid results (-got +want):\n %s", diff)
			}
		})
	}
}
