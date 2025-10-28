// main.go
package main

import (
	"fmt"
	"practice_usermanagement/config" // 确保这是你的模块名
	"practice_usermanagement/pkg/database"
	"practice_usermanagement/routes"
	"time"

	"github.com/gin-contrib/cors" // <--- 1. 导入
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDB()

	r := gin.Default()

	// CORS 配置
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	routes.RegisterRoutes(r)

	port := fmt.Sprintf(":%s", config.Config.ServerPort)
	if err := r.Run(port); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
