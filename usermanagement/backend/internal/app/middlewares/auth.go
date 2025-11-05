// 认证中间件，用于保护需要身份验证的路由，处理横切需求，JWT鉴权，日志记录，cors跨域，panic恢复等。
package middlewares

import (
	"backend/internal/pkg/e"
	"backend/internal/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			util.RespondError(c, http.StatusUnauthorized, e.ErrorAuthCheckTokenFail)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.RespondError(c, http.StatusUnauthorized, e.ErrorAuth)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			util.RespondError(c, http.StatusUnauthorized, e.ErrorAuthTokenTimeout)
			c.Abort()
			return
		}
		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
