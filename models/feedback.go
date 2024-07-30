package models

import "time"

type Feedback struct {
	Id         int       `json:"id"`          //自增id
	UserOuid   string    `json:"user_ouid"`   //用户ouid
	UserName   string    `json:"user_name"`   //用户名称
	UserAvatar string    `json:"user_avatar"` //用户头像
	Photo      string    `json:"photo"`       //照片
	Content    string    `json:"content"`     //内容
	GmtCreate  time.Time `json:"gmt_create"`  //创建时间
	GmtUpdate  time.Time `json:"gmt_update"`  //更新时间
}
