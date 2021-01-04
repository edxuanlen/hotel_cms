package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type User struct {
	Id           int    `gorm:"AUTO_INCREMENT"`
	IdentityCard string `gorm:"size:20"`
	Name         string `gorm:"size:20"`
	Phone        string `gorm:"size:20;not null; unique"`
	Password     string `gorm:"not null"`
	Salary       int
	Status       byte      `gorm:"not null"`
	EntryTime    time.Time `gorm:"not null"`
	LastPayday   time.Time `gorm:"not null"`
	CreateTime   time.Time `gorm:"not null"`
	UpdateTime   time.Time `gorm:"not null"`
	Access       byte      `gorm:"not null"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(Id interface{}) (User, error) {
	var user User
	result := DB.First(&user, Id)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
