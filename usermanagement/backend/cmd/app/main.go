// main.go
package main

import (
	"backend/internal/app/router"
	"backend/internal/config" // 确保这是你的模块名
	"backend/internal/pkg/database"
	"fmt"
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

	router.RegisterRoutes(r)

	port := fmt.Sprintf(":%s", config.Config.ServerPort)
	if err := r.Run(port); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
