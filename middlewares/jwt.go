package middlewares

import (
	"asis_quest/config"
	"asis_quest/helper"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					return config.JWT_KEY, nil
				})

				if err != nil {
					response := map[string]string{"message": "Unauthorized"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}

				if !token.Valid {
					response := map[string]string{"message": "Unauthorized"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}
				next.ServeHTTP(w, r)
			}
		} else {
			response := map[string]string{"message": "An Authorization header is required"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
	})
}
