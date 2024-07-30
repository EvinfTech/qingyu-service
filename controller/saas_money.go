package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"qiyu/dao"
	"qiyu/models/enum"
	"qiyu/models/param"
	"qiyu/util"
)

// UserRecharge 账户充值
func UserRecharge(c *gin.Context) {
	var param param.Recharge
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
	mender, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}
	// 获取所有充值规则，从大到小检查是否负责赠送条件
	presentMoney, err := RechargeRuleGetPresent(param.Money)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "查询赠送金额错误："+err.Error(), nil)
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// 修改余额，如果是普通会员，改为充值会员
		if err2 := dao.UserRecharge(tx, user, param.Money, presentMoney); err2 != nil {
			return err2
		}
		// 新增管理员充值记录
		return dao.AccountRecordRecharge(tx, user, enum.AccountChangeTypeAdminRecharge, param.Money, presentMoney, uint32(mender.Id))
	})

	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "账户充值错误："+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "账户充值成功", nil)
	return
}

// RechargeRuleGetPresent 获取所有充值规则，从大到小检查是否负责赠送条件
func RechargeRuleGetPresent(origin uint32) (presentMoney uint32, err error) {
	ruleAll, _, err := dao.RechargeRuleAll()
	if err != nil {
		return
	}
	for _, rule := range ruleAll {
		rate := origin / rule.OriginMoney // 倍率
		if rate >= 1 {
			origin -= rule.OriginMoney * rate
			presentMoney += rule.PresentMoney * rate
			if origin == 0 { // 因为 rate>=1，所以 origin 不可能小于0（即使把 uint32 转为 int）
				break
			}
		}
	}
	return presentMoney, nil
}
