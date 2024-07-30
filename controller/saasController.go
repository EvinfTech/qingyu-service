package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"qiyu/dao"
	"qiyu/kit"
	"qiyu/logic"
	"qiyu/models"
	"qiyu/models/enum"
	"qiyu/models/param"
	"qiyu/util"
	"strings"
	"time"
)

func SaasAddShop(c *gin.Context) {
	shopParam := param.SaasAddShopParam{}
	err := c.BindJSON(&shopParam)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	marshal, _ := json.Marshal(&shopParam.Tag)
	bytes, _ := json.Marshal(shopParam.Photo)
	err = db.Create(&models.Shop{
		Name:      shopParam.Name,
		Avatar:    shopParam.Avatar,
		Photo:     string(bytes),
		Address:   shopParam.Address,
		Longitude: shopParam.Longitude,
		Latitude:  shopParam.Latitude,
		Tag:       string(marshal),
		Phone:     shopParam.Phone,
		WorkTime:  shopParam.WorkTime,
		SiteCount: shopParam.SiteCount,
		Facility:  shopParam.Facility,
		Serve:     shopParam.Serve,
		GmtCreate: util.Now(),
		GmtUpdate: util.Now(),
	}).Error
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "操作失败:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// SaasGetShopDetail 获取门店详情
func SaasGetShopDetail(c *gin.Context) {
	param := param.SaasGetShopDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	shop := models.Shop{}
	db.Where("id = ?", param.ShopId).First(&shop)
	type SaasGetShopDetailVO struct {
		Name      string   `json:"name"`      //门店名称
		Avatar    string   `json:"avatar"`    //头像
		Photo     []string `json:"photo"`     //门店照片
		Address   string   `json:"address"`   //门店地址
		Longitude float64  `json:"longitude"` //经度
		Latitude  float64  `json:"latitude"`  //维度
		Tag       []string `json:"tag"`       //门店标签
		Phone     string   `json:"phone"`     //门店电话
		WorkTime  string   `json:"work_time"` //营业时间
		Facility  string   `json:"facility"`  //设施
		Serve     string   `json:"serve"`     //服务
		Desc      string   `json:"desc"`      //描述
	}
	var photo []string
	json.Unmarshal([]byte(shop.Photo), &photo)
	var tag []string
	json.Unmarshal([]byte(shop.Tag), &tag)
	vo := SaasGetShopDetailVO{
		Name:      shop.Name,
		Avatar:    shop.Avatar,
		Photo:     photo,
		Address:   shop.Address,
		Longitude: shop.Longitude,
		Latitude:  shop.Latitude,
		Tag:       tag,
		Phone:     shop.Phone,
		WorkTime:  shop.WorkTime,
		Facility:  shop.Facility,
		Serve:     shop.Serve,
		Desc:      shop.Desc,
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", vo)

}

// SaasUpdateShopDetail 更新门店信息
func SaasUpdateShopDetail(c *gin.Context) {
	param := param.SaasUpdateShopDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	photo, _ := json.Marshal(&param.ShopPhoto)
	tag, _ := json.Marshal(&param.Tag)
	shop := models.Shop{
		Id:        param.ShopId,
		Name:      param.ShopName,
		Avatar:    param.ShopAvatar,
		Photo:     string(photo),
		Address:   param.ShopAddress,
		Tag:       string(tag),
		Phone:     param.ShopPhone,
		WorkTime:  param.WorkTime,
		Desc:      param.Desc,
		Facility:  param.Facility,
		Serve:     param.Serve,
		GmtUpdate: util.Now(),
	}
	db.Table("shop").Where("id = ?", param.ShopId).Updates(&shop)
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasAddSite 添加场地
func SaasAddSite(c *gin.Context) {
	param := param.SaasAddSiteParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}

	siteTimeEnum, err := GetWorkTimeEnum(param.SiteStartTime, param.SiteEndTime)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "时间工具发生错误", nil)
		return

	}
	timeEnum, err := json.Marshal(&siteTimeEnum)

	shop := models.Shop{}
	db.First(&shop)
	// 获取所有商户的所有场地数量
	count, err := dao.SiteCount()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取场地数量错误:"+err.Error(), nil)
		return
	}

	if shop.BottomPrice == 0 {
		shop.BottomPrice = param.FreePrice
		shop.SiteCount = int(count + 1)
		db.Updates(&shop)
	}

	var dataBusy, busyTimeEnum []byte
	if len(param.DataBusy) != 0 {
		dataBusy, _ = json.Marshal(param.DataBusy)
	}
	if len(param.BusyEndTime) != 0 && len(param.BusyStartTime) != 0 {
		BusyTimeEnum, busyErr := GetWorkTimeEnum(param.BusyStartTime, param.BusyEndTime)
		if busyErr != nil {
			util.ResponseKit(c, http.StatusInternalServerError, "时间工具发生错误", nil)
			return
		}
		busyTimeEnum, _ = json.Marshal(&BusyTimeEnum)
	}

	if err = dao.SiteAdd(param, shop.Id, timeEnum, dataBusy, busyTimeEnum); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "添加场地失败", nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// SaasGetSiteList 获取场地列表
