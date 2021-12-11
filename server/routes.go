package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
