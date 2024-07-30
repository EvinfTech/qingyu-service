package models

import "time"

type Role struct {
	Id        int       `json:"id"`         //自增id
	RoleName  string    `json:"role_name"`  //角色名称
	Del       int       `json:"del"`        //删除状态  0:正常 1:删除
	Level     int       `json:"level"`      //等级
	GmtCreate time.Time `json:"gmt_create"` //创建时间
	GmtUpdate time.Time `json:"gmt_update"` //更新时间
}