func SaasGetSiteList(c *gin.Context) {
	param := param.SaasGetSiteListParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "请求失败", nil)
		return
	}

	// 获取所有场地
	siteAll, err := dao.SiteAll(param.ShopId)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取所有场地失败"+err.Error(), nil)
		return
	}

	type SaasGetSiteListVO struct {
		Id            int      `json:"id"`              //自增id
		Name          string   `json:"name"`            //场地名称
		ShopId        int      `json:"shop_id"`         //场地id
		SiteStartTime string   `json:"site_start_time"` //时间枚举标签
		SiteEndTime   string   `json:"site_end_time"`   //时间枚举标签
		Status        string   `json:"status"`          //状态
		FreePrice     int      `json:"free_price"`      //闲时价格
		BusyPrice     int      `json:"busy_price"`      //忙时价格
		BusyTimeEnum  []int    `json:"busy_time_enum"`  //忙时时间段枚举标签
		BusyStartTime string   `json:"busy_start_time"` //忙时开始时间
		BusyEndTime   string   `json:"busy_end_time"`   //忙时结束时间
		WeekendBusy   int      `json:"weekend_busy"`    //周末是否为忙时
		HolidayBusy   int      `json:"holiday_busy"`    //假期是否为忙时
		DataBusy      []string `json:"data_busy"`       //特殊忙时日期
	}
	var SaasGetSiteListVOList []SaasGetSiteListVO

	for _, siteOne := range siteAll {
		var time []int
		json.Unmarshal([]byte(siteOne.TimeEnum), &time)
		var busyTime []int
		json.Unmarshal([]byte(siteOne.BusyTimeEnum), &busyTime)
		var dataBusy []string
		json.Unmarshal([]byte(siteOne.DataBusy), &dataBusy)
		startTime := strings.Split(TimeEnum[time[0]], "~")
		endTime := strings.Split(TimeEnum[time[len(time)-1]], "~")

		busyStartTime := ""
		busyEndTime := ""
		if len(busyTime) != 0 {
			busyStartTime = strings.Split(TimeEnum[busyTime[0]], "~")[0]
			busyEndTime = strings.Split(TimeEnum[busyTime[len(busyTime)-1]], "~")[1]

		}
		SaasGetSiteListVOList = append(SaasGetSiteListVOList, SaasGetSiteListVO{
			Id:            siteOne.Id,
			Name:          siteOne.Name,
			SiteStartTime: startTime[0],
			SiteEndTime:   endTime[1],
			ShopId:        siteOne.ShopId,
			Status:        siteOne.Status,
			FreePrice:     siteOne.FreePrice,
			BusyPrice:     siteOne.BusyPrice,
			BusyStartTime: busyStartTime,
			BusyEndTime:   busyEndTime,
			WeekendBusy:   siteOne.WeekendBusy,
			HolidayBusy:   siteOne.HolidayBusy,
			DataBusy:      dataBusy,
		})
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", SaasGetSiteListVOList)
}

// SaasUpdateSite 更新场地信息
func SaasUpdateSite(c *gin.Context) {
	var param param.SaasUpdateSiteParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}

	enum, err := GetWorkTimeEnum(param.SiteStartTime, param.SiteEndTime)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "时间工具发生错误", nil)
		return
	}
	enumBytes, _ := json.Marshal(&enum)
	var busyEnumBytes []byte
	if len(param.BusyStartTime) != 0 && len(param.BusyEndTime) != 0 {
		busyEnum, errBusyEnum := GetWorkTimeEnum(param.BusyStartTime, param.BusyEndTime)
		if errBusyEnum != nil {
			util.ResponseKit(c, http.StatusInternalServerError, "时间工具发生错误", nil)
			return
		}
		busyEnumBytes, _ = json.Marshal(&busyEnum)
	}

	siteOne, err := dao.SiteGet(param.SiteId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusInternalServerError, "场地不存在", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查找场地错误:"+err.Error(), nil)
		}
		return
	}
	if err = dao.SiteUpdate(siteOne, param, enumBytes, busyEnumBytes); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "更新场地信息错误:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}
func SaasCancelOrder(c *gin.Context) {
	param := param.WxCancelOrderParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	if !CheckCancelOrder(param.OrderNo) {
		util.ResponseKit(c, http.StatusInternalServerError, "该订单不能取消", nil)
		return
	}
	db.Table("order").Where("order_no = ?", param.OrderNo).Updates(map[string]interface{}{"status": "C"})
	// 场地占用状态批量改为取消占用
	if err = dao.SiteUseUpdateNot(param.OrderNo); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "取消订单失败:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// SaasAddOrder 后台添加订单
func SaasAddOrder(c *gin.Context) {
	var param param.SaasAddOrderParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}

	mender, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}
	//加锁以后查一下场地
	//todo lock.Lock() 等以后再放开吧，现在走流程
	// TODO 因为有微信订场接口，所以此处简单加锁不能解决问题
	var siteIdSlice []int
	for _, v := range param.SiteDetail {
		siteIdSlice = append(siteIdSlice, v.SiteId)
	}
	//todo 这里需要检测时间
	gmtSiteUse, err := util.StringToTime("2006-01-02", param.GmtSiteUse)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "时间格式发生错误", nil)
		return

	}
	//如果对应的日期有场地，则进行判断，反之不进行判断
	siteUseList, err := dao.SiteUseListBySiteIdSlice(param.GmtSiteUse, siteIdSlice)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "查找场地占用失败:"+err.Error(), nil)
		return
	}
	for paramK, paramV := range param.SiteDetail {
		for siteUseK, siteUserV := range siteUseList {
			if paramK == siteUseK {
				var a []int
				json.Unmarshal([]byte(siteUserV.UseTimeEnum), &a)
				for _, v := range a {
					for _, vv := range paramV.TimeEnum {
						if v == vv {
							util.ResponseKit(c, http.StatusInternalServerError, util.ToString(paramV.SiteId)+"场地的"+TimeEnum[vv]+"时间被占", nil)
							return
						}
					}
				}
			}
		}

	}

	var (
		money   int
		orderNo string
	)
	err = db.Transaction(func(tx *gorm.DB) error {
		var (
			userOuid     string
			reserveName  string
			reservePhone string
			user         *models.User
		)
		if param.PayMode == dao.PayModeBalance {
			userOuid = param.UserOuid
			user, err = dao.UserGetOne(param.UserOuid)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = errors.New("会员不存在")
				} else {
					err = errors.New("查找会员失败:" + err.Error())
				}
				return err
			}
			reserveName = user.Name
			reservePhone = user.Phone
		} else {
			reserveName = param.ReserveName
			reservePhone = param.ReservePhone
		}
		money, orderNo, err = AddOrderAndSiteUse(tx, param.SiteDetail, userOuid, gmtSiteUse, "预约订场", enum.OrderTypeSaas, reserveName, reservePhone, param.Remark)
		if err != nil {
			return err
		}
		if param.PayMode == dao.PayModeBalance {
			if err = logic.UserMoneyPay(tx, user, money, orderNo, uint32(mender.Id)); err != nil {
				return err
			}
		}
		tx.Table("order").Where("order_no = ?", orderNo).Updates(map[string]interface{}{"user_name": param.UserName, "user_phone": param.UserPhone, "remark": param.Remark})
		return nil
	})

	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "订单创建失败:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", gin.H{"money": money, "order": orderNo})
}

