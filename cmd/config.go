package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address  string `env:"ADDRESS" envDefault:":8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	HeaderJwtSourceEnabled bool   `env:"HEADER_JWT_SOURCE_ENABLED" envDefault:"true"`
	HeaderJwtSourceName    string `env:"HEADER_JWT_SOURCE_NAME" envDefault:"Authorization"`
	HeaderJwtSourcePrefix  string `env:"HEADER_JWT_SOURCE_PREFIX" envDefault:"Bearer "`

	CookieJwtSourceEnabled bool   `env:"COOKIE_JWT_SOURCE_ENABLED" envDefault:"false"`
	CookieJwtSourceName    string `env:"COOKIE_JWT_SOURCE_NAME" envDefault:"token"`

	ClaimMapping map[string]string `env:"CLAIM_MAPPINGS" envDefault:"sub:x-jwt-user-id,iss:x-jwt-issuer"`
}

func mapParser(v string) (interface{}, error) {
	m := make(map[string]string)
	l := strings.Split(v, ",")

	for _, s := range l {
		kv := strings.Split(s, ":")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid mapping: %s", s)
		}
		m[kv[0]] = kv[1]
	}

	return m, nil
}

func loadConfig() (*Config, error) {
	config := &Config{}
	options := env.Options{
		RequiredIfNoDef: true,
	}
	err := env.ParseWithFuncs(config, map[reflect.Type]env.ParserFunc{
		reflect.TypeOf(map[string]string{}): mapParser,
	}, options)

	if err != nil {
		return nil, err
	}

	return config, nil
}
