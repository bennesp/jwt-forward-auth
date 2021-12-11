package sources

import "github.com/gin-gonic/gin"

type CookieSource struct {
	Name string
}

func NewCookieSource(name string) *CookieSource {
	return &CookieSource{
		Name: name,
	}
}

func (s *CookieSource) RetrieveJwt(c *gin.Context) (string, error) {
	cookie, err := c.Request.Cookie(s.Name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
