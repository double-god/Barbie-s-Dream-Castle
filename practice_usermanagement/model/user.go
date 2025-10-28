// 用户模型
// 定义数据长什么样，定义struct,gorm标签和json标签
package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	Username     string `gorm:"unique;not null" json:"username"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"` // JSON响应中忽略此字段
	Bio          string `json:"bio"`
	CreatedAt    int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    int64  `gorm:"autoUpdateTime" json:"updated_at"`
}

// SetPassword 设置用户密码，进行哈希处理
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

// CheckPassword 验证用户密码
func (u *User) CheckPassword(password string) bool { //bcrypt.CompareHashAndPassword 比较哈希值和明文密码.这是一种方法的定义
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
