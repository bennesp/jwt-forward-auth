package main

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`

	HeaderJwtSource       bool   `env:"HEADER_JWT_SOURCE" envDefault:"true"`
	HeaderJwtSourceName   string `env:"HEADER_JWT_SOURCE_NAME" envDefault:"Authorization"`
	HeaderJwtSourcePrefix string `env:"HEADER_JWT_SOURCE_PREFIX" envDefault:"Bearer "`

	CookieJwtSource     bool   `env:"COOKIE_JWT_SOURCE" envDefault:"false"`
	CookieJwtSourceName string `env:"COOKIE_JWT_SOURCE_NAME" envDefault:"token"`
}

func loadConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config, env.Options{
		RequiredIfNoDef: true,
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}
