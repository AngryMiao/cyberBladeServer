package middleware

import (
	"angrymiao-ai/config"
	"angrymiao-ai/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	maxAge = 12
)

var (
	allowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	allowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
)

func CORSMiddleware() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowMethods:     allowMethods,
		AllowHeaders:     allowHeaders,
		AllowCredentials: false,
		MaxAge:           maxAge * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			origin = strings.ToLower(origin)
			return pkg.IsAllowHost(origin, config.Conf.AllowHosts)
		},
	}
	return cors.New(corsConfig)
}
