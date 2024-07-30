package util

import "gorm.io/gorm"

// PageKit 分页插件
func PageKit(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
