package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"qiyu/dao"
	"qiyu/kit"
	logKit "qiyu/logger"
	"qiyu/logic"
	"qiyu/models/param"
	"qiyu/util"
)

// WxRechargeRuleList 微信--充值规则列表
func WxRechargeRuleList(c *gin.Context) {
	ruleAll, _, err := dao.RechargeRuleAll()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取充值规则列表错误："+err.Error(), nil)
		return
	}

	type RechargeRuleListVO struct {
		OriginMoney  uint32 `json:"origin_money"`
		PresentMoney uint32 `json:"present_money"`
	}
	list := make([]RechargeRuleListVO, 0, len(ruleAll))
	for _, rule := range ruleAll {
		list = append(list, RechargeRuleListVO{
			OriginMoney:  rule.OriginMoney,
			PresentMoney: rule.PresentMoney,
		})
	}
	util.ResponseKit(c, http.StatusOK, "查询成功", gin.H{"list": list})
}

// WxRechargeOrderAdd 新增微信充值订单
func WxRechargeOrderAdd(c *gin.Context) {
	var param param.WxRecharge
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}
	if param.Money == 0 {
		util.ResponseKit(c, http.StatusBadRequest, "充值金额不能为0", nil)
		return
	}

	// 检查用户是否存在
	user, err := dao.UserGetOne(param.Ouid)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取用户错误："+err.Error(), nil)
		return
	}

	orderNo := kit.GenerateOrderNo(PublicInt)
	PublicInt++

	var result gin.H
	err = db.Transaction(func(tx *gorm.DB) (errInner error) {
		if errInner = dao.PayAddRecharge(tx, user, orderNo, int(param.Money)); errInner != nil {
			return
		}
		result, errInner = wxPay(param.Type, user.WxOpenid, orderNo, int64(param.Money), c.ClientIP())
		return
	})

	if err != nil {
		logKit.Log.Println("预支付失败:", err.Error())
		util.ResponseKit(c, http.StatusInternalServerError, "预支付失败", nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", result)
}

// WxRechargeOrderStatus 查询充值订单状态
func WxRechargeOrderStatus(c *gin.Context) {
	var param param.WxGetOrderDetailParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}

	pay, err := dao.PayGetRecharge(param.OrderNo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusInternalServerError, "充值订单不存在", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查找充值订单错误:"+err.Error(), nil)
		}
		return
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"order_no": pay.OrderNo,
		"status":   pay.Status,
	})
}

// WxBalancePay 储值余额支付
func WxBalancePay(c *gin.Context) {
	var param param.BalancePay
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}

	// 获取该订单
	order, err := dao.OrderGet(param.OrderNo)
	if err != nil {
		return
	}
	if order.UserOuid != param.UserOuid {
		util.ResponseKit(c, http.StatusBadRequest, "不是该用户的订单", nil)
		return
	}
	if order.Status == "C" {
		util.ResponseKit(c, http.StatusBadRequest, "订单已取消，不可支付", nil)
		return
	}
	// 获取该用户
	user, err := dao.UserGetOne(order.UserOuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusInternalServerError, "会员不存在", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查找会员失败:"+err.Error(), nil)
		}
		return
	}
	mender, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// 用户金额相关
		if err = logic.UserMoneyPay(tx, user, order.Money, order.OrderNo, uint32(mender.Id)); err != nil {
			return err
		}
		// 修改订单表信息
		order.Status = dao.PayStatusYes
		order.GmtSuccess = util.Now()
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		// 修改支付表信息
		return dao.PayUpdate(tx, order.OrderNo, nil)
	})

	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "储值余额支付"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}
