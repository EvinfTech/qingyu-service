package logic

import (
	"errors"
	"gorm.io/gorm"
	"qiyu/dao"
	"qiyu/models"
)

func UserMoneyPay(tx *gorm.DB, user *models.User, money int, orderNo string, menderId uint32) error {
	if money > int(user.ActualBalance+user.PresentBalance) {
		return errors.New("储值余额不足")
	}
	if err := dao.UserMoneyPay(tx, user, money); err != nil {
		return err
	}
	return dao.AccountRecordPay(tx, user, uint32(money), orderNo, menderId)
}

func UserMoneyRefund(tx *gorm.DB, user *models.User, orderNo string, menderId uint32) error {
	// 获取退款订单关联的储值余额支付记录
	record, err := dao.AccountRecordGetPay(orderNo)
	if err != nil {
		return err
	}
	if err = dao.UserMoneyRefund(tx, user, record.ActualMoney, record.PresentMoney); err != nil {
		return err
	}
	return dao.AccountRecordPayRefund(tx, user, record.ActualMoney, record.PresentMoney, orderNo, menderId)
}
