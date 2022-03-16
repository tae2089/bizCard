package util

import (
	"bizCard/domain"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var hmacSampleSecret []byte

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
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Id, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

	return 1
}
