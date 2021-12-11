package server

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bennesp/traefik-jwt-forward-auth/services/jwt"
	"github.com/bennesp/traefik-jwt-forward-auth/sources"
	"github.com/gin-gonic/gin"
)

func (ctx *Server) handleGet(c *gin.Context) {
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

type Server struct {
	source     sources.Source
	jwtWrapper *jwt.JwtWrapper
}

func New(source sources.Source, jwtWrapper *jwt.JwtWrapper) *Server {
	return &Server{
		source:     source,
		jwtWrapper: jwtWrapper,
	}
}

func (ctx *Server) Start(port int) {
	log.Infof("Starting server on port %d...", port)

	r := gin.Default()
	r.GET("/", ctx.handleGet)

	http.ListenAndServe(":"+fmt.Sprint(port), r)
}
