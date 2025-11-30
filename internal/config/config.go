package config

import (
	"errors"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":3030"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, errors.New("failed to parse config")
	}
	return cfg, nil
}
