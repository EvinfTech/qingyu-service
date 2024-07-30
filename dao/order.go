package dao

import "qiyu/models"

const PayModeBalance = "储值支付"

// OrderGet 获取一个订单
func OrderGet(orderNo string) (order *models.Order, err error) {
	err = DB.Where("order_no = ?", orderNo).
		First(&order).Error
	return
}
