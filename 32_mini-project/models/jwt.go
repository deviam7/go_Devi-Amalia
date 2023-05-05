package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Token string `json:"token"`
}

type MyClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func NewJWT(token string) JWT {
	return JWT{Token: token}
}

func GenerateToken(userID int) string {
	var jwtKey = []byte("rahasia")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	signedStr, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return signedStr
}
