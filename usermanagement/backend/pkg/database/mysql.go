// pkg/database/mysql.go
package database

import (
	"fmt"
	"log"
	"practice_usermanagement/config" // 确保这是你的模块名 (practice_usermanagement?)
	"practice_usermanagement/model"
	"time" // <-- 1. 导入 time

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger" // <--- 2. 导入 GORM 的 logger
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBName,
	)

	// 【【【【【 3. 这是关键修改：开启 SQL 日志 】】】】】
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 【【【【【 设置日志级别为 Info 】】】】】
			IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound 错误
			Colorful:                  true,        // 彩色打印
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // <-- 【【【【【 使用我们新配置的 Logger 】】】】】
	})
	// 【【【【【 修改结束 】】】】】

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database connection successful.")
}
