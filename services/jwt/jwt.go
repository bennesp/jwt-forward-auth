package jwt

import (
	"errors"
	"net/url"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

// JwtWrapper is a simple wrapper service over golang-jwt/jwt and MicahParks/keyfunc
type JwtWrapper struct {
	config  *Config
	keyFunc jwt.Keyfunc
}

func New() (*JwtWrapper, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	if config.JwtSecretEnabled && config.JwtSecret == "" {
		return nil, errors.New("jwt secret is enabled but no secret is set")
	}

	if config.JwksEnabled && config.JwksUrl == "" {
		return nil, errors.New("jwks is enabled but no jwks url is set")
	}

	keyFunc, err := getKeyFunc(config)
	if err != nil {
		return nil, err
	}

	return &JwtWrapper{
		config:  config,
		keyFunc: keyFunc,
	}, nil
}

func isValidUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getKeyFunc(config *Config) (jwt.Keyfunc, error) {
	if !config.JwksEnabled && !config.JwtSecretEnabled {
		return nil, errors.New("no secret or jwks enabled")
	}

	if !config.JwksEnabled || !isValidUrl(config.JwksUrl) {
		log.Warn("Jwks not enabled or invalid URL, fallback using only JWT_SECRET")
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		}

		return keyFunc, nil
	}

	jwksOptions := keyfunc.Options{
		RefreshErrorHandler: func(err error) {
			log.Errorf("Error refreshing JWKS: %v", err)
		},
		RefreshInterval:   config.JwksRefreshInterval,
		RefreshRateLimit:  config.JwksRefreshRateLimit,
		RefreshTimeout:    config.JwksRefreshTimeout,
		RefreshUnknownKID: config.JwksRefreshUnknownKID,
	}

	jwks, err := keyfunc.Get(config.JwksUrl, jwksOptions)
	if err != nil {
		return nil, err
	}

	return func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Header["kid"]; !ok {
			if !config.JwtSecretEnabled {
				return nil, errors.New("no 'kid' found in jwt and jwt secret disabled, cannot validate")
			}

			return []byte(config.JwtSecret), nil
		}

		return jwks.Keyfunc(t)
	}, nil
}

func (jwtWrapper *JwtWrapper) Verify(jwtAsString string) (*jwt.Token, error) {
	log.WithField("jwt", jwtAsString).Debugf("Verifying JWT")

	token, err := jwt.Parse(jwtAsString, jwtWrapper.keyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (jwtWrapper *JwtWrapper) GetClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}
