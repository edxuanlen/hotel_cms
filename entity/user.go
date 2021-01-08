package entity

import (
	"golang.org/x/crypto/bcrypt"
	"hotel_cms/common"
	"time"
)

// User 用户模型
type User struct {
	// Id 用户id
	Id           int    `gorm:"AUTO_INCREMENT"`
	// IdentityCard 身份证
	IdentityCard string `gorm:"size:20"`
	// Name 姓名
	Name         string `gorm:"size:20"`
	// Phone 手机号
	Phone        string `gorm:"size:20;not null; unique"`
	// Password 密码
	Password     string `gorm:"not null"`
	// Salary 月薪
	Salary       int
	// Status 在职状态
	Status       byte      `gorm:"not null"`
	// Access 权限
	Access       byte      `gorm:"not null"`
	// EntryTime 入职时间
	EntryTime    time.Time `gorm:"not null"`
	// LastPayday 最后发薪日
	LastPayday   time.Time `gorm:"not null"`
	// CreateTime 创建时间
	CreateTime   time.Time `gorm:"not null"`
	// UpdateTime 更新时间
	UpdateTime   time.Time `gorm:"not null"`
}


// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), common.PassWordCost)
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
