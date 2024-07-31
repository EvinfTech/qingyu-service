package main

import (
	"qiyu/conf"
	"qiyu/controller"
	"qiyu/dao"
	"qiyu/kit"
	logKit "qiyu/logger"
	"qiyu/util"
)

func main() {
	// 配置文件初始化
	conf.Init()
	//初始化数据库
	dao.Init()
	//初始化日志模块
	logKit.Init()
	//文件夹初始化
	util.FileInit()

	//自动执行订单结算
	go kit.AutoSuperviseMoney()
	////初始化微信token
	go kit.AccessTokenInit()
	////微信初始化
	go kit.WxPayInit()

	//路由初始化
	controller.SetRouters()

}
