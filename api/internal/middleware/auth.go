package middleware

import (
	"context"
	"net/http"
	"strings"

	"acupofcoffee/common/errorx"
	"acupofcoffee/common/response"

	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	AccessSecret string
}

func NewAuthMiddleware(accessSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		AccessSecret: accessSecret,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Error(w, errorx.NewCodeError(http.StatusUnauthorized, "missing authorization header"))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(w, errorx.NewCodeError(http.StatusUnauthorized, "invalid authorization format"))
			return
		}

		tokenString := parts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.AccessSecret), nil
		})

		if err != nil || !token.Valid {
			response.Error(w, errorx.NewCodeError(http.StatusUnauthorized, "invalid or expired token"))
			return
		}

		// 将用户信息存入 context
		userID, ok := claims["userId"].(float64)
		if !ok {
			response.Error(w, errorx.NewCodeError(http.StatusUnauthorized, "invalid token claims"))
			return
		}

		ctx := context.WithValue(r.Context(), "userId", uint(userID))
		next(w, r.WithContext(ctx))
	}
}
