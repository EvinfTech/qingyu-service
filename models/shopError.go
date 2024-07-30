package models

import "time"

type ShopError struct {
	Id        int       `json:"id"`      //自增id
	ShopId    int       `json:"shop_id"` //门店id
	UserOuid  string    `json:"user_ouid"`
	Type      string    `json:"type"`       //错误类型
	Detail    string    `json:"detail"`     //详情
	Photos    string    `json:"photos"`     //照片
	GmtCreate time.Time `json:"gmt_create"` //创建时间
	GmtUpdate time.Time `json:"gmt_update"` //更新时间
}
