package sources

import "github.com/gin-gonic/gin"

type Source interface {
	RetrieveJwt(*gin.Context) (string, error)
}

func NewAuthorizationHeaderSource() Source {
	return &HeaderSource{
		headerName:        "Authorization",
		headerValuePrefix: "Bearer ",
	}
}

func NewTokenCookieSource() Source {
	return &CookieSource{
		cookieName: "token",
	}
}
