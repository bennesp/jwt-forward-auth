package server

import (
	"fmt"
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

	claims := ctx.jwtWrapper.GetClaims(token)
	for key, value := range ctx.claimMapping {
		c.Header(value, fmt.Sprintf("%v", claims[key]))
	}
	c.Status(http.StatusOK)
}
