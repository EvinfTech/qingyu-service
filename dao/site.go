package dao

import (
	"qiyu/models"
	"qiyu/models/param"
	"qiyu/util"
)

// SiteAll 所有场地
func SiteAll(shopId int) (siteAll []models.Site, err error) {
	err = DB.Where("shop_id = ? AND del = 0", shopId).Find(&siteAll).Error
	return
}

// SiteAllShop 所有商户的所有场地
func SiteAllShop() (siteAll []models.Site, err error) {
	err = DB.Where("del = 0").Find(&siteAll).Error
	return
}

// SiteCount 所有商户的所有场地数量
func SiteCount() (count int64, err error) {
	err = DB.Model(&models.Site{}).Where("del = 0").Count(&count).Error
	return
}

// SiteGet 获取一个场地
func SiteGet(siteId int) (siteOne *models.Site, err error) {
	err = DB.Where("id = ?", siteId).
		Where("del = 0").
		First(&siteOne).Error
	return
}

// SiteAdd 新增一个场地
func SiteAdd(param param.SaasAddSiteParam, shopId int, timeEnum, dataBusy, busyTimeEnum []byte) error {
	now := util.Now()
	siteOne := models.Site{
		ShopId:      shopId,
		Name:        param.Name,
		TimeEnum:    string(timeEnum),
		GmtCreate:   now,
		GmtUpdate:   now,
		Status:      param.Status,
		FreePrice:   param.FreePrice,
		BusyPrice:   param.BusyPrice,
		WeekendBusy: param.WeekendBusy,
		HolidayBusy: param.HolidayBusy,
	}
	if dataBusy != nil {
		siteOne.DataBusy = string(dataBusy)
	}
	if busyTimeEnum != nil {
		siteOne.BusyTimeEnum = string(busyTimeEnum)
	}
	return DB.Create(&siteOne).Error
}

// SiteUpdate 更新一个场地
func SiteUpdate(siteOne *models.Site, param param.SaasUpdateSiteParam, timeEnum, busyTimeEnum []byte) error {
	m := map[string]interface{}{
		"time_enum":    string(timeEnum),
		"name":         param.Name,
		"status":       param.Status,
		"free_price":   param.FreePrice,
		"busy_price":   param.BusyPrice,
		"holiday_busy": param.HolidayBusy,
		"weekend_busy": param.WeekendBusy,
		"gmt_update":   util.Now(),
	}
	if busyTimeEnum != nil {
		m["busy_time_enum"] = string(busyTimeEnum)
	}
	return DB.Model(siteOne).Updates(m).Error
}

// SiteDelete 删除一个场地
func SiteDelete(siteId int) error {
	return DB.Model(&models.Site{}).
		Where("id = ?", siteId).
		Updates(map[string]any{
			"del":        1,
			"gmt_update": util.Now(),
		}).Error
}