// SaasGetOrderList 查看订单列表
func SaasGetOrderList(c *gin.Context) {
	param := param.SaasGetOrderListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}

	if param.Size == 0 {
		param.Size = 10
	}
	if param.Page == 0 {
		param.Page = 1
	}
	var total int64
	var orderList []models.Order
	begin := db
	if len(param.OrderNo) != 0 {
		begin = begin.Where("order_no = ?", param.OrderNo)
	}
	if param.Type != 0 {
		begin = begin.Where("type = ?", param.Type)
	}
	if len(param.Status) != 0 {
		begin = begin.Where("status = ?", param.Status)
	}
	if len(param.UserName) != 0 {
		begin = begin.Where("user_name like ?", "%"+param.UserName+"%")
	}
	if param.StartTime != "" && param.EndTime != "" {
		_, err = util.StringToTime("2006-01-02", param.StartTime)
		if err != nil {
			util.ResponseKit(c, http.StatusInternalServerError, "开始时间格式错误:应该为2003-01-02", nil)
			return
		}
		_, err = util.StringToTime("2006-01-02", param.EndTime)
		if err != nil {
			util.ResponseKit(c, http.StatusInternalServerError, "结束时间格式错误:应该为2003-01-02", nil)
			return
		}

		//var yesterdayOrderCount int64
		orderNoSlice := dao.SiteUseOrderNoSlice(param.StartTime, param.EndTime)
		begin = begin.Where("order_no in ?", orderNoSlice)
	}
	begin.Table("order").Where("shop_id = ?", param.ShopId).Count(&total)
	begin.Where("shop_id = ?", param.ShopId).Scopes(util.PageKit(param.Page, param.Size)).Order("gmt_create DESC").Find(&orderList)
	type SaasGetOrderListVO struct {
		OrderNo      string                   `json:"order_no"`
		UserOuid     string                   `json:"user_ouid"`
		UserName     string                   `json:"user_name"`
		Type         int                      `json:"type"`
		Status       string                   `json:"status"`
		Money        int                      `json:"money"`
		GmtCreate    string                   `json:"gmt_create"`
		SiteDetail   []models.SiteOrderDetail `json:"site_detail"`
		UserPhone    string                   `json:"user_phone"`
		Remark       string                   `json:"remark"`
		GmtSiteUse   string                   `json:"gmt_site_use"`
		ReserveName  string                   `json:"reserve_name"`
		ReservePhone string                   `json:"reserve_phone"`
	}

	var VOList []SaasGetOrderListVO
	for _, v := range orderList {
		var detailList []models.SiteOrderDetail
		json.Unmarshal([]byte(v.SiteDetail), &detailList)
		user := models.User{}
		db.Where("ouid = ?", v.UserOuid).First(&user)
		VOList = append(VOList, SaasGetOrderListVO{
			OrderNo:      v.OrderNo,
			UserOuid:     v.UserOuid,
			UserName:     user.Name,
			Type:         v.Type,
			Status:       v.Status,
			Money:        v.Money,
			GmtCreate:    v.GmtCreate.Format("2006-01-02 15:04:05"),
			SiteDetail:   detailList,
			UserPhone:    v.UserPhone,
			Remark:       v.Remark,
			GmtSiteUse:   v.ReserveDate.Format("2006-01-02"),
			ReserveName:  v.ReserveName,
			ReservePhone: v.ReservePhone,
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"total": total, "list": VOList})
}

