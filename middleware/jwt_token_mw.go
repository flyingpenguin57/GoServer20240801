package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建 JWT 中间件
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("token")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token header"})
            c.Abort()
            return
        }

		if tokenString != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
		}

        c.Next()
    }
}