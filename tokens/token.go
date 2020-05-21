package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// type Token struct {
// }

// func NewToken() *Token {
// 	return &Token{}
// }

func CreateToken(userId uint64, secret string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
