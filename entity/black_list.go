package entity

import (
	"time"
)

// BlackList 黑名单
type BlackList struct {

	// Id black list id
	Id int `gorm:"AUTO_INCREMENT"`

	// Name 姓名
	Name string `gorm:"not null"`

	// Phone 手机号
	Phone string `gorm:"not null"`

	// Cause cause by
	Cause string `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// ModifyTime 更新时间
	ModifyTime time.Time `gorm:"not null"`
}
