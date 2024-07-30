package dao

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"gorm.io/gorm"
	"qiyu/models"
	"qiyu/util"
)

const (
	PayStatusYes = "Y"
	PayStatusNo  = "N"

	PayTypeWx       = "微信支付"
	PayTypeRefund   = "退款"
	PayTypeRecharge = "充值"
)

func PayAdd(tx *gorm.DB, order *models.Order, user *models.User) error {
	now := util.Now()
	pay := models.Pay{
		OrderNo:         order.OrderNo,
		TotalMoney:      order.OriginalPrice,
		Money:           order.Money,
		Status:          PayStatusNo,
		ResponseContext: "",
		PayForm:         user.Name,
		PayTo:           "跃迁科技",
		UserOuid:        user.Ouid,
		UserName:        user.Name,
		UserAvatar:      user.Avatar,
		Remark:          "",
		GmtCreate:       now,
		GmtUpdate:       now,
	}
	return tx.Create(&pay).Error
}

func PayUpdate(tx *gorm.DB, orderNo string, responseContext []byte) error {
	now := util.Now()
	return tx.Model(&models.Pay{}).Where("order_no = ?", orderNo).Updates(models.Pay{
		PayType:         PayTypeWx,
		Status:          PayStatusYes,
		ResponseContext: string(responseContext),
		Remark:          "支付成功",
		GmtResponse:     now,
		GmtUpdate:       now,
	}).Error
}

func PayRefund(tx *gorm.DB, order *models.Order, resp *refunddomestic.Refund) error {
	now := util.Now()
	pay := models.Pay{
		OrderNo:         order.OrderNo,
		PayType:         PayTypeRefund,
		TotalMoney:      order.OriginalPrice,
		Money:           order.Money,
		Status:          "Y",
		ResponseContext: util.ToString(resp),
		PayForm:         "跃迁科技",
		PayTo:           order.UserName,
		UserOuid:        order.UserOuid,
		UserName:        order.UserName,
		UserAvatar:      order.UserAvatar,
		Remark:          "退款订单",
		GmtResponse:     now,
		GmtCreate:       now,
		GmtUpdate:       now,
	}
	return tx.Create(&pay).Error
}

// PayAddRecharge 新增充值
func PayAddRecharge(tx *gorm.DB, user *models.User, orderNo string, money int) error {
	now := util.Now()
	pay := models.Pay{
		OrderNo:         orderNo,
		PayType:         PayTypeRecharge,
		TotalMoney:      money,
		Money:           money,
		Status:          PayStatusNo,
		ResponseContext: "",
		PayForm:         user.Name,
		PayTo:           "跃迁科技",
		UserOuid:        user.Ouid,
		UserName:        user.Name,
		UserAvatar:      user.Avatar,
		Remark:          "",
		GmtCreate:       now,
		GmtUpdate:       now,
	}
	return tx.Create(&pay).Error
}

// PayGetRecharge 获取一个充值订单
func PayGetRecharge(orderNo string) (*models.Pay, error) {
	var pay models.Pay
	err := DB.Where("order_no = ?", orderNo).
		Where("pay_type = ?", PayTypeRecharge).
		First(&pay).Error
	if err != nil {
		return nil, err
	}
	return &pay, nil
}

func PayUpdateRecharge(tx *gorm.DB, payOne *models.Pay, responseContext []byte) error {
	now := util.Now()
	return tx.Model(&payOne).Updates(models.Pay{
		Status:          PayStatusYes,
		ResponseContext: string(responseContext),
		Remark:          "支付成功",
		GmtResponse:     now,
		GmtUpdate:       now,
	}).Error
}
