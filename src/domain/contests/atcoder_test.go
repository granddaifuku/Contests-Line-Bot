package domain

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewAtcoderInfo(t *testing.T) {
	type args struct {
		name       string
		start      string
		duration   string
		ratedRange string
	}
	tests := []struct {
		name    string
		args    args
		want    AtcoderInfo
		wantErr error
	}{
		{
			name: "Success",
			args: args{
				name:       "AtCoder Begineer Contest 999",
				start:      "2021-12-31 21:00:00+0900",
				duration:   "100:100",
				ratedRange: "~ 1999",
			},
			want: AtcoderInfo{
				Name:       "AtCoder Begineer Contest 999",
				StartTime:  time.Date(2021, 12, 31, 21, 0, 0, 0, jst),
				EndTime:    time.Date(2022, 1, 5, 2, 40, 0, 0, jst),
				RatedRange: "~ 1999",
			},
			wantErr: nil,
		},
		{
			name: "Fail: No timezone suffix",
			args: args{
				name:       "AtCoder Begineer Contest 999",
				start:      "2021-12-31 21:00:00",
				duration:   "100:100",
				ratedRange: "~ 1999",
			},
			wantErr: nil,
		},
		{
			name: "Fail: Unable to parse time",
			args: args{
				name:       "AtCoder Begineer Contest 999",
				start:      "2021/12/31 21:00:00+0900",
				duration:   "100:100",
				ratedRange: "~ 1999",
			},
			wantErr: nil,
		},
		{
			name: "Fail: Unable to convert hours to int",
			args: args{
				name:       "AtCoder Begineer Contest 999",
				start:      "2021-12-31 21:00:00+0900",
				duration:   "ab:100",
				ratedRange: "~ 1999",
			},
			wantErr: nil,
		},
		{
			name: "Fail: Unable to convert minutes to int",
			args: args{
				name:       "AtCoder Begineer Contest 999",
				start:      "2021-12-31 21:00:00+0900",
				duration:   "100:ab",
				ratedRange: "~ 1999",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAtCoderInfo(tt.args.name, tt.args.start, tt.args.duration, tt.args.ratedRange)
			if err != nil {
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NewAtcoderInfo() Failed (-want +got):\n%s", diff)
			}
		})
	}
}
