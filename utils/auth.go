package utils

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/andres15alvarez/go_http_server/models"
	"github.com/golang-jwt/jwt"
)

var secretKey string = os.Getenv("SECRET_KEY")
var jwtKey = []byte(secretKey)

func GenerateToken(user models.User) (string, error) {
	claims := &models.Claims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(authorization string) (models.UserBase, error) {
	if !strings.Contains(authorization, "Bearer ") {
		log.Println("The token is invalid")
		return models.UserBase{}, errors.New("Invalid token")
	}
	tokenStr := strings.Split(authorization, " ")
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println("Error to parse token", err)
		return models.UserBase{}, err
	}
	if !tkn.Valid {
		log.Println("Invalid token", err)
		return models.UserBase{}, err
	}

	return models.UserBase{Id: claims.Id, Email: claims.Email}, err
}
