package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Home struct {
	Data string `json:"data"`
}
