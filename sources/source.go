package sources

import "github.com/gin-gonic/gin"

type Source interface {
	RetrieveJwt(*gin.Context) (string, error)
}
