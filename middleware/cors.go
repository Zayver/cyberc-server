package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zayver/cyberc-server/config"
)

type CorsMiddleware struct {
	config  config.ConfigHolder
}

func NewCorsMiddleware(config config.ConfigHolder) CorsMiddleware{
	return CorsMiddleware{
		config: config,
	}
}


func (c *CorsMiddleware) Setup() gin.HandlerFunc{
	var corsM gin.HandlerFunc
	if c.config.Env == "development" {
		corsM = cors.New(cors.Config{
			AllowCredentials: true,
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowHeaders:   []string{"*"},
			AllowMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		})
	}else{
		corsM = cors.New(cors.Config{
			AllowCredentials: true,
			AllowOriginFunc:  func(origin string) bool {
				return origin == "cyberc.vercel.app"
			},
			AllowHeaders:   []string{},
			AllowMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		})
	}
	return corsM
}
