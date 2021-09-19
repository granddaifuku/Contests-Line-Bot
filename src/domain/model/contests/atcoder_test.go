package contests

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewAtCoderInfo(t testing.T) {
	type args struct {
		name            string
		startTime       time.Time
		endTime         time.Time
		ratedRangeBegin int
		ratedRangeEnd   int
	}
	tests := []struct {
		name string
		args args
		want AtcoderInfo
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAtCoderInfo(tt.args.name, tt.args.startTime, tt.args.endTime, tt.args.ratedRangeBegin, tt.args.ratedRangeEnd)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NewAtcoderInfo() Failed (-want +got)\n%s", diff)
			}
		})
	}
}