// SaasUpdateremark 修改备注
func SaasUpdateRemark(c *gin.Context) {
	param := param.SaasUpdateRemarkParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	db.Table("order").Where("order_no = ?", param.OrderNo).Updates(map[string]interface{}{
		"remark":     param.Remark,
		"gmt_update": util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)

}

// SaasDelSite 删除场地
func SaasDelSite(c *gin.Context) {
	var param param.SaasDelSiteParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数错误:"+err.Error(), nil)
		return
	}

	// TODO 查看是否正在占用
	if err := dao.SiteDelete(param.SiteId); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "删除场地错误:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasAddRole 增加角色权限
func SaasAddRole(c *gin.Context) {
	param := param.SaasAddRoleParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	db.Create(&models.Role{
		Level:     param.Level,
		RoleName:  param.RoleName,
		GmtCreate: util.Now(),
		GmtUpdate: util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "添加成功", nil)

}

// GetFeedbackList 增加反馈列表
func GetFeedbackList(c *gin.Context) {
	param := param.GetFeedbackListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	var feedbackList []models.Feedback
	db.Scopes(util.PageKit(param.Page, param.Size)).Order("gmt_create DESC").Find(&feedbackList)
	var total int64
	db.Table("feedback").Count(&total)
	type GetFeedbackListVO struct {
		UserOuid   string   `json:"user_ouid"`   //用户ouid
		UserName   string   `json:"user_name"`   //用户名称
		UserAvatar string   `json:"user_avatar"` //用户头像
		Photo      []string `json:"photo"`       //照片
		Content    string   `json:"content"`     //内容
		GmtCreate  string   `json:"gmt_create"`  //创建时间
	}
	var VOList []GetFeedbackListVO
	for _, v := range feedbackList {
		vo := GetFeedbackListVO{
			UserOuid:   v.UserOuid,
			UserName:   v.UserName,
			UserAvatar: v.UserAvatar,
			Content:    v.Content,
			GmtCreate:  v.GmtCreate.Format("2006-01-02 15:04:05"),
		}
		json.Unmarshal([]byte(v.Photo), &vo.Photo)
		VOList = append(VOList, vo)
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"list": VOList, "total": total})

}

// SaasGetRoleList 获取所有角色名称
func SaasGetRoleList(c *gin.Context) {
	param := param.SaasGetRoleListParam{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	if param.Size == 0 {
		param.Size = 10
	}
	if param.Page == 0 {
		param.Page = 1
	}
	var roleList []models.Role
	var total int64
	db.Table("role").Where("del != 1").Count(&total)
	db.Where("del != 1").Scopes(util.PageKit(param.Page, param.Size)).Find(&roleList)
	type SaasGetRoleListVO struct {
		Id       int    `json:"id"`        //自增id
		RoleName string `json:"role_name"` //角色名称
		Level    int    `json:"level"`     //角色等级
	}
	var VOList []SaasGetRoleListVO
	for _, v := range roleList {
		VOList = append(VOList, SaasGetRoleListVO{
			Id:       v.Id,
			RoleName: v.RoleName,
			Level:    v.Level,
		})
	}
	util.ResponseKit(c, http.StatusOK, "查询成功", gin.H{"total": total, "list": VOList})
}

// SaasDelRole 删除角色
func SaasDelRole(c *gin.Context) {
	param := param.SaasDelRoleParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	db.Table("role").Where("id = ?", param.RoleId).Updates(map[string]interface{}{"del": 1, "gmt_update": util.Now()})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasUpdateRole 更新角色
func SaasUpdateRole(c *gin.Context) {
	param := param.SaasUpdateRoleParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Table("role").Where("id = ?", param.RoleId).Updates(map[string]interface{}{"role_name": param.RoleName, "level": param.Level, "gmt_update": util.Now()})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)

}

// SaasAddAdminUser 添加管理员用户
func SaasAddAdminUser(c *gin.Context) {
	param := param.SaasAddAdminUserParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	var count int64
	db.Table("admin_user").Where("name = ?", param.Name).Count(&count)
	if count > 0 {
		util.ResponseKit(c, http.StatusInternalServerError, "用户名已存在", nil)
		return
	}
	var role models.Role
	db.Where("id = ?", param.RoleId).First(&role)
	db.Create(&models.AdminUser{
		UserOuid:  param.UserOuid,
		Password:  param.Password,
		Avatar:    param.Avatar,
		Name:      param.Name,
		RoleId:    param.RoleId,
		RoleName:  role.RoleName,
		ShopId:    param.ShopId,
		GmtCreate: util.Now(),
		GmtUpdate: util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// SaasGetAdminUserList 查看所有管理用户
func SaasGetAdminUserList(c *gin.Context) {
	var adminUserList []models.AdminUser
	db.Where("del != 1").Find(&adminUserList)
	type SaasGetAdminUserListVO struct {
		Id       int    `json:"id"`        // 自增id
		UserOuid string `json:"user_ouid"` //用户唯一标识
		Avatar   string `json:"avatar"`    //头像
		Name     string `json:"name"`      //用户名
		RoleId   int    `json:"role_id"`   //用户角色
		RoleName string `json:"role_name"` //角色名称
	}
	var VOList []SaasGetAdminUserListVO
	for _, v := range adminUserList {
		VOList = append(VOList, SaasGetAdminUserListVO{
			Id:       v.Id,
			UserOuid: v.UserOuid,
			Avatar:   v.Avatar,
			Name:     v.Name,
			RoleId:   v.RoleId,
			RoleName: v.RoleName,
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", VOList)
}
func SaasGetAdminUserNameStatus(c *gin.Context) {
	param := param.SaasGetAdminUserNameStatusParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	var count int64
	db.Table("admin_user").Where("name = ?", param.Name).Count(&count)
	status := dao.PayStatusYes
	if count > 0 {
		status = dao.PayStatusNo
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", status)

}

// SaasUpdateAdminUser 更新管理用户信息
func SaasUpdateAdminUser(c *gin.Context) {
	param := param.SaasUpdateAdminUserParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	db.Table("admin_user").Where("user_ouid = ?", param.UserOuid).Updates(map[string]interface{}{
		"name":     param.Name,
		"password": param.Password,
		"role_id":  param.RoleId,
		"avatar":   param.Avatar,
	})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasDelAdminUser 删除管理员
func SaasDelAdminUser(c *gin.Context) {
	param := param.SaasDelAdminUserParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	db.Table("admin_user").Where("user_ouid = ?", param.UserOuid).Updates(map[string]interface{}{"del": 1, "gmt_create": util.Now()})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasAddMenu 增加菜单
func SaasAddMenu(c *gin.Context) {
	param := param.SaasAddMenuParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Create(&models.Menu{
		Pid:       param.Pid,
		Name:      param.Name,
		PName:     param.PName,
		Path:      param.Path,
		Hidden:    param.Hidden,
		Title:     param.Title,
		Component: param.Component,
		Icon:      param.Icon,
		GmtCreate: util.Now(),
		GmtUpdate: util.Now(),
		Grade:     param.Grade,
	})
	util.ResponseKit(c, http.StatusOK, "添加成功", nil)
}

// SaasGetMenuList 查看所有菜单列表
func SaasGetMenuList(c *gin.Context) {
	var menuList []models.Menu
	db.Find(&menuList)
	type SaasGetMenuListVO struct {
		Id        int    `json:"id"`        //自增id
		Pid       int    `json:"pid"`       //父id
		Name      string `json:"name"`      //名称
		PName     string `json:"p_name"`    //父名称
		Grade     int    `json:"grade"`     //菜单等级
		Path      string `json:"path"`      //路径
		Icon      string `json:"icon"`      //图标
		Hidden    int    `json:"hidden"`    //是否隐藏 0:不隐藏 1:隐藏
		Title     string `json:"title"`     //中文标题
		Component string `json:"component"` //视图路径
	}
	var VOList []SaasGetMenuListVO
	for _, v := range menuList {
		VOList = append(VOList, SaasGetMenuListVO{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:      v.Name,
			PName:     v.PName,
			Grade:     v.Grade,
			Path:      v.Path,
			Icon:      v.Icon,
			Hidden:    v.Hidden,
			Title:     v.Title,
			Component: v.Component,
		})

	}

	util.ResponseKit(c, http.StatusOK, "请求成功", VOList)
}

// SaasUpdateMenu 更新菜单
func SaasUpdateMenu(c *gin.Context) {
	param := param.SaasUpdateMenuParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Table("menu").Where("id = ?", param.Id).Updates(map[string]interface{}{"pid": param.Pid, "path": param.Path, "icon": param.Icon, "hidden": param.Hidden,
		"p_name": param.PName, "grade": param.Grade, "name": param.Name, "gmt_update": util.Now(), "title": param.Title, "component": param.Component})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasDelMenu 删除菜单
func SaasDelMenu(c *gin.Context) {
	param := param.SaasDelMenuParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Delete(&models.Menu{}, param.Id)
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasUpdateMenuRole 增加、更新角色菜单权限表
func SaasUpdateMenuRole(c *gin.Context) {
	param := param.SaasAddMenuRoleParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	role := models.MenuRole{}
	marshal, _ := json.Marshal(param.MenuIds)
	if errors.Is(db.Where("role_id = ?", param.RoleId).First(&role).Error, gorm.ErrRecordNotFound) {
		db.Create(&models.MenuRole{
			RoleId:    param.RoleId,
			MenuIds:   string(marshal),
			GmtCreate: util.Now(),
			GmtUpdate: util.Now(),
		})
	} else {
		db.Table("menu_role").Where("role_id = ?", param.RoleId).Updates(map[string]interface{}{
			"menu_ids":   string(marshal),
			"gmt_update": util.Now(),
		})
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasGetMenuRoleList 获取所有角色菜单列表
func SaasGetMenuRoleList(c *gin.Context) {
	param := param.SaasGetMenuRoleListParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	if param.Size == 0 {
		param.Size = 10
	}
	if param.Page == 0 {
		param.Page = 1
	}
	type SaasGetMenuRoleListVO struct {
		MenuId   int    `json:"menu_id"`
		MenuName string `json:"menu_name"`
	}
	role := models.MenuRole{}
	var total int64
	db.Table("role").Where("role_id = ?", param.RoleId).Count(&total)
	db.Where("role_id = ?", param.RoleId).Scopes(util.PageKit(param.Page, param.Size)).First(&role)
	var a []int
	json.Unmarshal([]byte(role.MenuIds), &a)

	var VOList []SaasGetMenuRoleListVO
	for _, v := range a {
		menu := models.Menu{}
		if !errors.Is(db.Where("id = ?", v).First(&menu).Error, gorm.ErrRecordNotFound) {
			VOList = append(VOList, SaasGetMenuRoleListVO{
				MenuId:   menu.Id,
				MenuName: menu.Name,
			})
		}

	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"total": total, "list": VOList})
}

func SaasDelMenuRole(c *gin.Context) {
	param := param.SaasDelMenuRoleVOParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Delete(&models.MenuRole{}, param.Id)
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)
}

// SaasLogin 登录
func SaasLogin(c *gin.Context) {
	var param param.SaasLoginParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}

	adminUser, err := dao.AdminUserGet(param.Name, param.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusBadRequest, "用户名密码不对", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查询管理员错误："+err.Error(), nil)
		}
		return
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"user_ouid": adminUser.UserOuid, "avatar": adminUser.Avatar})
	return
}

// SaasUpdatePassword 修改密码
func SaasUpdatePassword(c *gin.Context) {
	var param param.SaasUpdatePassword
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}

	adminUser, err := dao.AdminUserGet(param.Name, param.OldPassword)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusBadRequest, "用户名密码不对", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查询管理员错误："+err.Error(), nil)
		}
		return
	}

	if err = dao.AdminUserUpdatePassword(adminUser, param.NewPassword); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "修改密码错误："+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
	return
}

// SaasGetSiteReserve  查看场馆内场地预定信息
func SaasGetSiteReserve(c *gin.Context) {
	param := param.WxGetSiteReserveParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	time, err := util.StringToTime("2006-01-02", param.Date)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "时间格式有问题", nil)
		return
	}
	//todo 这里需要检测时间
	//time, err := util.StringToTime("2006-01-02", param.Date)
	//if err != nil {
	//	util.ResponseKit(c, http.StatusInternalServerError, "时间格式发生错误", nil)
	//	return
	//
	//}

	// 获取所有商户的所有场地
	siteAll, err := dao.SiteAllShop()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取所有场地失败"+err.Error(), nil)
		return
	}

	//场地价格
	type SitePrice struct {
		TimeEnum int `json:"time_enum"`
		Price    int `json:"price"`
	}
	type Site struct {
		SiteId                int         `json:"site_id"`
		SiteName              string      `json:"site_name"`
		Status                string      `json:"status"`
		OnlineReserveTimeEnum []int       `json:"online_reserve_time_enum"`
		SaasReserveTimeEnum   []int       `json:"saas_reserve_time_enum"`
		StoreTimeEnum         []int       `json:"store_time_enum"`
		Price                 []SitePrice `json:"price"`
	}

	var siteVO []Site

	for _, v := range siteAll {
		site := Site{
			SiteId:                v.Id,
			SiteName:              v.Name,
			Status:                v.Status,
			OnlineReserveTimeEnum: []int{},
			SaasReserveTimeEnum:   []int{},
		}
		//查看该场地都有哪些时间
		var timeEnum []int
		json.Unmarshal([]byte(v.TimeEnum), &timeEnum)
		site.StoreTimeEnum = timeEnum

		//计算每个时间的价格
		var sitePrice []SitePrice
		for _, vv := range timeEnum {
			sitePrice = append(sitePrice, SitePrice{
				TimeEnum: vv,
				Price:    SitePriceCalculate(v, time, vv),
			})
		}
		site.Price = sitePrice
		siteVO = append(siteVO, site)

	}

	siteUseList, err := dao.SiteUseList(param.Date)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "查找场地占用失败:"+err.Error(), siteVO)
		return
	}
	for _, v := range siteUseList {
		var a []int
		json.Unmarshal([]byte(v.UseTimeEnum), &a)
		for kk, vv := range siteVO {
			if v.SiteId == vv.SiteId {
				if v.Type == 1 {
					siteVO[kk].OnlineReserveTimeEnum = append(siteVO[kk].OnlineReserveTimeEnum, a...)
				} else {
					siteVO[kk].SaasReserveTimeEnum = append(siteVO[kk].SaasReserveTimeEnum, a...)
				}
			}
		}
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", siteVO)
}

// SaasGetUserList 查看所有用户
func SaasGetUserList(c *gin.Context) {
	var param param.SaasGetUserListParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}

	query := db.Model(&models.User{})
	if len(param.Name) > 0 {
		query = query.Where("name like ?", "%"+param.Name+"%")
	}
	if len(param.Phone) > 0 {
		query = query.Where("phone = ?", "%"+param.Phone+"%")
	}
	var total int64
	var userList []models.User
	query.Count(&total)
	query.Scopes(util.PageKit(param.Page, param.Size)).Find(&userList)
	type SaasGetUserListVO struct {
		Ouid           string       `json:"ouid"`         //用户ouid
		Name           string       `json:"name"`         //用户名
		Avatar         string       `json:"avatar"`       //用户头像
		Phone          string       `json:"phone"`        //用户手机号
		Birthday       int          `json:"birthday"`     //用户生日
		Status         string       `json:"status"`       //用户状态
		Sex            int          `json:"sex"`          //用户性别
		TotalCount     int          `json:"total_count"`  //统计运动次数
		TotalLength    int          `json:"total_length"` //统计运动时长
		GmtCreate      time.Time    `json:"gmt_create"`
		VipType        enum.VipType `json:"vip_type"`
		ActualBalance  uint32       `json:"actual_balance"`
		PresentBalance uint32       `json:"present_balance"`
	}
	var VOList []SaasGetUserListVO
	for _, v := range userList {
		VOList = append(VOList, SaasGetUserListVO{
			Ouid:           v.Ouid,
			Name:           v.Name,
			Avatar:         v.Avatar,
			Phone:          v.Phone,
			Birthday:       v.Birthday,
			Status:         v.Status,
			Sex:            v.Sex,
			TotalCount:     v.TotalCount,
			TotalLength:    v.TotalLength,
			GmtCreate:      v.GmtCreate,
			VipType:        v.VipType,
			ActualBalance:  v.ActualBalance,
			PresentBalance: v.PresentBalance,
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"total": total, "list": VOList})
}

// SaasUpdateUserInfo 修改会员信息
func SaasUpdateUserInfo(c *gin.Context) {
	param := param.SaasUpdateUserInfoParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Table("user").Where("ouid = ?", param.Ouid).Updates(map[string]interface{}{
		"name":       param.Name,
		"avatar":     param.Avatar,
		"phone":      param.Phone,
		"birthday":   param.Birthday,
		"status":     param.Status,
		"sex":        param.Sex,
		"gmt_update": util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// SaasUpdateAgreementAbout saas更新协议和关于我们
func SaasUpdateAgreementAbout(c *gin.Context) {
	param := param.SaasUpdateAgreementAboutParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	shop := models.Shop{}
	db.First(&shop)
	shop.AboutUs = param.AboutUs
	shop.Agreement = param.Agreement
	shop.GmtUpdate = util.Now()
	db.Save(&shop)
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)

}
func SaasGetAgreementAbout(c *gin.Context) {
	shop := models.Shop{}
	db.First(&shop)
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"agreement": shop.Agreement,
		"about_us":  shop.AboutUs,
	})
}

// SaasGetOperationalData saas查看首页运营数据
func SaasGetOperationalData(c *gin.Context) {
	shop := models.Shop{}
	if errors.Is(db.Take(&shop).Error, gorm.ErrRecordNotFound) {
		util.ResponseKit(c, http.StatusInternalServerError, "当前不存在门店，请先创建门店", nil)
		return
	}
	var todayMoney int //今日营业额
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("gmt_create > ? AND status = 'Y'", util.NewTime().NowStart()).Scan(&todayMoney)
	//var yesterdayMoney int //昨日营业额
	//db.Table("order").Select("COALESCE(SUM(money),0)").Where("gmt_create > ? AND gmt_create < ? AND status = 'Y'", util.GetZeroTime(util.Now().AddDate(0, 0, -1)), util.NewTime().NowStart()).Scan(&yesterdayMoney)

	//var yesterdayOrderCount int64
	todayOrderCount := dao.SiteUseTodayOrderCount()
	//db.Table("order").Where("gmt_create > ? AND gmt_create < ? AND status = 'Y'", util.GetZeroTime(util.Now().AddDate(0, 0, -1)), util.NewTime().NowStart()).Count(&yesterdayOrderCount)
	todaySiteUserCount, _ := getTodayAndYesterdaySiteUseCount()
	var todaySiteAveragePrice int //场均价格
	//var yesterdaySiteAveragePrice int //昨日场均价格
	if todayMoney != 0 && todaySiteUserCount != 0 {
		todaySiteAveragePrice = todayMoney / todaySiteUserCount
	}
	//if yesterdayMoney != 0 && yesterdaySiteUserCount != 0 {
	//	yesterdaySiteAveragePrice = yesterdayMoney / yesterdaySiteUserCount
	//}
	type SiteStatus struct {
		SiteId   int    `json:"site_id"`
		SiteName string `json:"site_name"` // 场地名称
		Status   string `json:"status"`    //状态
	}
	// 获取所有商户的所有场地
	siteAll, err := dao.SiteAllShop()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取场地失败"+err.Error(), nil)
		return
	}

	var siteStatusList []SiteStatus
	for _, v := range siteAll {
		if v.Status != "Y" {
			siteStatusList = append(siteStatusList, SiteStatus{
				SiteId:   v.Id,
				SiteName: v.Name,
				Status:   "维修中",
			})
			continue
		}
		siteStatusList = append(siteStatusList, SiteStatus{
			SiteId:   v.Id,
			SiteName: v.Name,
			Status:   "空闲",
		})
	}
	for _, v := range siteAll {
		if v.Status != "Y" {
			continue
		}
		siteUseList, err := dao.SiteUseListBySiteId(util.Now(), v.Id)
		if err != nil {
			util.ResponseKit(c, http.StatusInternalServerError, "查找场地占用失败:"+err.Error(), nil)
			return
		}
		for _, vv := range siteUseList {
			var jsonEnum []int
			json.Unmarshal([]byte(vv.UseTimeEnum), &jsonEnum)
			if len(jsonEnum) == 0 {
				continue
			}
			for _, vvv := range jsonEnum {
				split := strings.Split(TimeEnum[vvv], ":")
				if len(split) < 1 {
					continue
				}
				if util.Now().Hour() == util.StringToInt(split[0]) {
					for kkkk, vvvv := range siteStatusList {
						if vvvv.SiteId == vv.SiteId {
							siteStatusList[kkkk].Status = "使用中"
							continue
						}
					}
				}

			}
		}
	}
	weekSiteMoney := getWeekSiteMoney()
	var onLineCount int64
	var offlineCount int64
	db.Table("order").Where("status = 'Y' AND type = ?", enum.OrderTypeWx).Count(&onLineCount)
	db.Table("order").Where("status = 'Y' AND type = ?", enum.OrderTypeSaas).Count(&offlineCount)

	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"today_money": todayMoney,
		//"yesterday_money":              yesterdayMoney,
		"today_order_count": todayOrderCount,
		//"yesterday_order_count":        yesterdayOrderCount,
		"today_site_user_count": todaySiteUserCount,
		//"yesterday_site_user_count": yesterdaySiteUserCount,
		"today_site_average_price": todaySiteAveragePrice,
		//"yesterday_site_average_price": yesterdaySiteAveragePrice,
		"week_site_money":  weekSiteMoney,
		"site_status_list": siteStatusList,
		"service_date":     util.TimeSubDays(util.Now(), shop.GmtCreate),
		//"on_line_count":   onLineCount,
		//"off_line_count":  offlineCount,
	})

}

