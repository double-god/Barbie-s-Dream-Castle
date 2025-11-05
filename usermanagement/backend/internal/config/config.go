/*
读取 `.env`、`config.yaml` 等文件或环境变量。
只做：把配置信息加载到一个全局的 `struct` 中。
*/
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 全局配置结构体
type AppConfig struct {
	ServerPort string
	DBHost     string // <-- 修复：DBHOST -> DBHost
	DBPort     string // <-- 修复：DBPORT -> DBPort
	DBUser     string // <-- 修复：DBUSER -> DBUser
	DBPassword string // <-- 这个本来就是对的
	DBName     string // <-- 修复：DBNAME -> DBName
	JWTSecret  string
}

var Config AppConfig // 全局配置变量

// LoadConfig 加载配置
func LoadConfig() {
	// 尝试加载 .env 文件（如果存在）
	err := godotenv.Load() // 从当前目录加载 .env 文件,有助于配置管理环境变量
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	fmt.Println("--- 开始调试配置 ---")
	fmt.Println("读取到的 DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("读取到的 DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("读取到的 DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("--- 结束调试配置 ---")

	// 从环境变量读取配置
	Config = AppConfig{
		ServerPort: os.Getenv("HTTP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"), // <-- 直接用 os.Getenv
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
