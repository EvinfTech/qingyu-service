package models

import (
	"time"
)

type RechargeRule struct {
	ID           uint32    `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	OriginMoney  uint32    `gorm:"column:origin_money;type:decimal(10,2) unsigned;not null" json:"origin_money"`   // 原金额
	PresentMoney uint32    `gorm:"column:present_money;type:decimal(10,2) unsigned;not null" json:"present_money"` // 赠送金额
	CreatorID    uint32    `gorm:"column:creator_id;type:int unsigned;not null" json:"creator_id"`
	MenderID     uint32    `gorm:"column:mender_id;type:int unsigned;not null" json:"mender_id"`
	CreateAt     time.Time `gorm:"column:create_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt     time.Time `gorm:"column:update_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	IsDel        uint8     `gorm:"column:is_del;type:tinyint unsigned;not null;default:0" json:"is_del"`
}
