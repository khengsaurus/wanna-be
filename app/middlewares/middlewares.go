package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/khengsaurus/wanna-be/consts"
)

// func EnableCors(h http.Handler) http.Handler {
// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedHeaders:   []string{"*"},
// 		AllowCredentials: true,
// 	})

// 	return c.Handler(h)
// }

func WithContext(key consts.ContextKey, client interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return WithContextFn(key, client, next)
	}
}

func WithContextFn(key consts.ContextKey, client interface{}, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), key, client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
