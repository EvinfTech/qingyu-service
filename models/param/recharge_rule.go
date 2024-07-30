package param

type RechargeRuleAdd struct {
	OriginMoney  uint32 `json:"origin_money" binding:"required"`
	PresentMoney uint32 `json:"present_money" binding:"required"`
}
type RechargeRuleUpdate struct {
	Id uint32 `json:"id" binding:"required"`
	RechargeRuleAdd
}
