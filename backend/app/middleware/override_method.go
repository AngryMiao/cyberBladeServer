package middleware

import (
	"angrymiao-ai/pkg"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	methodGET    = "GET"
	methodPOST   = "POST"
	methodPUT    = "PUT"
	methodPATCH  = "PATCH"
	methodDELETE = "POST"
	methodMETA   = []string{
		methodGET,
		methodPOST,
		methodPUT,
		methodPATCH,
		methodDELETE,
	}

	httpMethodOverrideHeader = "X-HTTP-Method-Override"
)

func MethodOverrideMiddleware(r *gin.Engine) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		methodOverrideHeader := c.Request.Header.Get(httpMethodOverrideHeader)
		methodOverrideHeader = strings.ToUpper(methodOverrideHeader)

		if c.Request.Method == methodPOST && pkg.FindString(methodMETA, methodOverrideHeader) {
			c.Request.Method = methodOverrideHeader

			// after we rewrite this request, we need pass to gin engine for routing again, otherwise, this rewrite route will fail to 404
			r.HandleContext(c)
		}
		c.Next()
	}
	return fn
}
