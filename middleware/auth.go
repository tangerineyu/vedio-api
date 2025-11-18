package middleware

import (
	"errors"
	"net/http"
	"strings"
	"video-api/handler"
	"video-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 验证access token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			authHeader = c.Query("token")
			if authHeader == "" {
				handler.Error(c, http.StatusUnauthorized, "No Authorization header", "未提供token")
				c.Abort()
				return
			}

		}
		parts := strings.Split(authHeader, " ")
		tokenString := ""
		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenString = parts[1]
		} else {
			tokenString = authHeader
		}
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				handler.Error(c, http.StatusUnauthorized, "AUTH_TOKEN_EXPIRED", "token已过期")

			} else {
				handler.Error(c, http.StatusUnauthorized, "AUTH_INVALID", "token验证失败")
			}
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()

	}
}
