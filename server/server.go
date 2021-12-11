package server

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bennesp/traefik-jwt-forward-auth/services/jwt"
	"github.com/bennesp/traefik-jwt-forward-auth/sources"
	"github.com/gin-gonic/gin"
)

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
