package models

import "time"

type Order struct {
	Id              int       `json:"id"`               //自增id
	OrderNo         string    `json:"order_no"`         //订单号
	OrderName       string    `json:"order_name"`       //订单名称
	OriginalPrice   int       `json:"original_price"`   //原价
	Coupons         int       `json:"coupons"`          //优惠卷价格
	Money           int       `json:"money"`            //实付价格
	PayType         string    `json:"pay_type"`         //支付方式
	ShopId          int       `json:"shop_id"`          //门店id
	ShopName        string    `json:"shop_name"`        //门店名称
	ShopAvatar      string    `json:"shop_avatar"`      //门店头像
	UserOuid        string    `json:"user_ouid"`        //用户ouid
	UserName        string    `json:"user_name"`        //用户名称
	UserAvatar      string    `json:"user_avatar"`      //用户头像
	UserPhone       string    `json:"user_phone"`       //用户手机号
	SiteDetail      string    `json:"site_detail"`      //订场详细信息
	PrePay          string    `json:"pre_pay"`          //预支付信息
	Status          string    `json:"status"`           //订单状态 Y:已支付 N:待支付  C:已取消  U:已使用
	Remark          string    `json:"remark"`           //备注
	Type            int       `json:"type"`             //1:线上  10：线下
	SuperviseStatus string    `json:"supervise_status"` //结算状态 N:未结算  Y:已结算
	CheckNo         string    `json:"check_no"`         //订单核验码
	CheckQr         string    `json:"check_qr"`         //核验二维码
	ReserveName     string    `json:"reserve_name"`     //预约手机号
	ReservePhone    string    `json:"reserve_phone"`    //预约手机
	ReserveDate     time.Time `json:"reserve_date"`     //预约日期
	GmtSuccess      time.Time `json:"gmt_success"`      //支付成功时间
	GmtRefund       time.Time `json:"gmt_refund"`       //退款成功时间
	GmtCreate       time.Time `json:"gmt_create"`       //创建时间
	GmtUpdate       time.Time `json:"gmt_update"`       //更新时间
}
