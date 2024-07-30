package models

import "time"

type AdminUser struct {
	Id        int       `json:"id"`        // 自增id
	UserOuid  string    `json:"user_ouid"` //用户唯一标识
	Password  string    `json:"password"`  //密码
	Avatar    string    `json:"avatar"`    //头像
	Name      string    `json:"name"`      //用户名
	RoleId    int       `json:"role_id"`   //用户角色
	RoleName  string    `json:"role_name"` //角色名称
	ShopId    int       `json:"shop_id"`
	GmtCreate time.Time `json:"gmt_create"` //创建时间
	GmtUpdate time.Time `json:"gmt_update"` //更新时间
	Level     int       `json:"level"`      //等级
	Del       int       `json:"del"`        //删除状态 0:正常  1:已删除
}
