package utils

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}
