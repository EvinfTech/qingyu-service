package controller

import (
	"qiyu/dao"
	"qiyu/models"
	"qiyu/util"
	"sync"
	"time"
)

var TimeEnum map[int]string
var lock sync.Mutex
var PublicInt = 0

func init() {
	TimeEnum = map[int]string{
		1:  "00:00~01:00",
		2:  "01:00~02:00",
		3:  "02:00~03:00",
		4:  "03:00~04:00",
		5:  "04:00~05:00",
		6:  "05:00~06:00",
		7:  "06:00~07:00",
		8:  "07:00~08:00",
		9:  "08:00~09:00",
		10: "09:00~10:00",
		11: "10:00~11:00",
		12: "11:00~12:00",
		13: "12:00~13:00",
		14: "13:00~14:00",
		15: "14:00~15:00",
		16: "15:00~16:00",
		17: "16:00~17:00",
		18: "17:00~18:00",
		19: "18:00~19:00",
		20: "19:00~20:00",
		21: "20:00~21:00",
		22: "21:00~22:00",
		23: "22:00~23:00",
		24: "23:00~24:00",
	}
}

// CheckOrderByUserOuid 根据用户id进行判定，是否有已经超时的订单，进行关闭，注意，该订单要大于15分钟，避免和自动关闭脚本冲突
func CheckOrderByUserOuid(userOuid string) {
	go func() {
		//修改订单状态
		var orderList []models.Order
		//检测未支付订单进行关闭
		db.Table("order").Where("user_ouid = ? AND status = 'N' AND gmt_create < ?", userOuid, util.Now().Add(-time.Minute*16)).Find(&orderList)
		for _, v := range orderList {
			v.Status = "C"
			v.GmtUpdate = util.Now()
			db.Updates(&v)
			// 场地占用状态批量改为取消占用
			_ = dao.SiteUseUpdateNot(v.OrderNo)
		}
	}()
	//检测已支付订单是否已经使用
	go func() {
		var orderList []models.Order
		//检测未支付订单进行关闭
		db.Table("order").Where("user_ouid = ? AND status = 'Y' ", userOuid).Find(&orderList)
		for _, v := range orderList {
			siteTime, _ := GetOrderEndSiteTime(v.OrderNo)
			if siteTime.Before(util.Now()) {
				v.GmtUpdate = util.Now()
				v.Status = "U"
				db.Updates(&v)
			}
		}
	}()

}

// CheckOrderByShopId 根据店铺id进行判定，是否有已经超时的订单，进行关闭，注意，该订单要大于15分钟，避免和自动关闭脚本冲突
func CheckOrderByShopId(shopId int) {
	var orderList []models.Order
	db.Table("order").Where("shop_id = ? AND status = 'N' AND gmt_create < ?", shopId, util.Now().Add(-time.Minute*16)).Find(&orderList)
	for _, v := range orderList {
		v.Status = "C"
		v.GmtUpdate = util.Now()
		db.Updates(&v)
		// 场地占用状态批量改为取消占用
		_ = dao.SiteUseUpdateNot(v.OrderNo)
	}

}
