package models

import "time"

type Shop struct {
	Id              int       `json:"id"`               //自增id
	Name            string    `json:"name"`             //门店名称
	Avatar          string    `json:"avatar"`           //头像
	Photo           string    `json:"photo"`            //门店照片
	Address         string    `json:"address"`          //门店地址
	Longitude       float64   `json:"longitude"`        //经度
	Latitude        float64   `json:"latitude"`         //维度
	Tag             string    `json:"tag"`              //门店标签
	Phone           string    `json:"phone"`            //门店电话
	WorkTime        string    `json:"work_time"`        //营业时间
	SiteCount       int       `json:"site_count"`       //场地数量
	BottomPrice     int       `json:"bottom_price"`     //场馆最低价格
	Desc            string    `json:"desc"`             //描述
	Facility        string    `json:"facility"`         //设施
	Serve           string    `json:"serve"`            //服务
	AggregateAmount int       `json:"aggregate_amount"` //总金额
	Balance         int       `json:"balance"`          //余额
	Agreement       string    `json:"agreement"`        //协议
	AboutUs         string    `json:"about_us"`         //关于我们
	GmtCreate       time.Time `json:"gmt_create"`       //创建时间
	GmtUpdate       time.Time `json:"gmt_update"`       //更新时间
}
