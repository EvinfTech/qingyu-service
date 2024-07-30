package param

type Recharge struct {
	Ouid  string `json:"ouid" binding:"required"`  // 用户ouid
	Money uint32 `json:"money" binding:"required"` // 充值金额
}

type WxRecharge struct {
	Type string `json:"type" binding:"required"` // web  or h5
	Recharge
}

type BalancePay struct {
	UserOuid string `json:"user_ouid" binding:"required"`
	OrderNo  string `json:"order_no" binding:"required"`
}
