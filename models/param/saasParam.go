package param

type SaasAddShopParam struct {
	Name      string   `json:"name" binding:"required"`      //门店名称
	Avatar    string   `json:"avatar" binding:"required"`    //头像
	Photo     []string `json:"photo"`                        //门店照片
	Address   string   `json:"address" binding:"required"`   //门店地址
	Longitude float64  `json:"longitude" binding:"required"` //经度
	Latitude  float64  `json:"latitude" binding:"required"`  //维度
	Tag       []string `json:"tag"`                          //门店标签
	Phone     string   `json:"phone"`                        //门店电话
	WorkTime  string   `json:"work_time"`                    //营业时间
	SiteCount int      `json:"site_count"`                   //场地数量
	Facility  string   `json:"facility"`                     //场馆设施
	Serve     string   `json:"serve"`                        //场馆服务
}
type SaasAddSiteParam struct {
	Name          string   `json:"name"` //场地名称
	SiteStartTime string   `json:"site_start_time" binding:"required"`
	SiteEndTime   string   `json:"site_end_time" binding:"required"`
	BusyStartTime string   `json:"busy_start_time"`
	BusyEndTime   string   `json:"busy_end_time"`
	Status        string   `json:"status"`       //状态
	FreePrice     int      `json:"free_price"`   //闲时价格
	BusyPrice     int      `json:"busy_price"`   //忙时价格
	WeekendBusy   int      `json:"weekend_busy"` //周末是否为忙时
	HolidayBusy   int      `json:"holiday_busy"` //假期是否为忙时
	DataBusy      []string `json:"data_busy"`    //特殊忙时日期

}
type GetFeedbackListParam struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
type SaasUpdateAgreementAboutParam struct {
	Agreement string `json:"agreement" binding:"required"`
	AboutUs   string `json:"about_us" binding:"required"`
}
type SaasAddOrderParam struct {
	UserName     string       `json:"user_name"`
	UserPhone    string       `json:"user_phone"`
	Remark       string       `json:"remark"`
	ShopId       int          `json:"shop_id"`
	SiteDetail   []SiteDetail `json:"site_detail"`
	PayMode      string       `json:"pay_mode"`
	UserOuid     string       `json:"user_ouid"`
	ReserveName  string       `json:"reserve_name"`                    //预约手机号
	ReservePhone string       `json:"reserve_phone"`                   //预约手机
	GmtSiteUse   string       `json:"gmt_site_use" binding:"required"` //场地占用日期  (2020-05-05)
}
type SaasUpdateRemarkParam struct {
	OrderNo string `json:"order_no" binding:"required"`
	Remark  string `json:"remark" binding:"required"`
}
type SaasGetOrderListParam struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	ShopId    int    `json:"shop_id" binding:"required"`
	OrderNo   string `json:"order_no"` //
	Type      int    `json:"type"`
	Status    string `json:"status"`
	UserName  string `json:"user_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
type SaasGetShopDetailParam struct {
	ShopId int `json:"Shop_id"`
}
type SaasUpdateShopDetailParam struct {
	ShopId      int      `json:"id"`
	ShopName    string   `json:"name"`
	ShopAvatar  string   `json:"avatar"`
	ShopPhoto   []string `json:"photo"`
	ShopPhone   string   `json:"phone"`
	ShopAddress string   `json:"address"`
	WorkTime    string   `json:"work_time"`
	Tag         []string `json:"tag"`
	Desc        string   `json:"desc"`
	Facility    string   `json:"facility"` //场馆设施
	Serve       string   `json:"serve"`    //场馆服务
}
type SaasGetSiteListParam struct {
	ShopId int `json:"shop_id"`
}
type SaasUpdateSiteParam struct {
	SiteId        int    `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	SiteStartTime string `json:"site_start_time" binding:"required"`
	SiteEndTime   string `json:"site_end_time" binding:"required"`
	BusyStartTime string `json:"busy_start_time"`
	BusyEndTime   string `json:"busy_end_time"`
	FreePrice     int    `json:"free_price" binding:"required"`
	BusyPrice     int    `json:"busy_price" binding:"required"`
	Status        string `json:"status" binding:"required"`
	WeekendBusy   int    `json:"weekend_busy" `
	HolidayBusy   int    `json:"holiday_busy"`
}
type SaasDelSiteParam struct {
	SiteId int `json:"site_id"` //门店id
}
type SaasAddRoleParam struct {
	RoleName string `json:"role_name"`
	Level    int    `json:"level"`
}
type SaasDelRoleParam struct {
	RoleId int `json:"role_id" binding:"required"`
}
type SaasAddAdminUserParam struct {
	UserOuid string `json:"user_ouid" binding:"required"` //用户oud
	Name     string `json:"name" binding:"required"`      //姓名
	Password string `json:"password" binding:"required"`  //密码
	RoleId   int    `json:"role_id" binding:"required"`   //角色id
	Avatar   string `json:"avatar"`                       //头像
	ShopId   int    `json:"shop_id"`                      //门店id
}
type SaasGetAdminUserNameStatusParam struct {
	Name string `json:"name" binding:"required"`
}
type SaasUpdateAdminUserParam struct {
	UserOuid string `json:"user_ouid" binding:"required"` //用户oud
	Name     string `json:"name" binding:"required"`      //姓名
	Password string `json:"password" binding:"required"`  //密码
	RoleId   int    `json:"role_id" binding:"required"`   //角色id
	Avatar   string `json:"avatar"`                       //头像
}
type SaasDelAdminUserParam struct {
	UserOuid string `json:"user_ouid" binding:"required"`
}
type SaasAddMenuParam struct {
	Pid       int    `json:"pid"`                      //父id
	Name      string `json:"name" binding:"required"`  //名称
	PName     string `json:"p_name"`                   //父名称
	Grade     int    `json:"grade" binding:"required"` //菜单等级
	Path      string `json:"path" binding:"required"`  //路径
	Icon      string `json:"icon"`                     //图标
	Hidden    int    `json:"hidden"`                   //是否隐藏 0:不隐藏  1:隐藏
	Title     string `json:"title"`                    //中文标题
	Component string `json:"component"`                //视图路径
}
type SaasUpdateMenuParam struct {
	Id        int    `json:"id"`                       //id
	Pid       int    `json:"pid"`                      //父id
	Name      string `json:"name" binding:"required"`  //名称
	PName     string `json:"p_name"`                   //父名称
	Grade     int    `json:"grade" binding:"required"` //菜单等级
	Path      string `json:"path" binding:"required"`  //路径
	Icon      string `json:"icon"`                     //图标
	Hidden    int    `json:"hidden"`                   //是否隐藏 0:不隐藏  1:隐藏
	Title     string `json:"title"`                    //中文标题
	Component string `json:"component"`                //视图路径
}
type SaasDelMenuParam struct {
	Id int `json:"id"`
}
type SaasAddMenuRoleParam struct {
	RoleId  int   `json:"role_id" binding:"required"`
	MenuIds []int `json:"menu_ids" binding:"required"`
}
type SaasDelMenuRoleVOParam struct {
	Id int `json:"id" binding:"required"`
}
type SaasLoginParam struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SaasUpdatePassword struct {
	Name        string `json:"name" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type SaasGetMenuRoleListParam struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	RoleId int `json:"role_id"`
}
type SaasUpdateRoleParam struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name" binding:"required"` //角色名称
	Level    int    `json:"level" binding:"required"`     //等级
}
type SaasGetUserListParam struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type SaasUpdateUserInfoParam struct {
	Ouid     string `json:"ouid"`     //用户ouid
	Name     string `json:"name"`     //用户名
	Avatar   string `json:"avatar"`   //用户头像
	Phone    string `json:"phone"`    //用户手机号
	Birthday int    `json:"birthday"` //用户生日
	Status   string `json:"status"`   //用户状态
	Sex      int    `json:"sex"`      //用户性别

}

type SaasGetRoleListParam struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
type SaasWithdrawDepositParam struct {
	Money int `json:"money" binding:"required"`
}
type SaasGetBillListParam struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
type SaasCheckOrderParam struct {
	OrderNo string `json:"order_no" binding:"required"` //订单列表
	CheckNo string `json:"check_no" binding:"required"` //核验码
}
