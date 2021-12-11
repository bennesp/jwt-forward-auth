package sources

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderSource struct {
	Name   string
	Prefix string
}

func NewHeaderSource(Name, Prefix string) *HeaderSource {
	return &HeaderSource{
		Name:   Name,
		Prefix: Prefix,
	}
}

func (s *HeaderSource) RetrieveJwt(c *gin.Context) (string, error) {
	header := c.Request.Header.Get(s.Name)

	if header == "" {
		err := fmt.Errorf("no '%s' header", s.Name)
		return "", err
	}

	return strings.TrimPrefix(header, s.Prefix), nil
}
