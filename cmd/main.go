package main

import (
	"github.com/bennesp/traefik-jwt-forward-auth/server"
	"github.com/bennesp/traefik-jwt-forward-auth/services/jwt"
	"github.com/bennesp/traefik-jwt-forward-auth/sources"

	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.WithError(err).Fatal("Cannot load configuration. Exiting")
		return
	}

	log.Info("Initializing jwt (could take a couple of seconds if we have to fetch keys from a JWKS URL)...")
	jwtWrapper, err := jwt.New()
	if err != nil {
		log.WithError(err).Fatal("Cannot instantiate JWT wrapper. Exiting")
		return
	}

	var source sources.Source
	if config.CookieJwtSourceEnabled {
		source = sources.NewCookieSource(config.CookieJwtSourceName)
	}
	if config.HeaderJwtSourceEnabled {
		source = sources.NewHeaderSource(config.HeaderJwtSourceName, config.HeaderJwtSourcePrefix)
	}

	if source == nil || config.CookieJwtSourceEnabled == config.HeaderJwtSourceEnabled {
		log.Fatal("Exactly one between Header source or Cookie source must be enabled")
		return
	}

	server := server.New(source, jwtWrapper)
	server.Start(config.Address)
}
