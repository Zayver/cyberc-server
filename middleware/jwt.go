package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zayver/cyberc-server/service"
)

type JwtMiddleware struct{
	jwtService service.JwtService
}

func NewJwtMiddleware(jwt service.JwtService) JwtMiddleware{
	return JwtMiddleware{
		jwtService: jwt,
	}
}

func (j *JwtMiddleware) Handler() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 && t[0] == "Bearer"{
			token := t[1]
			authorized, _ := j.jwtService.CheckToken(token)
			if !authorized {
				ctx.Status(http.StatusUnauthorized)
				ctx.Abort()
				return
			}
			ctx.Next()
			return 
		}
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	}
}