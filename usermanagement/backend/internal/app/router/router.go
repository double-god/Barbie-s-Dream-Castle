package router

import (
	"backend/internal/app/router/routes"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 统一对外暴露的路由注册入口
func RegisterRoutes(r *gin.Engine) {
	routes.RegisterRoutes(r)
}
