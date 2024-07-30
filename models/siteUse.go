package models

import "time"

type SiteUse struct {
	Id           int       `json:"id"`            //自增id
	SiteId       int       `json:"site_id"`       //场地id
	ShopId       int       `json:"shop_id"`       //门店id
	UseDate      string    `json:"use_date"`      //占用日期
	UseTimeEnum  string    `json:"use_time_enum"` //占用时间标签
	UserOuid     string    `json:"user_ouid"`     //用户ouid
	UserName     string    `json:"user_name"`     //用户名称
	UserAvatar   string    `json:"user_avatar"`   //用户头像
	UserPhone    string    `json:"user_phone"`    //用户手机号
	OrderNo      string    `json:"order_no"`      //订单编号
	OperatorOuid string    `json:"operator_ouid"` //操作人员的ouid
	Status       string    `json:"status"`        //占用状态 //Y:占用中 N:取消占用
	Type         int       `json:"type"`          //1:线上  10：线下
	GmtCreate    time.Time `json:"gmt_create"`    //创建时间
	GmtUpdate    time.Time `json:"gmt_update"`    //更新时间
}
