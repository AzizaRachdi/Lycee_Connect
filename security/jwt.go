// internal/security/jwt.go
package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("change-this-secret")

type Claims struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string) (string, error) {
	claims := &Claims{
		Sub:  userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
