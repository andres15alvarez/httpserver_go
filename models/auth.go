package models

import "github.com/golang-jwt/jwt"

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Claims struct {
	Id    int64
	Email string
	jwt.StandardClaims
}
