package entity

import (
	"time"
)

// RoomCard room card
type RoomCard struct {
	// Id 用户id
	Id int `gorm:"AUTO_INCREMENT"`

	// RoomId the room which is binding in
	RoomId int

	// Status 房间数量
	Status int `gorm:"not null, default: 0"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"not null"`
}
