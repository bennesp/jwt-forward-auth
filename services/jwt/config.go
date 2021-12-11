package jwt

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	JwksEnabled           bool          `env:"JWKS_ENABLED" envDefault:"true"`
	JwksUrl               string        `env:"JWKS_URL" envDefault:""`
	JwksRefreshInterval   time.Duration `env:"JWKS_REFRESH_INTERVAL" envDefault:"1h"`
	JwksRefreshRateLimit  time.Duration `env:"JWKS_REFRESH_RATE_LIMIT" envDefault:"5m"`
	JwksRefreshTimeout    time.Duration `env:"JWKS_REFRESH_TIMEOUT" envDefault:"5s"`
	JwksRefreshUnknownKID bool          `env:"JWKS_REFRESH_UNKNOWN_KID" envDefault:"true"`

	JwtSecretEnabled bool   `env:"JWT_SECRET_ENABLED" envDefault:"false"`
	JwtSecret        string `env:"JWT_SECRET" envDefault:""`
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
