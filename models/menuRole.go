package models

import "time"

type MenuRole struct {
	Id        int       `json:"id"`         //自增id
	RoleId    int       `json:"role_id"`    //角色id
	MenuIds   string    `json:"menu_ids"`   //菜单id
	GmtCreate time.Time `json:"gmt_create"` //创建时间
	GmtUpdate time.Time `json:"gmt_update"` //更新时间
}
