package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		allowed := false

		switch origin {
		case "http://localhost:5173":
			allowed = true

		case "https://nightveilrecords.pages.dev":
			allowed = true
		}

		if allowed {
			w.Header().Set(
				"Access-Control-Allow-Origin",
				origin,
			)

			w.Header().Set(
				"Access-Control-Allow-Methods",
				"GET, POST, PUT, DELETE, OPTIONS",
			)

			w.Header().Set(
				"Access-Control-Allow-Headers",
				"Content-Type, Authorization",
			)

			w.Header().Set(
				"Access-Control-Allow-Credentials",
				"true",
			)
		}

		// preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
