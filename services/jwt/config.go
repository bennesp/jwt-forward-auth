package jwt

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	JwksUrl               string        `env:"JWKS_URL" envDefault:""`
	JwksRefreshInterval   time.Duration `env:"JWKS_RefreshInterval" envDefault:"1h"`
	JwksRefreshRateLimit  time.Duration `env:"JWKS_RefreshRateLimit" envDefault:"5m"`
	JwksRefreshTimeout    time.Duration `env:"JWKS_RefreshTimeout" envDefault:"5s"`
	JwksRefreshUnknownKID bool          `env:"JWKS_RefreshUnknownKID" envDefault:"true"`

	JwtSecret string `env:"SECRET" envDefault:""`
}

func loadConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config, env.Options{
		RequiredIfNoDef: true,
		Prefix:          "JWT_",
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}
