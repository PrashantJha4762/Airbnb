package middleware

import (
	dbconfig "AuthInGoService/config/db"
	envconfig "AuthInGoService/config/env"
	repo "AuthInGoService/db/repositories"
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
			return []byte(envconfig.GetString("JWT_SECRET_KEY", "token")), nil
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

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbconn, err := dbconfig.SetUpDB()
			if err != nil {
				http.Error(w, "Authorization service is unavailable", http.StatusInternalServerError)
				return
			}
			defer dbconn.Close()
			RequireAllRolesWithRepository(repo.NewUserRoleRepository(dbconn), roles...)(next).ServeHTTP(w, r)
		})
	}
}

// RequireAllRolesWithRepository authorizes a request with an application-owned
// repository. Prefer this form when wiring routes: it reuses the application's
// database pool instead of opening a connection for every request.
func RequireAllRolesWithRepository(userRoleRepository repo.UserRoleRepository, roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := UserIDFromContext(r.Context())
			if !ok || userID < 1 {
				http.Error(w, "Authentication is required", http.StatusUnauthorized)
				return
			}

			if len(roles) == 0 {
				http.Error(w, "No roles configured for this route", http.StatusInternalServerError)
				return
			}

			for _, role := range roles {
				roleName := strings.TrimSpace(role)
				if roleName == "" {
					http.Error(w, "Invalid role configured for this route", http.StatusInternalServerError)
					return
				}

				hasRole, err := userRoleRepository.HasRole(userID, roleName)
				if err != nil {
					http.Error(w, "Failed to fetch user roles", http.StatusInternalServerError)
					return
				}
				if !hasRole {
					http.Error(w, "Forbidden: required role is missing", http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAllPermissions only allows a request through when the authenticated
// user has every requested permission. It remains for callers outside the app
// wiring; use RequireAllPermissionsWithRepository in routers.
func RequireAllPermissions(permissions ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbconn, err := dbconfig.SetUpDB()
			if err != nil {
				http.Error(w, "Authorization service is unavailable", http.StatusInternalServerError)
				return
			}
			defer dbconn.Close()
			RequireAllPermissionsWithRepository(repo.NewUserRoleRepository(dbconn), permissions...)(next).ServeHTTP(w, r)
		})
	}
}

func RequireAllPermissionsWithRepository(userRoleRepository repo.UserRoleRepository, permissions ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := UserIDFromContext(r.Context())
			if !ok || userID < 1 {
				http.Error(w, "Authentication is required", http.StatusUnauthorized)
				return
			}
			if len(permissions) == 0 {
				http.Error(w, "No permissions configured for this route", http.StatusInternalServerError)
				return
			}
			for _, permission := range permissions {
				permissionName := strings.TrimSpace(permission)
				if permissionName == "" {
					http.Error(w, "Invalid permission configured for this route", http.StatusInternalServerError)
					return
				}
				hasPermission, err := userRoleRepository.HasPermission(userID, permissionName)
				if err != nil {
					http.Error(w, "Failed to fetch user permissions", http.StatusInternalServerError)
					return
				}
				if !hasPermission {
					http.Error(w, "Forbidden: required permission is missing", http.StatusForbidden)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
