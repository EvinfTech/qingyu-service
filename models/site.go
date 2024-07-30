package models

import "time"

type Site struct {
	Id           int       `json:"id"`             //自增id
	Name         string    `json:"name"`           //场地名称
	ShopId       int       `json:"shop_id"`        //场地id
	TimeEnum     string    `json:"time_enum"`      //时间枚举标签
	GmtCreate    time.Time `json:"gmt_create"`     //创建时间
	GmtUpdate    time.Time `json:"gmt_update"`     //更新时间
	Status       string    `json:"status"`         //状态
	FreePrice    int       `json:"free_price"`     //闲时价格
	BusyPrice    int       `json:"busy_price"`     //忙时价格
	BusyTimeEnum string    `json:"busy_time_enum"` //忙时时间段枚举标签
	WeekendBusy  int       `json:"weekend_busy"`   //周末是否为忙时
	HolidayBusy  int       `json:"holiday_busy"`   //假期是否为忙时
	DataBusy     string    `json:"data_busy"`      //特殊忙时日期
	Del          int       `json:"del"`            //删除标记 0:正常  1:删除
}
type SiteOrderDetail struct {
	SiteId   int    `json:"site_id"`
	SiteName string `json:"site_name"`
	TimeEnum []int  `json:"time_enum"`
	Money    int    `json:"money"`
}
