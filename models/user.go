package models

import (
	"qiyu/models/enum"
	"time"
)

type User struct {
	Id             int          `json:"id"`                                                                                 //id
	Ouid           string       `json:"ouid"`                                                                               //用户ouid
	Name           string       `json:"name"`                                                                               //用户名
	Avatar         string       `json:"avatar"`                                                                             //用户头像
	Phone          string       `json:"phone"`                                                                              //用户手机号
	Birthday       int          `json:"birthday"`                                                                           //用户生日
	Status         string       `json:"status"`                                                                             //用户状态
	Sex            int          `json:"sex"`                                                                                //用户性别
	TotalCount     int          `json:"total_count"`                                                                        //统计运动次数
	TotalLength    int          `json:"total_length"`                                                                       //统计运动时长
	WxOpenid       string       `json:"wx_openid"`                                                                          //微信openid
	WxUnionid      string       `json:"wx_unionid"`                                                                         //微信unionid
	Introduce      string       `json:"introduce"`                                                                          //个人简介
	GmtCreate      time.Time    `json:"gmt_create"`                                                                         //创建时将
	GmtUpdate      time.Time    `json:"gmt_update"`                                                                         //更新时间
	VipType        enum.VipType `gorm:"column:vip_type;type:tinyint unsigned;not null;default:1" json:"vip_type"`           // 会员类型，1普通会员；2充值会员
	ActualBalance  uint32       `gorm:"column:actual_balance;type:decimal(10,2) unsigned;not null" json:"actual_balance"`   // 实际余额
	PresentBalance uint32       `gorm:"column:present_balance;type:decimal(10,2) unsigned;not null" json:"present_balance"` // 赠送余额
}
