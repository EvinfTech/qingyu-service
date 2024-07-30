package dao

import (
	"gorm.io/gorm"
	"qiyu/models"
	"qiyu/util"
	"time"
)

const (
	SiteUseStatusY = "Y"
	SiteUseStatusN = "N"
)

// SiteUseFirst 一个订单的一个场地占用
func SiteUseFirst(orderNo string) (siteUse *models.SiteUse, err error) {
	err = DB.Where("order_no = ?", orderNo).
		First(&siteUse).Error
	return
}

// SiteUseAll 一个订单的所有场地占用
func SiteUseAll(orderNo string) (siteUseAll []*models.SiteUse, err error) {
	err = DB.Where("order_no = ?", orderNo).
		Find(&siteUseAll).Error
	return
}

// SiteUseList 根据占用日期 获取场地占用列表
func SiteUseList(useDate string) (siteUseList []*models.SiteUse, err error) {
	err = DB.Where("use_date = ?", useDate).
		Where("status = 'Y'").
		Find(&siteUseList).Error
	return
}

// SiteUseListBySiteId 根据占用日期、场地ID 获取场地占用列表
func SiteUseListBySiteId(useDate time.Time, siteId int) (siteUseList []*models.SiteUse, err error) {
	err = DB.Where("use_date = ?", util.Date2String(useDate)).
		Where("site_id = ?", siteId).
		Where("status = 'Y'").
		Find(&siteUseList).Error
	return
}

// SiteUseListBySiteIdAll 根据占用日期、场地ID 获取场地占用列表，不限状态
func SiteUseListBySiteIdAll(useDate string, siteId int) (siteUseList []*models.SiteUse, err error) {
	err = DB.Where("use_date = ?", useDate).
		Where("site_id = ?", siteId).
		Find(&siteUseList).Error
	return
}

// SiteUseListBySiteIdSlice 根据占用日期和场地ID切片 获取场地占用列表
func SiteUseListBySiteIdSlice(useDate string, siteIdSlice []int) (siteUseList []*models.SiteUse, err error) {
	err = DB.Where("use_date = ?", useDate).
		Where("site_id in ?", siteIdSlice).
		Where("status = 'Y'").
		Find(&siteUseList).Error
	return
}

// SiteUseListByShopId 根据商家ID、占用日期 获取场地占用列表
func SiteUseListByShopId(shopId int, useDate time.Time) (siteUseList []*models.SiteUse, err error) {
	err = DB.Where("shop_id = ?", shopId).
		Where("use_date = ?", util.Date2String(useDate)).
		Where("status = 'Y'").
		Find(&siteUseList).Error
	return
}

// SiteUseOrderNoSlice 时间内正在占用的场地的订单列表
func SiteUseOrderNoSlice(startTime, endTime string) []string {
	var orderNoList []models.SiteUse
	DB.Model(&models.SiteUse{}).Select("DISTINCT order_no").
		Where("use_date >= ?", startTime).
		Where("use_date <= ?", endTime).
		Where("status = 'Y'").
		Find(&orderNoList)
	orderNoSlice := make([]string, 0, len(orderNoList))
	for _, one := range orderNoList {
		orderNoSlice = append(orderNoSlice, one.OrderNo)
	}
	return orderNoSlice
}

// SiteUseTodayOrderCount 今日占用场地对应的订单数量
func SiteUseTodayOrderCount() int64 {
	var todayOrderQ struct {
		Count int64
	}
	DB.Model(&models.SiteUse{}).Select("Count(DISTINCT order_no) as count").
		Where("use_date = ?", util.NewTime().GetYmd()).
		Where("status = 'Y'").
		Find(&todayOrderQ)
	return todayOrderQ.Count
}

// SiteUseAdd 新增场地占用
func SiteUseAdd(tx *gorm.DB, user *models.User, siteId int, gmtSiteUse time.Time, useTimeEnum []byte, orderNo string, orderType int) error {
	now := util.Now()
	return tx.Create(&models.SiteUse{
		SiteId:       siteId,
		UseDate:      util.Date2String(gmtSiteUse),
		UseTimeEnum:  string(useTimeEnum),
		UserOuid:     user.Ouid,
		UserName:     user.Name,
		UserAvatar:   user.Avatar,
		UserPhone:    user.Phone,
		OrderNo:      orderNo,
		Type:         orderType,
		OperatorOuid: user.Ouid,
		Status:       PayStatusYes,
		GmtCreate:    now,
		GmtUpdate:    now,
	}).Error
}

// SiteUseUpdateNot 场地占用状态批量改为取消占用
func SiteUseUpdateNot(orderNo string) error {
	return DB.Model(&models.SiteUse{}).
		Where("order_no = ?", orderNo).
		Updates(map[string]any{
			"status":     "N",
			"gmt_update": util.Now(),
		}).Error
}
