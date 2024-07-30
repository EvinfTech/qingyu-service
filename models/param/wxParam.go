package param

type GetOuidFromCodeParam struct {
	Code      string `json:"code" binding:"required"`
	PhoneCode string `json:"phone_code"`
}

// UserOuidParam 用户id
type UserOuidParam struct {
	UserOuid string `json:"user_ouid" binding:"required"`
}

type WxUpdateUserInfoParam struct {
	UserOuid  string `json:"user_ouid" binding:"required"` //用户ouid
	Name      string `json:"name"`                         //姓名
	Avatar    string `json:"avatar"`                       //用户头像
	Phone     string `json:"phone"`                        //手机号
	Birthday  int    `json:"birthday"`                     //生日
	Sex       int    `json:"sex"`                          //性别
	Introduce string `json:"introduce"`                    //个人简介
}

type SiteDetail struct {
	SiteId   int   `json:"site_id"`
	TimeEnum []int `json:"time_enum"`
}

type WxAddOrderParam struct {
	UserOuid     string       `json:"user_ouid" binding:"required"`
	SiteDetail   []SiteDetail `json:"site_detail" binding:"required"`
	GmtSiteUse   string       `json:"gmt_site_use" binding:"required"` //场地占用日期  (2020-05-05)
	ReserveName  string       `json:"reserve_name"`                    //预约手机号
	ReservePhone string       `json:"reserve_phone"`                   //预约手机
	Remark       string       `json:"remark"`                          //备注
}
type WxGetSiteReserveParam struct {
	Date string `json:"date" binding:"required"` //想要查看的时间
}
type WxGetOrderDetailParam struct {
	OrderNo string `json:"order_no"`
}
type WxAddFeedbackParam struct {
	UserOuid string   `json:"user_ouid" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Photo    []string `json:"photo"`
}
type WxCancelOrderParam struct {
	OrderNo string `json:"order_no" binding:"required"`
}
type WxAddActivityParam struct {
	ActivityName   string       `json:"activity_name"`    //活动名称
	ActivityAvatar string       `json:"activity_avatar"`  //活动头像
	TotalCount     int          `json:"total_count"`      //活动总计人数
	PriceMan       int          `json:"price_man"`        //男生单价
	PriceWoman     int          `json:"price_woman"`      //女生单价
	ShopId         int          `json:"shop_id"`          //门店id
	ActivityDate   string       `json:"activity_date"`    //活动时间（2023-11-11）
	ActivityTime   string       `json:"activity_time"`    //活动时间段(可以手填)
	CreateUserOuid string       `json:"create_user_ouid"` //创建者id
	SiteDetail     []SiteDetail `json:"site_detail"`      //场地详情
	remark         string       `json:"remark"`           //备注
}
type WxJoinActivityParam struct {
	ActivityId int    `json:"activity_id" binding:"required"` //活动id
	UserOuid   string `json:"user_ouid" binding:"required"`   // 用户ouid
	ManCount   int    `json:"man_count" `                     //报名男人数
	WomanCount int    `json:"woman_count"`                    //报名女人数
	remark     string `json:"remark"`                         //备注
	Phone      string `json:"phone"`
	Name       string `json:"name"` //名称
}
type WxGetActivityListParam struct {
	UserOuid   string `json:"user_ouid" binding:"required"`   //用户ouid
	ActiveDate string `json:"active_date" binding:"required"` //查询活动的时间
}
type WxGetActivityDetailParam struct {
	ActivityId int    `json:"activity_id" binding:"required"`
	UserOuid   string `json:"user_ouid" binding:"required"`
}
type WxGetActivityApplyListParam struct {
	ActivityId int `json:"activity_id" binding:"required"`
}
type WxCancelActivityApplyParam struct {
	UserOuid   string `json:"user_ouid" binding:"required"`
	ActivityId int    `json:"activity_id" binding:"required"`
}
type WxCancelActivityParam struct {
	UserOuid   string `json:"user_ouid" binding:"required"`
	ActivityId int    `json:"activity_id" binding:"required"`
}
type GetPhoneParam struct {
	PhoneCode string `json:"phone_code" binding:"required"`
}
type WxGetReserveDetailParam struct {
	UserOuid string `json:"user_ouid"`
	OrderNo  string `json:"order_no" binding:"required"`
}
type WxAddShopErrorParam struct {
	UserOuid string   `json:"user_ouid" binding:"required"`
	Type     string   `json:"type"`   //错误类型
	Detail   string   `json:"detail"` //详情
	Photos   []string `json:"photos"` //照片
}
type GetPhoneCodeParam struct {
	Phone string `json:"phone" binding:"required"` //手机号

}
type PhoneCodeLoginParam struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
type WxGetOrderListParam struct {
	UserOuid string `json:"user_ouid" binding:"required"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Status   string `json:"status"`
}
type WxGetMyReserveListParam struct {
	UserOuid string `json:"user_ouid" binding:"required"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Status   string `json:"status"`
}
type WxPayParam struct {
	OrderNo string `json:"order_no" binding:"required"`
	Type    string `json:"type" binding:"required"` //web  ro  h5
}
type WxGetWxIdParam struct {
	Type     string `json:"type"` //A:微信小程序  B:微信内置浏览器
	Code     string `json:"code"`
	UserOuid string `json:"user_ouid"`
}
