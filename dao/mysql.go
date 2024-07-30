package dao

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	logKit "qiyu/logger"
	"qiyu/models"
	"qiyu/util"
	"time"
)

// DB 本方法内不能使用util工具
var DB *gorm.DB

// Init 初始化入口
func Init() {

}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	// log.Infof(format, args...)
	fmt.Printf(format, args...)
	fmt.Println()
	log.Println("====数据库查询=====", args[1:])
}
func init() {
	var err error
	newLogger := logger.New(
		//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		Writer{},
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	dsn := "root:woshimima@tcp(127.0.0.1:3306)/qiyu_open?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	//解决苦逼的mysql高版本groupby问题
	DB.Raw("SET GLOBAL sql_mode='STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION';").Row()
	db, err := DB.DB()
	//最大链接时长
	if err != nil {
		//todo  这里数据库连接出错，需要自行进行优化
		logKit.Log.Println("数据库连接发生错误:" + err.Error())

	}
	db.SetConnMaxLifetime(1 * time.Hour)

	shop := models.Shop{}
	if errors.Is(DB.First(&shop).Error, gorm.ErrRecordNotFound) {
		DB.Create(&models.Shop{
			Name:            "门店001",
			Avatar:          "",
			Photo:           "",
			Address:         "店铺默认地址",
			Longitude:       120.192519,
			Latitude:        35.975585,
			Tag:             "",
			Phone:           "",
			WorkTime:        "",
			SiteCount:       0,
			BottomPrice:     0,
			Desc:            "",
			Facility:        "",
			Serve:           "",
			AggregateAmount: 0,
			Balance:         0,
			GmtCreate:       util.Now(),
			GmtUpdate:       util.Now(),
		})
	}
	admin := models.AdminUser{}
	if errors.Is(DB.First(&admin).Error, gorm.ErrRecordNotFound) {
		DB.Create(&models.AdminUser{
			UserOuid:  util.GenerateOuid(),
			Password:  "admin",
			Avatar:    "avatar/171385753493.png",
			Name:      "admin",
			RoleId:    1,
			RoleName:  "管理员",
			ShopId:    0,
			Level:     0,
			GmtCreate: util.Now(),
			GmtUpdate: util.Now(),
		})
	}

	logKit.Log.Println("数据库链接成功")
}
