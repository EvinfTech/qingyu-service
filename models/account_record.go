package models

import (
	"qiyu/models/enum"
	"time"
)

type AccountRecord struct {
	ID             uint32                 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Ouid           string                 `gorm:"column:ouid;type:varchar(64);not null" json:"ouid"`
	ChangeType     enum.AccountChangeType `gorm:"column:change_type;type:tinyint unsigned;not null" json:"change_type"`               // 变更类型
	ActualMoney    uint32                 `gorm:"column:actual_money;type:decimal(10,2);not null" json:"actual_money"`                // 实际金额
	PresentMoney   uint32                 `gorm:"column:present_money;type:decimal(10,2);not null" json:"present_money"`              // 赠送金额
	NewActual      uint32                 `gorm:"column:new_actual;type:decimal(10,2) unsigned;not null" json:"new_actual"`           // 新实际金额
	NewPresent     uint32                 `gorm:"column:new_present;type:decimal(10,2) unsigned;not null" json:"new_present"`         // 新赠送金额
	RelatedOrderNo string                 `gorm:"column:related_order_no;type:varchar(31) unsigned;not null" json:"related_order_no"` // 关联订单号
	CreateAt       time.Time              `gorm:"column:create_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	CreatorID      uint32                 `gorm:"column:creator_id;type:int unsigned;not null" json:"creator_id"`
}