// 获取今日和昨日的场地占用数量
func getTodayAndYesterdaySiteUseCount() (todaySiteUserCount int, yesterdaySiteUserCount int) {
	now := util.Now()
	todaySiteUseList, _ := dao.SiteUseList(util.Date2String(now))
	for _, v := range todaySiteUseList {
		var a []int
		json.Unmarshal([]byte(v.UseTimeEnum), &a)
		todaySiteUserCount += len(a)
	}

	yesterdaySiteUseList, _ := dao.SiteUseList(util.Date2String(now.AddDate(0, 0, -1)))
	for _, v := range yesterdaySiteUseList {
		var a []int
		json.Unmarshal([]byte(v.UseTimeEnum), &a)
		yesterdaySiteUserCount += len(a)
	}
	return
}
func getWeekSiteMoney() map[string]interface{} {
	mm := make(map[string]interface{})
	mm[util.Now().Format("2006-01-02")] = getDaySiteMoney(util.Now())
	mm[util.Now().AddDate(0, 0, -1).Format("2006-01-02")] = getDaySiteMoney(util.Now().AddDate(0, 0, -1))
	mm[util.Now().AddDate(0, 0, -2).Format("2006-01-02")] = getDaySiteMoney(util.Now().AddDate(0, 0, -2))
	mm[util.Now().AddDate(0, 0, -3).Format("2006-01-02")] = getDaySiteMoney(util.Now().AddDate(0, 0, -3))
	mm[util.Now().AddDate(0, 0, -4).Format("2006-01-02")] = getDaySiteMoney(util.Now().AddDate(0, 0, -4))
	mm[util.Now().AddDate(0, 0, -5).Format("2006-01-02")] = getDaySiteMoney(util.Now().AddDate(0, 0, -5))
	mm[util.Now().AddDate(0, 0, -6).Format("2006-01-02")] = getDaySiteMoney(util.Now().AddDate(0, 0, -6))
	return mm
}

