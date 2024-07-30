package models

import (
	"time"
)

type Activity struct {
	Id               int       `json:"id"`                 //自增id
	Name             string    `json:"name"`               //活动名称
	Avatar           string    `json:"avatar"`             //活动头像
	CreateUserOuid   string    `json:"create_user_ouid"`   //发起用户ouid
	CreateUserName   string    `json:"create_user_name"`   //发起用户名称
	CreateUserAvatar string    `json:"create_user_avatar"` //发起用户头像
	CreateUserPhone  string    `json:"create_user_phone"`  //发起用户手机号
	ShopId           int       `json:"shop_id"`            //门店id
	ShopName         string    `json:"shop_name"`          //门店名称
	ShopAvatar       string    `json:"shop_avatar"`        //门店名称
	SiteDetail       string    `json:"site_detail"`        //场地信息
	ShopAddress      string    `json:"shop_address"`       //门店详细地址
	Longitude        float64   `json:"longitude"`          //经度
	Latitude         float64   `json:"latitude"`           //维度
	ActivityDate     time.Time `json:"activity_date"`      //活动日期
	ActivityTime     string    `json:"activity_time"`      //活动时间段
	PriceMan         int       `json:"price_man"`          //男生单价
	PriceWoman       int       `json:"price_woman"`        //女生单价
	TotalCount       int       `json:"total_count"`        //总人数
	ApplyCount       int       `json:"apply_count"`        //已申请人数
	remark           string    `json:"remark"`             //活动备注
	Status           string    `json:"status"`             //活动状态
	GmtCreate        time.Time `json:"gmt_create"`         //创建时将
	GmtUpdate        time.Time `json:"gmt_update"`         //更新时间
}
