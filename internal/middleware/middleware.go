package middleware

import (
	"errors"
	"net/http"
	"social-network/internal/config"
	"social-network/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := config.Get().SECRET_JWT

	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthenticated"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userID", userID)
		ctx.Set("username", username)
		ctx.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := config.Get().SECRET_JWT

	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthenticated"))
			return
		}

		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("userID", userID)
		ctx.Set("username", username)
		ctx.Next()
	}
}
