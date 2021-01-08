package entity

import (
	"time"
	"github.com/shopspring/decimal"
)

// AccountBook account book
type AccountBook struct {
	// Id account bool record id
	Id int `gorm:"AUTO_INCREMENT"`

	// Amount amount of money
	Amount Decimal `gorm:"not null"`

	// Type the record type
	// TODO

	// BeginTime the begin time when the order take effect
	BeginTime time.Time `gorm:"not null"`

	// EndTime the end time when the order end
	EndTime time.Time `gorm:"not null"`

	// Date the begin date when the order take effect(used for union index)
	Date int `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"not null"`
}
