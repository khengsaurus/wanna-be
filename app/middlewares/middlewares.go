package middlewares

import (
	"net/http"
	"strings"
)

// func EnableCors(h http.Handler) http.Handler {
// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedHeaders:   []string{"*"},
// 		AllowCredentials: true,
// 	})

// 	return c.Handler(h)
// }

func SetHeader(header string, value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(header, value)
			next.ServeHTTP(w, r)
		})
	}
}

func VerifyHeader(header string, value string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			to_verify := r.Header.Get(header)
			if to_verify == "" || !strings.HasPrefix(to_verify, value) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
