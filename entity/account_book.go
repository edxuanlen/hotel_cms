package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

// AccountBook account book
type AccountBook struct {
	// Id account bool record id
	Id int `gorm:"AUTO_INCREMENT"`

	// Amount amount of money
	Amount decimal.Decimal `gorm:"not null"`

	// Type the record type
	// 0  房间收入
	// 1  水电开销
	// 2  日用品开销
	// 3  人员开销
	// 4  其他收入
	// 5  其他支出
	Type byte `gorm:"not null"`

	// Leader 负责人
	Leader string `gorm:"not null"`

	// LeaderId 负责人id
	LeaderId int `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// ModifyTime 更新时间
	ModifyTime time.Time `gorm:"not null"`
}
