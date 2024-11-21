package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	jwt.StandardClaims

	UserID int    `json:"user_id"`
	Role   string `json:"role,omitempty"`
}
