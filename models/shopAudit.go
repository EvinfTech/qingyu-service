package models

import "time"

type ShopAudit struct {
	Id              int       `json:"id"`               // 自增id
	Name            string    `json:"name"`             // 名字
	Address         string    `json:"address"`          // 门店地址
	Longitude       float64   `json:"longitude"`        // 经度
	Latitude        float64   `json:"latitude"`         // 维度
	SiteCount       int       `json:"site_count"`       // 场地数量
	GatePhoto       string    `json:"gate_photo"`       // 门口照片
	IndoorPhoto     string    `json:"indoor_photo"`     // 室内照片
	BusinessLicense string    `json:"business_license"` // 营业执照
	IdCardFront     string    `json:"id_card_front"`    // 身份证正面
	IdCardBack      string    `json:"id_card_back"`     // 身份证反面
	UserOuid        string    `json:"user_ouid"`        // 用户ouid
	UserName        string    `json:"user_name"`        // 联系人名称
	UserPhone       string    `json:"user_phone"`       // 联系人电话
	ShopId          int       `json:"shop_id"`
	Desc            string    `json:"desc"`       // 备注
	Status          string    `json:"status"`     // 审核状态 Y:审核成功  N:审核中
	GmtCreate       time.Time `json:"gmt_create"` // 创建时间
	GmtUpdate       time.Time `json:"gmt_update"` // 更新时间
}
