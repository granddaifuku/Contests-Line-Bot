package envs

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

type Env struct {
	Host          string `envconfig:"DB_HOST" required:"true"`
	UserName      string `envconfig:"DB_USER_NAME" required:"true"`
	Password      string `envconfig:"DB_PASSWORD" required:"true"`
	Port          string `envconfig:"DB_PORT" required:"true"`
	Name          string `envconfig:"DB_NAME" required:"true"`
	ChannelSecret string `envconfig:"LINE_CHANNEL_SECRET" required:"true"`
	ChannelToken  string `envconfig:"LINE_CHANNEL_TOKEN" required:"true"`
}

func LoadEnv() (*Env, error) {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, xerrors.Errorf("Error when Processing Envs: %w", err)
	}

	return &env, nil
}
