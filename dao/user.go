package dao

import (
	"gorm.io/gorm"
	"qiyu/models"
	"qiyu/models/enum"
)

func UserGetOne(ouid string) (*models.User, error) {
	var user models.User
	if err := DB.Where("ouid = ?", ouid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UserRecharge(tx *gorm.DB, user *models.User, actualMoney, presentMoney uint32) error {
	newMap := map[string]any{
		"actual_balance":  gorm.Expr("actual_balance + ?", actualMoney),
		"present_balance": gorm.Expr("present_balance + ?", presentMoney),
	}
	if user.VipType == enum.VipTypeCommon {
		newMap["vip_type"] = enum.VipTypeRecharge
	}
	return tx.Model(user).Updates(newMap).Error
}

func UserMoneyPay(tx *gorm.DB, user *models.User, money int) error {
	var newMap map[string]any
	if int(user.ActualBalance) >= money {
		newMap = map[string]any{
			"actual_balance": gorm.Expr("actual_balance - ?", money),
		}
	} else {
		newMap = map[string]any{
			"actual_balance":  gorm.Expr("actual_balance - ?", user.ActualBalance),
			"present_balance": gorm.Expr("present_balance - ?", money-int(user.ActualBalance)),
		}
	}
	return tx.Model(user).Updates(newMap).Error
}

func UserMoneyRefund(tx *gorm.DB, user *models.User, actualMoney, presentMoney uint32) error {
	return tx.Model(user).Updates(map[string]any{
		"actual_balance":  gorm.Expr("actual_balance + ?", actualMoney),
		"present_balance": gorm.Expr("present_balance + ?", presentMoney),
	}).Error
}
