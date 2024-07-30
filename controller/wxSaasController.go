package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"net/http"
	"qiyu/dao"
	"qiyu/models"
	"qiyu/models/param"
	"qiyu/util"
)

// SaasAddShopAudit 提交审核
func SaasAddShopAudit(c *gin.Context) {
	param := param.SaasAddShopAuditParam{}
	audit := models.ShopAudit{}
	if !errors.Is(db.Where("user_ouid = ?", param.UserOuid).First(&audit).Error, gorm.ErrRecordNotFound) {
		util.ResponseKit(c, http.StatusInternalServerError, "用户已经申请过", gin.H{"status": audit.Status})
		return
	}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	db.Create(&models.ShopAudit{
		Name:            param.Name,
		Address:         param.Address,
		Longitude:       param.Longitude,
		Latitude:        param.Latitude,
		SiteCount:       param.SiteCount,
		GatePhoto:       param.GatePhoto,
		IndoorPhoto:     param.IndoorPhoto,
		BusinessLicense: param.BusinessLicense,
		IdCardFront:     param.IdCardFront,
		IdCardBack:      param.IdCardBack,
		UserOuid:        param.UserOuid,
		UserName:        param.UserName,
		UserPhone:       param.UserPhone,
		Desc:            param.Desc,
		Status:          dao.PayStatusNo,
		GmtCreate:       util.Now(),
		GmtUpdate:       util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}
func SaasWxGetShopAudit(c *gin.Context) {
	param := param.UserOuidParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "用户ouid不能为空", nil)
		return
	}
	audit := models.ShopAudit{}
	if errors.Is(db.Where("user_ouid = ?", param.UserOuid).First(&audit).Error, gorm.ErrRecordNotFound) {
		util.ResponseKit(c, http.StatusOK, "该用户还未申请场馆入住", gin.H{"status": "NULL"})
		return
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"status": audit.Status})

}
func SaasWxGetShopDetail(c *gin.Context) {
	param := param.UserOuidParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "用户ouid不能为空", nil)
		return
	}
	audit := models.ShopAudit{}
	if errors.Is(db.Where("user_ouid = ?", param.UserOuid).First(&audit).Error, gorm.ErrRecordNotFound) {
		util.ResponseKit(c, http.StatusOK, "该用户还未申请场馆入住", gin.H{"status": "NULL"})
		return
	}
	if audit.Status != dao.PayStatusYes {
		util.ResponseKit(c, http.StatusInternalServerError, "还未申请成功", gin.H{"status": audit.Status})
		return
	}
	shop := models.Shop{}
	db.Where("id = ?", audit.ShopId).First(&shop)
	var photos []string
	json.Unmarshal([]byte(shop.Photo), &photos)
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"site_count":    shop.SiteCount,
		"week_money":    20000,
		"month_money":   5000000,
		"day_reserve":   10,
		"monty_reserve": 56,
		"avatar":        shop.Avatar,
		"name":          shop.Name,
		"photo":         photos,
		"longitude":     shop.Longitude,
		"latitude":      shop.Latitude,
		"address":       shop.Address,
		"phone":         shop.Phone,
		"work_time":     shop.WorkTime,
		"desc":          shop.Desc,
		"shop_id":       shop.Id,
	})
}
func SaasWxUpdateShopDetail(c *gin.Context) {
	param := param.SaasWxUpdateShopDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	marshal, _ := json.Marshal(&param.Photo)
	db.Table("shop").Where("id = ?", param.ShopId).Updates(map[string]interface{}{
		"avatar":    param.Avatar,
		"name":      param.Name,
		"photo":     string(marshal),
		"longitude": param.Longitude,
		"latitude":  param.Latitude,
		"address":   param.Address,
		"phone":     param.Phone,
		"work_time": param.WorkTime,
		"desc":      param.Desc,
	})
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)

}
func SaasWxGetOrderList(c *gin.Context) {
	param := param.SaasWxGetOrderListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}

	begin := db
	type SaasWxGetOrderListVO struct {
		OrderNo   string   `json:"order_no"`
		OrderName string   `json:"order_name"`
		SiteList  []string `json:"site_list"`
		Time      string   `json:"time"`
		Money     int      `json:"money"`
		Status    string   `json:"status"`
	}
	var orderList []models.Order
	begin.Where("shop_id = ?", param.ShopId).Order("gmt_create DESC").Limit(20).Find(&orderList)
	var VOList []SaasWxGetOrderListVO
	for _, v := range orderList {
		var siteDetail []models.SiteOrderDetail
		var siteNameList []string
		json.Unmarshal([]byte(v.SiteDetail), &siteDetail)
		for _, vv := range siteDetail {
			siteNameList = append(siteNameList, vv.SiteName)
		}
		VOList = append(VOList, SaasWxGetOrderListVO{
			OrderNo:   v.OrderNo,
			OrderName: v.OrderName,
			SiteList:  siteNameList,
			Time:      v.GmtCreate.Format("2006-01-02"),
			Money:     v.Money,
			Status:    v.Status,
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", VOList)
}
func SaasWxGetOrderDetail(c *gin.Context) {
	param := param.SaasWxGetOrderDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	order := models.Order{}
	db.Where("order_no = ?", param.OrderNO).First(&order)
	type SaasWxGetOrderDetailVO struct {
		OrderName   string                   `json:"order_name"`
		Money       int                      `json:"money"`
		Status      string                   `json:"status"`
		GmtCreate   string                   `json:"gmt_create"`
		GmtSuccess  string                   `json:"gmt_success"`
		GmtRefund   string                   `json:"gmt_refund"`
		UserName    string                   `json:"user_name"`
		UserPhone   string                   `json:"user_phone"`
		SiteDetail  []models.SiteOrderDetail `json:"site_detail"`
		UseDate     string                   `json:"use_date"`
		ReserveDate string                   `json:"reserve_date"` //预约日期

	}
	var siteDetail []models.SiteOrderDetail
	json.Unmarshal([]byte(order.SiteDetail), &siteDetail)
	use, _ := dao.SiteUseFirst(order.OrderNo)
	GmtPay := ""
	if !order.GmtSuccess.IsZero() {
		GmtPay = order.GmtSuccess.Format("2006-01-02 15:04:05")
	}

	util.ResponseKit(c, http.StatusOK, "请球成功", SaasWxGetOrderDetailVO{
		OrderName:   order.OrderName,
		Money:       order.Money,
		Status:      order.Status,
		GmtCreate:   order.GmtCreate.Format("2006-01-02"),
		GmtSuccess:  GmtPay,
		GmtRefund:   order.GmtRefund.Format("2006-01-02"),
		UserName:    order.UserName,
		UserPhone:   order.UserPhone,
		SiteDetail:  siteDetail,
		UseDate:     use.UseDate[:10],
		ReserveDate: order.ReserveDate.Format("2006-01-02"),
	})
}
func SaasWxGetEarnings(c *gin.Context) {
	param := param.SaasWxGetEarningsParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	var monthMoney int
	db.Table("order").Where("shop_id = ? AND status = 'Y' AND gmt_create >?", param.ShopId, util.GetFirstDateOfMonth(util.Now())).Select("IFNULL(SUM(money), 0)").Scan(&monthMoney)
	var weekMoney int
	db.Table("order").Where("shop_id = ? AND status = 'Y' AND gmt_create >?", param.ShopId, util.GetFirstDateOfWeek(util.Now())).Select("IFNULL(SUM(money), 0)").Scan(&weekMoney)

	var orderList []models.Order
	db.Where("shop_id = ?", param.ShopId).Or("status = 'Y' AND status = 'R'").Limit(15).Find(&orderList)
	type OrderListVO struct {
		OrderNo   string `json:"order_no"`
		OrderName string `json:"order_name"`
		GmtCreate string `json:"gmt_create"`
		Money     int    `json:"money"`
	}
	var VOList []OrderListVO
	for _, v := range orderList {
		vo := OrderListVO{
			OrderNo:   v.OrderNo,
			OrderName: v.OrderNo,
			GmtCreate: v.GmtCreate.Format("2006-01-02"),
		}
		if v.Status == dao.PayStatusYes {
			vo.Money = v.Money
		} else {
			vo.Money = -v.Money
		}
		VOList = append(VOList, vo)
	}
	var weekEarnings = gin.H{
		"x_data": []string{"1", "2", "3", "4", "5", "6", "7"},
		"y_data": []string{"20", "30", "40", "20", "60", "80", "50"},
	}
	util.ResponseKit(c, http.StatusOK, "查询成功", gin.H{
		"month_money":   monthMoney,
		"week_money":    weekMoney,
		"week_earnings": weekEarnings,
		"order_list":    VOList,
	})
}
func SaasWxGetActivityList(c *gin.Context) {
	param := param.SaasWxGetActivityListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	var activityList []models.Activity
	db.Where("shop_id = ?", param.ShopId).Order("gmt_create DESC").Find(&activityList)
	type SaasWxGetActivityListVO struct {
		ActivityId int    `json:"activity_id"`
		Name       string `json:"name"`
		UserName   string `json:"user_name"`
		Time       string `json:"time"`
		TotalCount int    `json:"total_count"`
		ApplyCount int    `json:"apply_count"`
		Avatar     string `json:"avatar"`
		Status     string `json:"status"`
		Address    string `json:"address"`
	}
	var VOList []SaasWxGetActivityListVO
	for _, v := range activityList {
		vo := SaasWxGetActivityListVO{
			ActivityId: v.Id,
			Name:       v.Name,
			UserName:   v.CreateUserName,
			Time:       v.ActivityDate.Format("2006-01-02") + " " + v.ActivityTime,
			TotalCount: v.TotalCount,
			ApplyCount: v.ApplyCount,
			Avatar:     v.Avatar,
			Address:    v.ShopAddress,
		}

		if util.Now().After(v.ActivityDate) {
			vo.Status = dao.PayStatusNo
		} else {
			vo.Status = dao.PayStatusYes
		}
		VOList = append(VOList, vo)
	}
	util.ResponseKit(c, http.StatusOK, "查询成功", VOList)
}
func SaasWxGetActivityDetail(c *gin.Context) {
	param := param.SaasWxGetActivityDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	activity := models.Activity{}
	db.Where("id = ?", param.ActivityId).First(&activity)

	var applyList []models.ActivityApply
	db.Where("activity_id = ? AND status !='C'", param.ActivityId).Limit(6).First(&applyList)
	type applyListVO struct {
		Avatar string `json:"avatar"`
		Name   string `json:"name"`
		Sex    int    `json:"sex"`
	}
	var applyListVo []applyListVO
	for _, v := range applyList {
		applyListVo = append(applyListVo, applyListVO{
			Avatar: v.UserAvatar,
			Name:   v.UserName,
			Sex:    v.UserSex,
		})
	}

	util.ResponseKit(c, http.StatusOK, "查询成功", gin.H{
		"activity_detail": activity,
		"price": gin.H{
			"price_man":   activity.PriceMan,
			"price_woman": activity.PriceWoman,
		},
		"apply_list": applyListVo,
	},
	)
}
func SaasWxGetSiteList(c *gin.Context) {
	param := param.SaasWxGetSiteListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	// 获取所有场地
	siteAll, err := dao.SiteAll(param.ShopId)
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
		SiteId          int         `json:"site_id"`
		SiteName        string      `json:"site_name"`
		ReserveTimeEnum []int       `json:"reserve_time_enum"`
		StoreTimeEnum   []int       `json:"store_time_enum"`
		Price           []SitePrice `json:"price"`
	}

	var siteVO []Site

	for _, v := range siteAll {
		site := Site{
			SiteId:          v.Id,
			SiteName:        v.Name,
			ReserveTimeEnum: []int{},
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
				Price:    SitePriceCalculate(v, util.Now(), vv),
			})
		}
		site.Price = sitePrice
		siteVO = append(siteVO, site)

	}

	siteUseList, _ := dao.SiteUseListByShopId(param.ShopId, util.Now())
	for _, v := range siteUseList {
		var a []int
		json.Unmarshal([]byte(v.UseTimeEnum), &a)
		for kk, vv := range siteVO {
			if v.SiteId == vv.SiteId {
				siteVO[kk].ReserveTimeEnum = append(siteVO[kk].ReserveTimeEnum, a...)
			}
		}
	}

	type siteListVO struct {
		SiteName   string `json:"site_name"`
		Status     string `json:"status"`      //状态
		FreePrice  int    `json:"free_price"`  //闲时价格
		BusyPrice  int    `json:"busy_price"`  //忙时价格
		TodayCount int64  `json:"today_count"` //今日预约次数
	}
	var VOList []siteListVO
	for _, v := range siteAll {
		var count int64
		db.Where("site_id = ? AND use_date = ?", v.Id, util.Now().Format("2006-01-02")).Count(&count)
		VOList = append(VOList, siteListVO{
			SiteName:   v.Name,
			Status:     v.Status,
			FreePrice:  v.FreePrice,
			BusyPrice:  v.BusyPrice,
			TodayCount: count,
		})
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"site_detail": siteVO,
		"site_list":   VOList,
	})

}
func SaasWxGetSiteReserve(c *gin.Context) {
	param := param.SaasWxGetSiteReserveParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	siteUse, _ := dao.SiteUseListBySiteIdAll(param.Date, param.SiteId)
	type SaasWxGetSiteReserveVO struct {
		OrderNo   string `json:"order_no"`
		TimeEnum  []int  `json:"time"`
		UserName  string `json:"user_name"`
		UserPhone string `json:"user_phone"`
		GmtCreate string `json:"gmt_create"`
	}
	var VOList []SaasWxGetSiteReserveVO
	for _, v := range siteUse {
		var timeEnum []int
		json.Unmarshal([]byte(v.UseTimeEnum), &timeEnum)
		VOList = append(VOList, SaasWxGetSiteReserveVO{
			OrderNo:   v.OrderNo,
			TimeEnum:  timeEnum,
			UserName:  v.UserName,
			UserPhone: v.UserPhone,
			GmtCreate: v.GmtCreate.Format("22006-01-02 15:04"),
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", VOList)
}
