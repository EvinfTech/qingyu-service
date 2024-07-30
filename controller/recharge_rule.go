package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"qiyu/dao"
	"qiyu/models/param"
	"qiyu/util"
)

// RechargeRuleList 充值规则列表
func RechargeRuleList(c *gin.Context) {
	ruleAll, total, err := dao.RechargeRuleAll()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取充值规则列表错误："+err.Error(), nil)
		return
	}

	type RechargeRuleListVO struct {
		Id           uint32 `json:"id"`
		OriginMoney  uint32 `json:"origin_money"`
		PresentMoney uint32 `json:"present_money"`
		MenderId     uint32 `json:"mender_id"`
		UpdateAt     string `json:"update_time"`
	}
	list := make([]RechargeRuleListVO, 0, len(ruleAll))
	for _, rule := range ruleAll {
		list = append(list, RechargeRuleListVO{
			Id:           rule.ID,
			OriginMoney:  rule.OriginMoney,
			PresentMoney: rule.PresentMoney,
			MenderId:     rule.MenderID,
			UpdateAt:     util.Time2String(rule.UpdateAt),
		})
	}
	util.ResponseKit(c, http.StatusOK, "查询成功", gin.H{"list": list, "total": total})
}

// RechargeRuleAdd 充值规则新增
func RechargeRuleAdd(c *gin.Context) {
	var param param.RechargeRuleAdd
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}
	if param.OriginMoney == 0 || param.PresentMoney == 0 {
		util.ResponseKit(c, http.StatusBadRequest, "金额不能为0", nil)
		return
	}

	creator, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}
	// 检查是否有原金额相同的规则
	if isExist, err := dao.RechargeRuleExistOrigin(param.OriginMoney); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "检查是否有原金额相同的规则错误："+err.Error(), nil)
		return
	} else if isExist {
		util.ResponseKit(c, http.StatusBadRequest, "已存在原金额相同的规则", nil)
		return
	}
	// 新增充值规则
	if err = dao.RechargeRuleAdd(param.OriginMoney, param.PresentMoney, uint32(creator.Id)); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "新增充值规则错误："+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// RechargeRuleUpdate 充值规则修改
func RechargeRuleUpdate(c *gin.Context) {
	var param param.RechargeRuleUpdate
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}
	if param.OriginMoney == 0 || param.PresentMoney == 0 {
		util.ResponseKit(c, http.StatusBadRequest, "金额不能为0", nil)
		return
	}

	mender, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}
	// 获取该充值规则
	theRule, err := dao.RechargeRuleGet(param.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusBadRequest, "充值规则不存在", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "获取充值规则错误："+err.Error(), nil)
		}
		return
	}
	// 检查是否有原金额相同、且不是该充值规则 的规则
	if isExist, err := dao.RechargeRuleExistOriginNotId(param.OriginMoney, theRule.ID); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "检查是否有原金额相同的规则错误："+err.Error(), nil)
		return
	} else if isExist {
		util.ResponseKit(c, http.StatusBadRequest, "已存在原金额相同的规则", nil)
		return
	}
	// 修改充值规则
	if err = dao.RechargeRuleUpdate(theRule, param.OriginMoney, param.PresentMoney, uint32(mender.Id)); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "修改充值规则错误："+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// RechargeRuleDelete 充值规则删除
func RechargeRuleDelete(c *gin.Context) {
	var param param.BaseDetail
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}

	mender, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}
	if err = dao.RechargeRuleDelete(param.Id, uint32(mender.Id)); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "删除充值规则错误："+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}
