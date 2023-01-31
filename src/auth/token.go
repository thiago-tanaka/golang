package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["userId"] = userId
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("secret"))
}
