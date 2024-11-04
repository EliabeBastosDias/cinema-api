package middlewares

import (
	"errors"
	"net/http"
	"strings"

	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtHandler jwthandler.JWTHandler, logger logger.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			err := errors.New("authorization header required")
			logger.Error("Authorization header missing", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			err := errors.New("invalid authorization header format")
			logger.Error("Invalid Authorization header format", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := jwtHandler.ValidateToken(tokenString)
		if err != nil {
			logger.Error("Token validation failed", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}
