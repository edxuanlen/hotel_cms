package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type partner struct {

	// Id account bool record id
	Id int `gorm:"AUTO_INCREMENT"`


	// PartnerName 合作伙伴名字
	PartnerName string `gorm:"not null"`

	// Discount 折扣力度
	Discount decimal.Decimal `gorm:"not null"`

	// Status 当前合作状态
	// 0 不合作
	// 1 合作
	Status byte `gorm:"not null"`

	// Leader 负责人
	Leader string `gorm:"not null"`

	// LeaderId 负责人id
	LeaderId int `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// ModifyTime 更新时间
	ModifyTime time.Time `gorm:"not null"`
}
