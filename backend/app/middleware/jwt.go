package middleware

import (
	"angrymiao-ai/app/auth"
	"angrymiao-ai/app/exception"
	"angrymiao-ai/app/mapping"
	"angrymiao-ai/config"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	userIdKey = "user_id"
)

func setDefaultUserKey(c *gin.Context, claims *auth.JWTClaims) {
	var userID int
	var userRole = mapping.UserRoleAnonymous

	if claims != nil {
		userID = claims.UserID
		userRole = claims.Role
	}

	c.Set(mapping.UserIDKey, userID)
	c.Set(mapping.UserRoleKey, userRole)
}

func BaseJWTAuthMiddleware(c *gin.Context, isForce bool, ruleFunc func(*gin.Context, *auth.JWTClaims) bool) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		if isForce {
			exception.UnauthorizedException(c, "")
			return
		} else {
			c.Set(userIdKey, 0)
			return
		}
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 {
		exception.UnauthorizedException(c, exception.InvalidAuthMsg)
		return
	}

	if parts[0] != config.Conf.JWT.Key {
		exception.UnauthorizedException(c, exception.InvalidAuthMsg)
		return
	}

	if parts[1] != "iFVXoa27z4GnyWKmhcuOLTe4qzj8k96bSSb5QIlL" {
		exception.UnauthorizedException(c, exception.InvalidAuthMsg)
		return
	}

	//claims, err := auth.ParseToken(parts[1])
	//if err != nil {
	//	exception.UnauthorizedException(c, exception.InvalidAuthMsg)
	//	return
	//}
	//
	//// 设置用户信息
	//setDefaultUserKey(c, claims)
	//
	//// 自定义规则函数
	//if ruleFunc != nil && !ruleFunc(c, claims) {
	//	return
	//}

	c.Next()
}

func adminRule(c *gin.Context, claims *auth.JWTClaims) (allow bool) {
	value, exists := c.Get(mapping.UserRoleKey)
	if !exists || value != mapping.UserRoleAdmin {
		exception.UnauthorizedException(c, exception.InvalidAuthMsg)
		return false
	}
	return true
}

func JWTAuthMiddlewareForAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		BaseJWTAuthMiddleware(c, true, adminRule)
	}
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		BaseJWTAuthMiddleware(c, true, nil)
	}
}

func JWTAuthNotForceMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		BaseJWTAuthMiddleware(c, false, nil)
	}
}
