package dao

import (
	"errors"
	"gorm.io/gorm"
	"qiyu/models"
)

func RechargeRuleAll() (ruleAll []*models.RechargeRule, count int64, err error) {
	q := DB.Where("is_del = 0")
	if err = q.Model(&models.RechargeRule{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err = q.Order("origin_money desc").Find(&ruleAll).Error; err != nil {
		return nil, 0, err
	}
	return
}

func RechargeRuleGet(ruleId uint32) (rule *models.RechargeRule, err error) {
	err = DB.Where("id = ?", ruleId).
		Where("is_del = 0").
		First(&rule).Error
	if err != nil {
		return nil, err
	}
	return
}

func RechargeRuleExistOrigin(originMoney uint32) (bool, error) {
	var rule models.RechargeRule
	err := DB.Where("origin_money = ?", originMoney).
		Where("is_del = 0").
		First(&rule).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func RechargeRuleExistOriginNotId(originMoney uint32, ruleId uint32) (bool, error) {
	var rule models.RechargeRule
	err := DB.Where("origin_money = ?", originMoney).
		Not("id = ?", ruleId).
		Where("is_del = 0").
		First(&rule).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func RechargeRuleAdd(originMoney, presentMoney uint32, creatorId uint32) error {
	rule := models.RechargeRule{
		CreatorID:    creatorId,
		OriginMoney:  originMoney,
		PresentMoney: presentMoney,
		MenderID:     creatorId,
	}
	return DB.Create(&rule).Error
}

func RechargeRuleUpdate(rule *models.RechargeRule, originMoney, presentMoney uint32, menderID uint32) error {
	rule.OriginMoney = originMoney
	rule.PresentMoney = presentMoney
	rule.MenderID = menderID
	return DB.Save(&rule).Error
}

func RechargeRuleDelete(ruleId, menderId uint32) error {
	return DB.Model(&models.RechargeRule{}).
		Where("id = ?", ruleId).
		Updates(map[string]any{
			"mender_id": menderId,
			"is_del":    1,
		}).Error
}
