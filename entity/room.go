package entity

import (
	"time"
)

// Room 用户模型
type Room struct {
	// Id 用户id
	Id int `gorm:"AUTO_INCREMENT"`

	// Type 房间类型
	//0 单人房
	//1 标间
	//2 大床
	//3 亲子房
	//4 套房
	Type byte `gorm:"not null"`

	// Price 房间单价
	Price int `gorm:"not null"`

	// Number 房间数量
	Number int `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"not null"`
}
