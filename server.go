package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bennesp/traefik-jwt-forward-auth/services/jwt"
	"github.com/bennesp/traefik-jwt-forward-auth/sources"
	"github.com/gin-gonic/gin"
)

func (ctx *Context) handleGet(c *gin.Context) {
	jwtString, err := ctx.source.RetrieveJwt(c)

	if err != nil {
		log.WithError(err).Error("Error retrieving JWT from source")
		c.Status(http.StatusUnauthorized)
		return
	}

	token, err := ctx.jwtWrapper.Verify(jwtString)

	if err != nil {
		log.WithError(err).WithField("jwt", jwtString).Error("Error verifying JWT")
		c.Status(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		log.WithField("jwt", jwtString).Error("JWT is not valid")
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusOK)
}

type Context struct {
	source     sources.Source
	jwtWrapper *jwt.JwtWrapper
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.WithError(err).Fatal("Cannot load configuration. Exiting")
		return
	}

	l, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.WithError(err).Error("Cannot parse log level")
	} else {
		log.SetLevel(l)
	}

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

	ctx := &Context{
		source:     source,
		jwtWrapper: jwtWrapper,
	}

	log.Infof("Starting server on port %d...", config.Port)
	r := gin.Default()
	r.GET("/", ctx.handleGet)
	http.ListenAndServe(":"+fmt.Sprint(config.Port), r)
}
