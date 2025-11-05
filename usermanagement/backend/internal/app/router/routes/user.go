// 路由注册，URL和处理函数的映射表，`r.GET(...)`, `r.POST(...)`, `r.Group(...)` 以及应用中间件等
package routes

import (
	controller "backend/internal/app/handler"
	"backend/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)

		userRoutes := api.Group("/user")
		userRoutes.Use(middleware.AuthMiddleware())
		{
			userRoutes.GET("/search", controller.SearchUsers)
			userRoutes.GET("/:uid", controller.GetUser)
			userRoutes.PUT("", controller.UpdateUser) //更新用户信息,使用空的路径，对应/api/user
			userRoutes.GET("/me", controller.GetMe)
		}
	}
}
