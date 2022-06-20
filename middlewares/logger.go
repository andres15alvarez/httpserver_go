package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/andres15alvarez/go_http_server/utils"
)

func Logger() utils.Middleware {
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
