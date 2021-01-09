package entity

import (
	"time"
)

// CommodityInventory 仓库库存
type CommodityInventory struct {

	// Id account bool record id
	Id int `gorm:"AUTO_INCREMENT"`

	// CommodityName 物品名
	CommodityName string `gorm:"not null"`

	// Stock 库存
	Stock int `gorm:"not null"`

	// Leader 负责人
	Leader string `gorm:"not null"`

	// LeaderId 负责人id
	LeaderId int `gorm:"not null"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"not null"`

	// ModifyTime 更新时间
	ModifyTime time.Time `gorm:"not null"`
}
