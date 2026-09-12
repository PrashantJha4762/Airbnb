package middleware

import (
	config "AuthInGoService/config/env"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDContextKey contextKey = "userId"

func UserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userIDContextKey).(int)
	return userID, ok
}

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization Handler is required", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Bearer must be present", http.StatusUnauthorized)
			return
		}
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}
		claims := jwt.MapClaims{}

		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return ([]byte(config.GetString("JWT_SECRET_KEY", "token"))), nil
		})
		if err != nil || !parsedToken.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		UserId, ok := claims["user_id"].(float64)
		email, okemail := claims["email"].(string)
		if !ok || !okemail {
			http.Error(w, "Invalid Token: user_id or email not found", http.StatusUnauthorized)
			return
		}
		userID := int(UserId)
		if float64(userID) != UserId || userID < 1 {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println("Authenticated User ID:", userID, " Email:", email)
		ctx := context.WithValue(r.Context(), userIDContextKey, userID)
		ctx = context.WithValue(ctx, "userId", fmt.Sprintf("%d", userID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
