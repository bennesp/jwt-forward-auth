package sources

import "github.com/gin-gonic/gin"

type CookieSource struct {
	cookieName string
}

func NewCookieSource(cookieName string) *CookieSource {
	return &CookieSource{
		cookieName: cookieName,
	}
}

func (s *CookieSource) RetrieveJwt(c *gin.Context) (string, error) {
	cookie, err := c.Request.Cookie(s.cookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
