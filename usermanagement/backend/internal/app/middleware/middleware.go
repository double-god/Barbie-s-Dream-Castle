// middleware 包提供了一些中间件的封装
package middleware

import (
	"backend/internal/app/middlewares"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 代理到实际 middlewares 包，便于过渡
func AuthMiddleware() gin.HandlerFunc {
	return middlewares.AuthMiddleware()
}
