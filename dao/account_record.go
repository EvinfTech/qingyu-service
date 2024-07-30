package dao

import (
	"gorm.io/gorm"
	"qiyu/models"
	"qiyu/models/enum"
)

// AccountRecordGetPay 获取一条储值余额支付记录
func AccountRecordGetPay(orderNo string) (*models.AccountRecord, error) {
	var record models.AccountRecord
	err := DB.Where("related_order_no = ?", orderNo).
		Where("change_type = ?", enum.AccountChangeTypePay).
		First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func AccountRecordRecharge(tx *gorm.DB, user *models.User, changeType enum.AccountChangeType, actualMoney, presentMoney uint32, creatorId uint32) error {
	return tx.Create(&models.AccountRecord{
		Ouid:         user.Ouid,
		ChangeType:   changeType,
		ActualMoney:  actualMoney,
		PresentMoney: presentMoney,
		NewActual:    user.ActualBalance + actualMoney,
		NewPresent:   user.PresentBalance + presentMoney,
		CreatorID:    creatorId,
	}).Error
}

func AccountRecordPay(tx *gorm.DB, user *models.User, money uint32, relatedOrderNo string, creatorId uint32) error {
	var actualMoney, presentMoney, newActualMoney, newPresentMoney uint32
	if user.ActualBalance >= money {
		actualMoney = money
		presentMoney = 0
		newActualMoney = user.ActualBalance - money
		newPresentMoney = user.PresentBalance
	} else {
		actualMoney = user.ActualBalance
		presentMoney = money - user.ActualBalance
		newActualMoney = 0
		newPresentMoney = user.PresentBalance - presentMoney
	}
	return tx.Create(&models.AccountRecord{
		Ouid:           user.Ouid,
		ChangeType:     enum.AccountChangeTypePay,
		ActualMoney:    actualMoney,
		PresentMoney:   presentMoney,
		NewActual:      newActualMoney,
		NewPresent:     newPresentMoney,
		RelatedOrderNo: relatedOrderNo,
		CreatorID:      creatorId,
	}).Error
}

func AccountRecordPayRefund(tx *gorm.DB, user *models.User, actualMoney, presentMoney uint32, relatedOrderNo string, creatorId uint32) error {
	return tx.Create(&models.AccountRecord{
		Ouid:           user.Ouid,
		ChangeType:     enum.AccountChangeTypePayRefund,
		ActualMoney:    actualMoney,
		PresentMoney:   presentMoney,
		NewActual:      user.ActualBalance + actualMoney,
		NewPresent:     user.PresentBalance + presentMoney,
		RelatedOrderNo: relatedOrderNo,
		CreatorID:      creatorId,
	}).Error
}
