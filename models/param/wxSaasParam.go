package param

type SaasAddShopAuditParam struct {
	Name            string  `json:"name"`             // 名字
	Address         string  `json:"address"`          // 门店地址
	Longitude       float64 `json:"longitude"`        // 经度
	Latitude        float64 `json:"latitude"`         // 维度
	SiteCount       int     `json:"site_count"`       // 场地数量
	GatePhoto       string  `json:"gate_photo"`       // 门口照片
	IndoorPhoto     string  `json:"indoor_photo"`     // 室内照片
	BusinessLicense string  `json:"business_license"` // 营业执照
	IdCardFront     string  `json:"id_card_front"`    // 身份证正面
	IdCardBack      string  `json:"id_card_back"`     // 身份证反面
	UserOuid        string  `json:"user_ouid"`        // 用户ouid
	UserName        string  `json:"user_name"`        // 联系人名称
	UserPhone       string  `json:"user_phone"`       // 联系人电话
	Desc            string  `json:"desc"`             // 备注
}
type SaasWxUpdateShopDetailParam struct {
	ShopId    int      `json:"shop_id"`
	Avatar    string   `json:"avatar"`
	Name      string   `json:"name"`
	Photo     []string `json:"photo"`
	Longitude float64  `json:"longitude"`
	Latitude  float64  `json:"latitude"`
	Address   string   `json:"address"`
	Phone     string   `json:"phone"`
	WorkTime  string   `json:"work_time"`
	Desc      string   `json:"desc"`
}

type SaasWxGetOrderListParam struct {
	ShopId int    `json:"shop_id" binding:"required"`
	Status string `json:"status"`
}
type SaasWxGetOrderDetailParam struct {
	OrderNO string `json:"order_no" binding:"required"`
}
type SaasWxGetEarningsParam struct {
	ShopId int `json:"shop_id" binding:"required"`
}
type SaasWxGetActivityListParam struct {
	ShopId int `json:"shop_id" binding:"required"`
}
type SaasWxGetActivityDetailParam struct {
	ActivityId int `json:"activity_id"`
}
type SaasWxGetSiteListParam struct {
	ShopId int `json:"shop_id"`
}
type SaasWxGetSiteReserveParam struct {
	SiteId int    `json:"site_id" binding:"required"`
	Date   string `json:"date" binding:"required"` //想要查看的时间
}
