package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiyu/dao"
	"qiyu/util"
)

var db = dao.DB

func SetRouters() {
	r := gin.Default()
	//开启中间件 允许使用跨域请求
	r.Use(util.Cors(), cors.Default())
	r.StaticFS("/avatar", http.Dir("static/avatar"))
	r.StaticFS("/photo", http.Dir("static/photo"))
	r.StaticFS("/web", http.Dir("static/h5"))
	r.StaticFS("/qr", http.Dir("static/qr"))
	r.StaticFS("/saas", http.Dir("static/saas"))
	r.GET("/", func(c *gin.Context) {
		// 指定重定向路由
		c.Request.URL.Path = "/web"

		// 继续后续处理
		r.HandleContext(c)
	})
	{
		common := r.Group("common")
		//获取枚举
		common.POST("get/enum", CommonGetEnum)
		//获取验证码
		common.POST("get-phone-code", GetPhoneCode)
		//验证码登录
		common.POST("phone-code-login", PhoneCodeLogin)
		//上传头像
		common.POST("upload/avatar", UploadAvatar)
		//上传图片
		common.POST("upload/photo", UploadPhoto)
	}
	//微信小程序接口
	{
		//微信接口
		wx := r.Group("wx")
		//微信登录接口
		wx.POST("login", WxLogin)
		//获取用户信息
		wx.POST("get/user/info", WxGetUserInfo)
		//获取用户openId和unionId
		wx.POST("get/wx/id", WxGetWxId)
		//获取手机号
		wx.POST("get/phone", WxGetPhone)
		//更新用户信息
		wx.POST("update/user/info", WxUpdateUserInfo)
		// 充值规则列表
		wx.POST("recharge-rule/list", WxRechargeRuleList)
		// 新增微信充值订单
		wx.POST("recharge-order/add", WxRechargeOrderAdd)
		// 查询充值订单状态
		wx.POST("recharge-order/status", WxRechargeOrderStatus)
		//查看球馆详情
		wx.POST("get/shop/detail", WxGetShopDetail)
		//球馆详情下面的订场概括
		wx.POST("get/shop/surplus/count", WxGetShopSurplusCount)
		//小程序定场
		wx.POST("add/order", WxAddOrder)
		//微信支付
		wx.POST("pay", WxPay)
		//微信支付回调接口
		wx.POST("pay/notify", WxPayNotify)
		// 储值余额支付
		wx.POST("balance/pay", WxBalancePay)
		//查看某场馆的7天内场地预定信息
		wx.POST("get/site/reserve", WxGetSiteReserve)
		//用户查看订单
		wx.POST("get/order/list", WxGetOrderList)
		//查看订单详情
		wx.POST("get/order/detail", WxGetOrderDetail)
		//查看订单状态
		wx.POST("get/order/status", WxGetOrderStatus)
		//查看我的预约
		wx.POST("get/my/reserve/list", WxGetMyReserveList)
		//查看预约详情
		wx.POST("get/reserve/detail", WxGetReserveDetail)
		//填写反馈
		wx.POST("add/feedback", WxAddFeedback)
		//取消订单
		wx.POST("cancel/order", WxCancelOrder)
		//最近预约
		wx.POST("recently/reserve", WxRecentlyReserve)
		//场馆报错信息
		wx.POST("add/shop/error", WxAddShopError)
		//微信查看协议和关于我们
		wx.POST("get/agreement/about", WxGetAgreementAbout)
	}
	//后台接口
	{
		//saas接口
		saas := r.Group("saas")
		//后台核验订单
		saas.POST("check/order", SaasCheckOrder)
		//首页运营数据
		saas.POST("get/operational/data", SaasGetOperationalData)
		//获取saas资金流水页面
		saas.POST("get/shop/capital", SaasGetShopCapital)
		//提现
		saas.POST("/withdraw/deposit", SaasWithdrawDeposit)
		//查看账单流水
		saas.POST("get/bill/list", SaasGetBillList)
		//添加门店
		saas.POST("add/shop", SaasAddShop)
		//获取门店详情
		saas.POST("get/shop/detail", SaasGetShopDetail)
		//更新门店信息
		saas.POST("update/shop/detail", SaasUpdateShopDetail)
		//添加场地表
		saas.POST("add/site", SaasAddSite)
		//获取场地列表
		saas.POST("get/site/list", SaasGetSiteList)
		//更新场地信息
		saas.POST("update/site", SaasUpdateSite)
		//取消订单
		saas.POST("cancel/order", SaasCancelOrder)
		// 后台添加订单
		saas.POST("add/order", SaasAddOrder)
		//查看门店订单列表
		saas.POST("get/order/list", SaasGetOrderList)
		//修改备注
		saas.POST("update/remark", SaasUpdateRemark)
		//删除场地
		saas.POST("del/site", SaasDelSite)
		//查看反馈列表
		saas.POST("get/feedback/list", GetFeedbackList)
		//------------------------------------------------------
		//增加角色
		//saas.POST("add/role", SaasAddRole)
		////查看所有角色
		//saas.POST("get/role/list", SaasGetRoleList)
		////删除角色
		//saas.POST("del/role", SaasDelRole)
		////更新角色
		//saas.POST("update/role", SaasUpdateRole)

		//添加管理员用户
		saas.POST("add/admin/user", SaasAddAdminUser)
		//查看所有用户列表
		saas.POST("get/admin/user/list", SaasGetAdminUserList)
		//查看用户名是否可用
		saas.POST("get/admin/userName/status", SaasGetAdminUserNameStatus)
		//更新信息
		saas.POST("update/admin/user", SaasUpdateAdminUser)
		//删除管理员
		saas.POST("del/admin/user", SaasDelAdminUser)
		////添加菜单
		//saas.POST("add/menu", SaasAddMenu)
		////查看菜单列表
		//saas.POST("get/menu/list", SaasGetMenuList)
		////更新菜单信息
		//saas.POST("update/menu", SaasUpdateMenu)
		////删除菜单
		//saas.POST("del/menu", SaasDelMenu)
		////增加角色菜单权限表
		//saas.POST("update/menuRole", SaasUpdateMenuRole)
		////获取角色菜单权限列表
		//saas.POST("get/menuRole/list", SaasGetMenuRoleList)
		////删除角色菜单权限列表
		//saas.POST("del/menu/role", SaasDelMenuRole)
		//登录接口
		saas.POST("login", SaasLogin)
		saas.POST("update-password", SaasUpdatePassword)
		//查看球馆的预约情况
		saas.POST("get/site/reserve", SaasGetSiteReserve)
		//查看会员列表
		saas.POST("get/user/list", SaasGetUserList)
		//修改会员信息
		saas.POST("update/user/info", SaasUpdateUserInfo)
		// 会员充值
		saas.POST("/user/recharge", UserRecharge)
		//saas更新协议和关于我们
		saas.POST("update/agreement/about", SaasUpdateAgreementAbout)
		//saas查看协议
		saas.POST("get/agreement/about", SaasGetAgreementAbout)

		{
			// 充值规则
			rule := saas.Group("recharge-rule")
			rule.POST("list", RechargeRuleList)
			rule.POST("add", RechargeRuleAdd)
			rule.POST("update", RechargeRuleUpdate)
			rule.POST("delete", RechargeRuleDelete)
		}
	}
	r.Run(":8002")
}
