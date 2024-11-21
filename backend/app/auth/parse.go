package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JWTClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return verifyKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		if isExpire(claims.ExpiresAt) {
			return nil, errors.New("token expired")
		}

		if claims.Issuer != issuer {
			return nil, errors.New("invalid token issuer")
		}

		return claims, nil
	}

	return nil, err
}

func isExpire(expire int64) bool {
	if expire-time.Now().Unix() < 0 {
		return true
	}
	return false
}
