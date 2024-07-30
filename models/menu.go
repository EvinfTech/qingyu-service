package models

import "time"

type Menu struct {
	Id        int       `json:"id"`         //自增id
	Pid       int       `json:"pid"`        //父id
	Name      string    `json:"name"`       //名称
	PName     string    `json:"p_name"`     //父名称
	Path      string    `json:"path"`       //路径
	Icon      string    `json:"icon"`       //图标
	Hidden    int       `json:"hidden"`     //是否隐藏 0:不隐藏 1:隐藏
	Title     string    `json:"title"`      //中文标题
	Component string    `json:"component"`  //视图路径
	GmtCreate time.Time `json:"gmt_create"` //创建时间
	GmtUpdate time.Time `json:"gmt_update"` //更新时间
	Grade     int       `json:"grade"`      //菜单等级
}
