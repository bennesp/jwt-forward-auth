package sources

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderSource struct {
	headerName        string
	headerValuePrefix string
}

func NewHeaderSource(headerName, headerValuePrefix string) *HeaderSource {
	return &HeaderSource{
		headerName:        headerName,
		headerValuePrefix: headerValuePrefix,
	}
}

func (s *HeaderSource) RetrieveJwt(c *gin.Context) (string, error) {
	header := c.Request.Header.Get(s.headerName)

	if header == "" {
		err := fmt.Errorf("no '%s' header", s.headerName)
		return "", err
	}

	return strings.TrimPrefix(header, s.headerValuePrefix), nil
}
