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
		URL:           "postgresql://localhost:5432/test?user=postgres&password=password",
		ChannelSecret: "dummy_channel_secret",
		ChannelToken:  "dummy_channel_token",
	}
	got, err := LoadEnv()
	assert.Nil(t, err)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("LoadEnv() returned invalid results (-got +want):\n %s", diff)
	}

	// When the environmental variable lacks
	os.Unsetenv("DB_URL")
	_, err = LoadEnv()
	assert.Error(t, err)
}
