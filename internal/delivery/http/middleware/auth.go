package middleware

import (
	"context"
	"net/http"
	"strings"

	jwtservice "github.com/MaksimCpp/TaskManager/internal/service/jwt"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func Auth(jwtService *jwtservice.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token.", http.StatusUnauthorized)
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Invalid token format.", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := jwtService.ParseToken(token)
			if err != nil {
				http.Error(w, "Invalid token.", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
