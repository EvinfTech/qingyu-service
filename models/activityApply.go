package models

import (
	"time"
)

type ActivityApply struct {
	Id               int       `json:"id"`                 //自增id
	ActivityId       int       `json:"activity_id"`        //活动id
	ActivityName     string    `json:"activity_name"`      //活动名称
	ActivityAvatar   string    `json:"activity_avatar"`    //活动头像
	CreateUserOuid   string    `json:"create_user_ouid"`   //活动发起人
	CreateUserName   string    `json:"create_user_name"`   //创始人名称
	CreateUserAvatar string    `json:"create_user_avatar"` //创始人头像
	CreateUserPhone  string    `json:"create_user_phone"`  //创始人手机号
	UserOuid         string    `json:"user_ouid"`          // 用户ouid
	UserName         string    `json:"user_name"`          //用户名称
	UserPhone        string    `json:"user_phone"`         //用户手机号
	UserAvatar       string    `json:"user_avatar"`        //用户头像
	UserSex          int       `json:"user_sex"`           //用户性别
	ManCount         int       `json:"man_count"`          //报名男人数
	OrderNo          string    `json:"order_no"`           //订单号
	WomanCount       int       `json:"woman_count"`        //报名女人数
	Status           string    `json:"status"`             //目前的状态 Y:已报名 N:已报名待支付 C:已取消 W:候选人 U:候选人待支付
	Remark           string    `json:"remark"`             //备注
	GmtCreate        time.Time `json:"gmt_create"`         //创建时间
	GmtUpdate        time.Time `json:"gmt_update"`         //更新时间
}
