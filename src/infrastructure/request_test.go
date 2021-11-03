package infrastructure

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

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
			rp := &requestPersistence{}
			err := rp.DecodeJson(tt.args.body, &tt.args.target)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			if diff := cmp.Diff(tt.args.target, tt.want); diff != "" {
				t.Errorf("requestPersistence.DecodeJson() do not work properly (-got, +want):\n %s", diff)
			}
		})
	}
}
