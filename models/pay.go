package models

import "time"

type Pay struct {
	Id              int       `json:"id"`               //自增id
	OrderNo         string    `json:"order_no"`         //订单编号
	PayType         string    `json:"pay_type"`         //支付类型
	TotalMoney      int       `json:"total_money"`      //订单总额
	Money           int       `json:"money"`            //实付金额
	Status          string    `json:"status"`           //支付状态
	ResponseContext string    `json:"response_context"` //第三方返回的信息
	PayForm         string    `json:"pay_form"`         //从谁转
	PayTo           string    `json:"pay_to"`           //转到谁
	UserOuid        string    `json:"user_ouid"`        //用户ouid
	UserName        string    `json:"user_name"`        //用户名称
	UserAvatar      string    `json:"user_avatar"`      //用户头像
	GmtResponse     time.Time `json:"gmt_response"`     //第三方返回时间
	GmtCreate       time.Time `json:"gmt_create"`       //创建时间
	GmtUpdate       time.Time `json:"gmt_update"`       //更新时间
	Remark          string    `json:"remark"`           //备注
	OutOrderNo      string    `json:"out_order_no"`     //外部订单号
	OrderName       string    `json:"order_name"`       //订单名称
	OrderId         int       `json:"order_id"`         //订单id
}

// Transaction 微信支付resource对象解密后的参数
type Transaction struct {
	Amount          TransactionAmount `json:"amount,omitempty"`
	Appid           string            `json:"appid,omitempty"`
	Attach          string            `json:"attach,omitempty"`
	BankType        string            `json:"bank_type,omitempty"`
	Mchid           string            `json:"mchid,omitempty"`
	OutTradeNo      string            `json:"out_trade_no,omitempty"`
	Payer           TransactionPayer  `json:"payer,omitempty"`
	PromotionDetail []PromotionDetail `json:"promotion_detail,omitempty"`
	SuccessTime     string            `json:"success_time,omitempty"`
	TradeState      string            `json:"trade_state,omitempty"`
	TradeStateDesc  string            `json:"trade_state_desc,omitempty"`
	TradeType       string            `json:"trade_type,omitempty"`
	TransactionId   string            `json:"transaction_id,omitempty"`
}

// TransactionAmount
type TransactionAmount struct {
	Currency      string `json:"currency,omitempty"`
	PayerCurrency string `json:"payer_currency,omitempty"`
	PayerTotal    int64  `json:"payer_total,omitempty"`
	Total         int64  `json:"total,omitempty"`
}
type TransactionPayer struct {
	Openid string `json:"openid,omitempty"`
}

// PromotionDetail
type PromotionDetail struct {
	// 券ID
	CouponId string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type string `json:"type,omitempty"`
	// 优惠券面额
	Amount int64 `json:"amount,omitempty"`
	// 活动ID，批次ID
	StockId string `json:"stock_id,omitempty"`
	// 单位为分
	WechatpayContribute int64 `json:"wechatpay_contribute,omitempty"`
	// 单位为分
	MerchantContribute int64 `json:"merchant_contribute,omitempty"`
	// 单位为分
	OtherContribute int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    string                 `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

// PromotionGoodsDetail
type PromotionGoodsDetail struct {
	// 商品编码
	GoodsId string `json:"goods_id"`
	// 商品数量
	Quantity int64 `json:"quantity"`
	// 商品价格
	UnitPrice int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark string `json:"goods_remark,omitempty"`
}
