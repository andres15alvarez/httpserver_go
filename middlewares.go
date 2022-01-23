package main

import (
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headers := r.Header
			_, exist := headers["Authorization"]
			if exist {
				f(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	}
}

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.Method, r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
