package envs

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	// Success Pattern
	want := &Env{
		Host:          "localhost",
		UserName:      "postgres",
		Password:      "password",
		Port:          "5432",
		Name:          "test",
		ChannelSecret: "dummy_channel_secret",
		ChannelToken:  "dummy_channel_token",
	}
	got, err := LoadEnv()
	assert.Nil(t, err)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("LoadEnv() returned invalid results (-got +want):\n %s", diff)
	}

	// When the environmental variable lacks
	os.Unsetenv("DB_USER_NAME")
	got, err = LoadEnv()
	assert.Error(t, err)
}
