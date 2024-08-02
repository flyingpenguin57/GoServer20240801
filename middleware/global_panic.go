package middleware

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Recovery 中间件
func RecoveryMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // 记录错误信息
                log.Printf("panic recovered: %s\n", err)

                // 返回标准的错误响应
                c.JSON(http.StatusOK, gin.H{
                    "error": fmt.Sprintf("Internal Server Error: %s", err),
                })
                c.Abort()
            }
        }()
        c.Next()
    }
}