// 获取某一天场地的金额
func getDaySiteMoney(time time.Time) map[string]int {
	var orderList []models.Order
	db.Where("gmt_create > ? AND gmt_create < ?", util.GetZeroTime(time), util.GetZeroTime(time.AddDate(0, 0, 1))).Find(&orderList)
	// 获取所有商户的所有场地
	siteAll, _ := dao.SiteAllShop()

	m := make(map[string]int)
	for _, v := range siteAll {
		m[v.Name] = 0
	}
	for _, v := range orderList {
		var siteOrderDetail []models.SiteOrderDetail
		json.Unmarshal([]byte(v.SiteDetail), &siteOrderDetail)
		for _, vv := range siteOrderDetail {
			for kkk, _ := range m {
				if kkk == vv.SiteName {
					m[kkk] += vv.Money
				}
			}
		}
	}
	return m
}

// 获取某一天场地的金额
func getDaySiteCount(time time.Time) map[string]int {
	var orderList []models.Order
	db.Where("gmt_create > ? AND gmt_create < ?", util.GetZeroTime(time), util.GetZeroTime(time.AddDate(0, 0, 1))).Find(&orderList)
	// 获取所有商户的所有场地
	siteAll, _ := dao.SiteAllShop()

	m := make(map[string]int)
	for _, v := range siteAll {
		m[v.Name] = 0
	}
	for _, v := range orderList {
		var siteOrderDetail []models.SiteOrderDetail
		json.Unmarshal([]byte(v.SiteDetail), &siteOrderDetail)
		for _, vv := range siteOrderDetail {
			for kkk, _ := range m {
				if kkk == vv.SiteName {
					m[kkk] = m[kkk] + 1
				}
			}
		}
	}
	return m
}

