package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type contextKey string

const UserContextKey contextKey = "user"

type AuthUser struct {
	Username string
	Role     string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access-token")
		if err != nil {
			http.Error(w, "JWT cookie missing", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		// Load .env and secret key
		_ = godotenv.Load()
		secretKey := []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			log.Println("Unauthorized")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		username, uok := claims["username"].(string)
		role, rok := claims["role"].(string)
		if !uok || !rok {
			http.Error(w, "Missing token fields", http.StatusUnauthorized)
			return
		}

		// Set user in context
		user := AuthUser{
			Username: username,
			Role:     role,
		}
		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
