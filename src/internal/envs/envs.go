package envs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Env struct {
	URL           string `envconfig:"DB_URL" required:"true"`
	ChannelSecret string `envconfig:"LINE_CHANNEL_SECRET" required:"true"`
	ChannelToken  string `envconfig:"LINE_CHANNEL_TOKEN" required:"true"`
}

func LoadEnv() (*Env, error) {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &env, nil
}
