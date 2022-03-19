package util

import (
	"bizCard/domain"
	"github.com/golang-jwt/jwt"
	"time"
)

var hmacSampleSecret = []byte("secret")

type MyCustomClaims struct {
	Id int
	domain.UserInfo
	jwt.StandardClaims
}

func CreateJwt(userInfo domain.UserInfo, userId int) (string, error) {

	claims := &MyCustomClaims{
		userId, userInfo, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSampleSecret)
	return tokenString, err
}

func ParseJwt(tokenString string) int {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		return 0
	}
	claims, ok := token.Claims.(*MyCustomClaims)

	if !(ok && token.Valid) {
		return -1
	}

	return claims.Id
}