// SaasGetShopCapital 获取门店资金流水
func SaasGetShopCapital(c *gin.Context) {
	var shop models.Shop
	db.Take(&shop)
	var superviseMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("status = 'Y'").Scan(&superviseMoney)

	var AggregateOrderMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Scan(&AggregateOrderMoney)
	var AggregateOrderPayMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("status = 'Y'").Scan(&AggregateOrderPayMoney)
	var AggregateOrderRefundMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("status = 'C'").Scan(&AggregateOrderRefundMoney)

	var TodayOrderMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("gmt_create > ?", util.NewTime().NowStart()).Scan(&TodayOrderMoney)
	var TodayOrderPayMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("status = 'Y' AND gmt_create >?", util.NewTime().NowStart()).Scan(&TodayOrderPayMoney)
	var TodayOrderRefundMoney int
	db.Table("order").Select("COALESCE(SUM(money),0)").Where("status = 'C' AND gmt_create >?", util.NewTime().NowStart()).Scan(&TodayOrderRefundMoney)

	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"service_date":                    util.TimeSubDays(util.Now(), shop.GmtCreate),
		"aggregate_amount":                shop.AggregateAmount,
		"balance":                         shop.Balance,
		"supervise_money":                 superviseMoney,
		"aggregate_volume_of_transaction": AggregateOrderMoney,
		"aggregate_order_money":           AggregateOrderMoney,
		"aggregate_order_pay_money":       AggregateOrderPayMoney,
		"aggregate_order_refund_money":    AggregateOrderRefundMoney,
		"today_volume_of_transaction":     TodayOrderMoney,
		"today_order_money":               TodayOrderMoney,
		"today_order_pay_money":           TodayOrderPayMoney,
		"today_order_refund_money":        TodayOrderRefundMoney,
	})

}

