package middlewares

import (
	"net/http"

	"github.com/andres15alvarez/go_http_server/utils"
)

func CheckAuth() utils.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headers := r.Header
			token, exist := headers["Authorization"]
			if exist {
				_, err := utils.ParseToken(token[0])
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				f(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	}
}
