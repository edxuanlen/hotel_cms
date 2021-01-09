package entity

import (
	"time"
)

// RoomOrder room order
type RoomOrder struct {
	// Id 用户id
	Id int `gorm:"AUTO_INCREMENT"`

	// Phone the phone who use to order
	Phone string `gorm:"not null, unique"`

	// Name the name of the user who order
	Name string `gorm:"not null"`

	// BeginTime the begin time when the order take effect
	BeginTime time.Time `gorm:"not null"`

	// EndTime the end time when the order end
	EndTime time.Time `gorm:"not null"`

	// Date the begin date when the order take effect(used for union index)
	Date int `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// ModifyTime 更新时间
	ModifyTime time.Time `gorm:"not null"`
}