// SaasWithdrawDeposit 提现
func SaasWithdrawDeposit(c *gin.Context) {
	param := param.SaasWithdrawDepositParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	shop := models.Shop{}
	db.First(&shop)
	if shop.Balance == 0 {
		util.ResponseKit(c, http.StatusOK, "余额为0", nil)
		return
	}
	if shop.Balance < param.Money {
		util.ResponseKit(c, http.StatusOK, "余额不足", nil)
		return
	}
	db.Table("shop").Where("id = ?", shop.Id).Updates(map[string]interface{}{"gmt_update": util.Now(), "balance": shop.Balance - param.Money})
	db.Create(&models.ShopBill{
		TransactionSerialNo: kit.GenerateOrderNo(1),
		Type:                "T",
		Money:               -param.Money,
		Balance:             shop.Balance - param.Money,
		GmtCreate:           util.Now(),
		GmtUpdate:           util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)

}

// SaasGetBillList 查看账单流水
func SaasGetBillList(c *gin.Context) {
	param := param.SaasGetBillListParam{}
	c.BindJSON(&param)

	var total int64
	db.Model(&models.ShopBill{}).Count(&total)
	var billList []models.ShopBill
	db.Scopes(util.PageKit(param.Page, param.Size)).Order("gmt_create DESC").Find(&billList)
	type SaasGetBillListVo struct {
		Id                  int    `json:"id"`                    //自增id
		TransactionSerialNo string `json:"transaction_serial_no"` //交易流水号
		Type                string `json:"type"`                  //流水类型 (T:提现 J:结算)
		Money               int    `json:"money"`                 //流水金额
		Balance             int    `json:"balance"`               //余额
		Time                string `json:"time"`                  //时间
	}
	var voList []SaasGetBillListVo
	for _, v := range billList {
		voList = append(voList, SaasGetBillListVo{
			Id:                  v.Id,
			TransactionSerialNo: v.TransactionSerialNo,
			Type:                v.Type,
			Money:               v.Money,
			Balance:             v.Balance,
			Time:                v.GmtCreate.Format("2006-01-02 15:04:05"),
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"total": total,
		"list":  voList,
	})

}

// SaasCheckOrder 后台增加订单核验
func SaasCheckOrder(c *gin.Context) {
	param := param.SaasCheckOrderParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	order := models.Order{}
	db.Where("order_no = ?", param.OrderNo).First(&order)
	if order.Status != "Y" {
		util.ResponseKit(c, http.StatusOK, "状态不能进行核验", nil)
		return
	}
	siteTime, err := GetOrderEndSiteTime(param.OrderNo)
	if err != nil {
		util.ResponseKit(c, http.StatusOK, err.Error(), nil)
		return
	}
	if siteTime.Before(util.Now()) {
		util.ResponseKit(c, http.StatusOK, "核验失败，已经超时", nil)
		return
	}
	order.GmtUpdate = util.Now()
	order.Status = "U"
	if db.Updates(&order).RowsAffected == 0 {
		util.ResponseKit(c, http.StatusOK, "核验失败", nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "核验成功", nil)

}
