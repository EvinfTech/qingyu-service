package dao

import (
	"qiyu/models"
	"qiyu/util"
)

func AdminUserFirst() (*models.AdminUser, error) {
	var adminUser models.AdminUser
	if err := DB.First(&adminUser).Error; err != nil {
		return nil, err
	}
	return &adminUser, nil
}

func AdminUserGet(name, password string) (*models.AdminUser, error) {
	var adminUser models.AdminUser
	if err := DB.Where("name = ? AND password = ?", name, password).First(&adminUser).Error; err != nil {
		return nil, err
	}
	return &adminUser, nil
}

func AdminUserUpdatePassword(adminUser *models.AdminUser, newPassword string) error {
	adminUser.Password = newPassword
	adminUser.GmtUpdate = util.Now()
	return DB.Save(&adminUser).Error
}
