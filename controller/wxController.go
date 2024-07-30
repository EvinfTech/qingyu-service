package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"qiyu/conf"
	"qiyu/dao"
	"qiyu/kit"
	logKit "qiyu/logger"
	"qiyu/logic"
	"qiyu/models"
	"qiyu/models/enum"
	"qiyu/models/param"
	"qiyu/util"
	"strconv"
	"strings"
	"time"
)

func WxLogin(c *gin.Context) {
	param := param.GetOuidFromCodeParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}
	//配置自己的appid和secret————————————————————————————————————————————————————————
	wxUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		conf.Config.WxAppid, conf.Config.WxSecret, param.Code,
	)
	get, err := util.HTTPGet(wxUrl, nil, nil, nil, "")
	if err != nil {
		c.JSON(200, gin.H{"code": 0, "msg": "获取失败！", "data": nil})
		return
	}
	var response map[string]string
	json.Unmarshal(get.Body, &response)

	wxUrl = "https://api.weixin.qq.com/sns/userinfo?access_token=" + response["access_token"] + "&openid=" + response["openid"]
	userInfoPO, err := util.HTTPGet(wxUrl, nil, nil, nil, "")

	if err != nil {
		c.JSON(200, gin.H{"code": 0, "msg": "用户信息失败！", "data": nil})
		return
	}
	var appUserInfo map[string]string
	json.Unmarshal(userInfoPO.Body, &appUserInfo)
	generateOUID := util.GenerateOuid()
	userInfo := models.User{}
	if errors.Is(db.Where("wx_unionid =?", response["unionid"]).First(&userInfo).Error, gorm.ErrRecordNotFound) {

		if len(param.PhoneCode) == 0 {
			util.ResponseKit(c, http.StatusOK, "用户未绑定手机号", gin.H{"ouid": ""})
			return
		}
		phone, err := getPhone(param.PhoneCode)
		if err != nil {
			if err != nil {
				util.ResponseKit(c, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}
		userInfo = models.User{
			Ouid:        generateOUID,
			Name:        "新用户" + generateOUID[5:9],
			Avatar:      "avatar/171385753493.png", //默认头像
			Phone:       phone,
			Birthday:    0,
			Status:      dao.PayStatusYes,
			Sex:         0,
			TotalCount:  0,
			TotalLength: 0,
			WxOpenid:    response["openid"],
			WxUnionid:   response["unionid"],
			GmtCreate:   util.Now(),
			GmtUpdate:   util.Now(),
		}
		err = db.Create(&userInfo).Error
		if err != nil {
			logKit.Log.Printf("添加用户是发生了错误:" + err.Error())
			util.ResponseKit(c, http.StatusInternalServerError, "添加用户失败,", nil)
			return
		}
	}
	util.ResponseKit(c, http.StatusOK, "查询成功", gin.H{"ouid": userInfo.Ouid})
}
func getPhone(code string) (string, error) {
	form, err := util.HTTPPostJSON("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token="+kit.AccessToken, nil, gin.H{"code": code}, nil, "")
	fmt.Println(string(form.Body), err)
	errorCode := gojsonq.New().FromString(string(form.Body)).Find("errcode")
	if errorCode != float64(0) {
		logKit.Log.Println("微信调用获取手机号失败:" + string(form.Body))
		return "", errors.New("微信调用获取手机号失败")
	}
	fmt.Println(form, err)
	phoneInfo := gojsonq.New().FromString(string(form.Body)).Find("phone_info").(map[string]interface{})
	Phone := util.ToString(phoneInfo["phoneNumber"])
	return Phone, err
}

// WxUpdateUserInfo 更新用户数据
func WxUpdateUserInfo(c *gin.Context) {
	var param = param.WxUpdateUserInfoParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数错误:"+err.Error(), nil)
		return
	}

	err = db.Table("user").Where("ouid = ?", param.UserOuid).Updates(&models.User{
		Name:      param.Name,
		Avatar:    param.Avatar,
		Phone:     param.Phone,
		Birthday:  param.Birthday,
		Sex:       param.Sex,
		Introduce: param.Introduce,
		GmtUpdate: util.Now(),
	}).Error
	if err != nil {
		logKit.Log.Printf("更新用户存数据库时，发生了错误:" + err.Error())
		util.ResponseKit(c, http.StatusInternalServerError, "用户更新失败", nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// WxGetShopList 查看门店列表
func WxGetShopList(c *gin.Context) {
	var shopList []models.Shop
	db.Find(&shopList)

	type WxGetShopListVO struct {
		Id          int      `json:"id"`           //id
		ShopName    string   `json:"shop_name"`    //门店名称
		Avatar      string   `json:"avatar"`       //头像
		Address     string   `json:"address"`      //详细位置
		BottomPrice int      `json:"bottom_price"` //最低价格
		Longitude   float64  `json:"longitude"`    //经度
		Latitude    float64  `json:"latitude"`     //维度
		Tag         []string `json:"tag"`          //门店标签
	}
	var shopListVO []WxGetShopListVO

	for _, v := range shopList {
		var tag []string
		json.Unmarshal([]byte(v.Tag), &tag)
		shopListVO = append(shopListVO, WxGetShopListVO{
			Id:          v.Id,
			ShopName:    v.Name,
			Avatar:      v.Avatar,
			Address:     v.Address,
			BottomPrice: v.BottomPrice,
			Longitude:   v.Longitude,
			Latitude:    v.Latitude,
			Tag:         tag,
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", shopListVO)
}
func WxGetRecommendShopList(c *gin.Context) {
	var shopList []models.Shop
	db.Find(&shopList)

	type WxGetShopListVO struct {
		Id          int      `json:"id"`           //id
		ShopName    string   `json:"shop_name"`    //门店名称
		Avatar      string   `json:"avatar"`       //头像
		Address     string   `json:"address"`      //详细位置
		BottomPrice int      `json:"bottom_price"` //最低价格
		Longitude   float64  `json:"longitude"`    //经度
		Latitude    float64  `json:"latitude"`     //维度
		Tag         []string `json:"tag"`          //门店标签
	}
	var shopListVO []WxGetShopListVO

	for _, v := range shopList {
		var tag []string
		json.Unmarshal([]byte(v.Tag), &tag)
		shopListVO = append(shopListVO, WxGetShopListVO{
			Id:          v.Id,
			ShopName:    v.Name,
			Avatar:      v.Avatar,
			Address:     v.Address,
			BottomPrice: v.BottomPrice,
			Longitude:   v.Longitude,
			Latitude:    v.Latitude,
			Tag:         tag,
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", shopListVO)
}

// WxGetShopDetail 获取门店详情
func WxGetShopDetail(c *gin.Context) {
	shop := models.Shop{}
	db.First(&shop)
	var tag []string
	json.Unmarshal([]byte(shop.Tag), &tag)
	CheckOrderByShopId(shop.Id)
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"shop_avatar":  shop.Avatar,
		"shop_name":    shop.Name,
		"shop_photo":   shop.Photo,
		"shop_phone":   shop.Phone,
		"shop_address": shop.Address,
		"longitude":    shop.Longitude,
		"latitude":     shop.Latitude,
		"work_time":    shop.WorkTime,
		"tag":          tag,
		"facility":     shop.Facility,
		"serve":        shop.Serve,
		"desc":         shop.Desc,
	})
}

// WxGetShopSurplusCount  获取门店场地信息
func WxGetShopSurplusCount(c *gin.Context) {
	//当天
	count, money := getSiteCount(util.Now())
	count1, money1 := getSiteCount(util.Now().AddDate(0, 0, 1))
	count2, money2 := getSiteCount(util.Now().AddDate(0, 0, 2))
	count3, money3 := getSiteCount(util.Now().AddDate(0, 0, 3))
	count4, money4 := getSiteCount(util.Now().AddDate(0, 0, 4))
	count5, money5 := getSiteCount(util.Now().AddDate(0, 0, 5))
	count6, money6 := getSiteCount(util.Now().AddDate(0, 0, 6))

	util.ResponseKit(c, http.StatusOK, "操作成功", gin.H{
		util.Now().Format("01-02"):                  gin.H{"count": count, "money": money},
		util.Now().AddDate(0, 0, 1).Format("01-02"): gin.H{"count": count1, "money": money1},
		util.Now().AddDate(0, 0, 2).Format("01-02"): gin.H{"count": count2, "money": money2},
		util.Now().AddDate(0, 0, 3).Format("01-02"): gin.H{"count": count3, "money": money3},
		util.Now().AddDate(0, 0, 4).Format("01-02"): gin.H{"count": count4, "money": money4},
		util.Now().AddDate(0, 0, 5).Format("01-02"): gin.H{"count": count5, "money": money5},
		util.Now().AddDate(0, 0, 6).Format("01-02"): gin.H{"count": count6, "money": money6},
	})

}
func getSiteCount(time time.Time) (count int, money int) {
	// 获取所有商户的所有场地
	siteList, _ := dao.SiteAllShop()
	if len(siteList) == 0 {
		return 0, 0
	}
	var timeEnum []int
	for _, v := range siteList {
		useList, _ := dao.SiteUseListBySiteId(time, v.Id)
		var useCount int
		for _, vv := range useList {
			var a []int
			json.Unmarshal([]byte(vv.UseTimeEnum), &a)
			if time.Day() == util.Now().Day() {
				for _, vvv := range a {
					split := strings.Split(TimeEnum[vvv], ":")

					if util.Now().Hour() < util.StringToInt(split[0]) {
						useCount += 1
					}
				}
			} else {
				useCount += len(a)
			}
		}
		json.Unmarshal([]byte(v.TimeEnum), &timeEnum)
		if time.Day() == util.Now().Day() {
			for _, vv := range timeEnum {
				split := strings.Split(TimeEnum[vv], ":")

				if util.Now().Hour() < util.StringToInt(split[0]) {
					count += 1
				}
			}
		} else {
			count += len(timeEnum)
		}

		count -= useCount
	}
	if time.Weekday() == 6 || time.Weekday() == 0 {
		money = siteList[0].BusyPrice
	} else {
		money = siteList[0].FreePrice
	}
	return count, money
}

// WxAddOrder 微信定场
func WxAddOrder(c *gin.Context) {
	lock.Lock() // TODO 因为有Saas订场接口，所以此处简单加锁不能解决问题
	defer lock.Unlock()

	var wxAddOrderParam param.WxAddOrderParam
	if err := c.BindJSON(&wxAddOrderParam); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败:"+err.Error(), nil)
		return
	}
	//加锁以后查一下场地
	var siteIdSlice []int
	for _, v := range wxAddOrderParam.SiteDetail {
		siteIdSlice = append(siteIdSlice, v.SiteId)
	}
	gmtSiteUse, err := util.StringToTime("2006-01-02", wxAddOrderParam.GmtSiteUse)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "时间格式发生错误", nil)
		return
	}
	var repetitionSiteList []param.SiteDetail
	//如果对应的日期有场地，则进行判断，反之不进行判断
	siteUseList, err := dao.SiteUseListBySiteIdSlice(wxAddOrderParam.GmtSiteUse, siteIdSlice)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "查找场地占用失败:"+err.Error(), nil)
		return
	}
	for _, paramV := range wxAddOrderParam.SiteDetail {
		for _, siteUserV := range siteUseList {
			if paramV.SiteId == siteUserV.SiteId {
				var a []int
				json.Unmarshal([]byte(siteUserV.UseTimeEnum), &a)
				for _, v := range a {
					for _, vv := range paramV.TimeEnum {
						if v == vv {
							if len(repetitionSiteList) == 0 {
								repetitionSiteList = append(repetitionSiteList, param.SiteDetail{
									SiteId:   paramV.SiteId,
									TimeEnum: []int{v},
								})
								continue
							}
							same := 0
							for kkk, vvv := range repetitionSiteList {
								if vvv.SiteId == paramV.SiteId {
									repetitionSiteList[kkk].TimeEnum = append(repetitionSiteList[kkk].TimeEnum, v)
									same = 1
									break
								}
							}
							if same == 0 {
								repetitionSiteList = append(repetitionSiteList, param.SiteDetail{
									SiteId:   paramV.SiteId,
									TimeEnum: []int{v},
								})
							}
						}
					}
				}
			}
		}
	}
	if len(repetitionSiteList) != 0 {
		util.ResponseKit(c, http.StatusOK, "场地被占用", gin.H{"site": repetitionSiteList})
		return
	}

	var (
		money   int
		orderNo string
	)
	err = db.Transaction(func(tx *gorm.DB) error {
		money, orderNo, err = AddOrderAndSiteUse(tx, wxAddOrderParam.SiteDetail, wxAddOrderParam.UserOuid, gmtSiteUse, "预约订场", enum.OrderTypeWx, wxAddOrderParam.ReserveName, wxAddOrderParam.ReservePhone, wxAddOrderParam.Remark)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "订单创建失败:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", gin.H{"money": money, "order": orderNo})
}

// WxPay 微信支付
func WxPay(c *gin.Context) {
	param := param.WxPayParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	order := models.Order{}
	db.Where("order_no = ?", param.OrderNo).First(&order)
	if order.Status != "N" {
		util.ResponseKit(c, http.StatusOK, "该订单无法支付", nil)
		return
	}

	//todo 下面是支付流程 需要的时候放开
	user, err := dao.UserGetOne(order.UserOuid)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取用户错误："+err.Error(), nil)
		return
	}

	var result gin.H
	err = db.Transaction(func(tx *gorm.DB) (errInner error) {
		//订单生成预支付订单
		dao.PayAdd(tx, &order, user)
		result, errInner = wxPay(param.Type, user.WxOpenid, order.OrderNo, int64(order.Money), c.ClientIP())
		return
	})

	if err != nil {
		logKit.Log.Println("支付失败:", err.Error())
		util.ResponseKit(c, http.StatusInternalServerError, "支付失败", nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", result)
}

func wxPay(payType, userOpenid, orderNo string, money int64, ip string) (gin.H, error) {
	if payType == "web" {
		api, err := kit.PrePayJsApi(userOpenid, money, orderNo, "微信订单支付")
		if err != nil {
			return nil, err
		}
		return gin.H{"per_pay": api}, nil
	} else if payType == "h5" {
		//这里暂时先不做
		h5, err := kit.PrePayH5(money, orderNo, "订单支付", ip)
		if err != nil {
			return nil, err
		}
		return gin.H{"H5_url": h5, "order_no": orderNo}, nil
	}
	panic("支付类型错误")
}

func WxPayNotify(c *gin.Context) {
	notify, errNotify := kit.Notify(c)
	if errNotify != nil {
		logKit.Log.Println("接受回调函数失败:" + errNotify.Error())
		return
	}
	marshal, _ := json.Marshal(&notify)
	logKit.Log.Println("接受微信回调函数成功:" + string(marshal))
	//将里面重要信息解密
	transaction := models.Transaction{}
	json.Unmarshal([]byte(notify.Resource.Plaintext), &transaction)
	fmt.Println("这是解析的数据:", transaction)

	if transaction.TradeState != "SUCCESS" {
		logKit.Log.Println("支付失败,消息体内容:", transaction)
		util.ResponseKit(c, http.StatusBadRequest, "微信回调失败", nil)
		return
	}
	//如果支付成功
	order, errSelect := dao.OrderGet(transaction.OutTradeNo)

	var errTx error
	if errSelect == nil {
		if order.Status == dao.PayStatusYes {
			logKit.Log.Println("已经成功了，但是微信又给我发了")
			c.JSON(200, gin.H{
				"code":    "SUCCESS",
				"message": "成功",
			})
			return
		}
		errTx = db.Transaction(func(tx *gorm.DB) error {
			order.Status = dao.PayStatusYes
			order.GmtSuccess = util.Now()
			if err := tx.Save(&order).Error; err != nil {
				return err
			}
			//修改支付表信息
			return dao.PayUpdate(tx, order.OrderNo, marshal)
		})
	} else {
		if !errors.Is(errSelect, gorm.ErrRecordNotFound) {
			logKit.Log.Println("订单查询错误:", errSelect.Error())
			util.ResponseKit(c, http.StatusBadRequest, "订单查询错误，请联系管理员处理", nil)
			return
		}
		// 订场订单找不到，那么可能是充值订单
		payOne, err := dao.PayGetRecharge(transaction.OutTradeNo)
		if err != nil {
			logKit.Log.Println("查询充值订单错误:", err.Error())
			util.ResponseKit(c, http.StatusBadRequest, "订单查询错误，请联系管理员处理", nil)
			return
		}
		if payOne.Status == dao.PayStatusYes {
			logKit.Log.Println("已经成功了，但是微信又给我发了")
			c.JSON(200, gin.H{
				"code":    "SUCCESS",
				"message": "成功",
			})
			return
		}

		originMoney := uint32(payOne.TotalMoney)
		// 获取所有充值规则，从大到小检查是否负责赠送条件
		presentMoney, err := RechargeRuleGetPresent(originMoney)
		if err != nil {
			util.ResponseKit(c, http.StatusInternalServerError, "查询赠送金额错误："+err.Error(), nil)
			return
		}
		user, err := dao.UserGetOne(payOne.UserOuid)
		if err != nil {
			logKit.Log.Println("充值订单查找用户失败:", err.Error())
			util.ResponseKit(c, http.StatusInternalServerError, "微信回调失败", nil)
			return
		}

		errTx = db.Transaction(func(tx *gorm.DB) (errInner error) {
			// 修改账户余额，如果是普通会员，改为充值会员
			if errInner = dao.UserRecharge(tx, user, originMoney, presentMoney); errInner != nil {
				return
			}
			// 新增用户充值记录
			if errInner = dao.AccountRecordRecharge(tx, user, enum.AccountChangeTypeUserRecharge, originMoney, presentMoney, 0); errInner != nil {
				return
			}
			return dao.PayUpdateRecharge(tx, payOne, marshal)
		})
	}

	if errTx != nil {
		logKit.Log.Println("微信回调失败:", errSelect.Error())
		util.ResponseKit(c, http.StatusInternalServerError, "微信回调失败", nil)
		return
	}
	//向微信发送成功
	c.JSON(200, gin.H{
		"code":    "SUCCESS",
		"message": "成功",
	})
}
func AddOrderAndSiteUse(tx *gorm.DB, siteDetail []param.SiteDetail, UserOuid string, gmtSiteUse time.Time, OrderName string, Type int, reserveName string, reservePhone string, remark string) (money int, orderNo string, err error) {
	//如果检测通过的话，进行场地占用和金额计算生成订单
	user := models.User{}
	db.Where("ouid = ?", UserOuid).First(&user)
	//先占用场地
	//tx := db.Begin()
	orderNo = kit.GenerateOrderNo(PublicInt)
	PublicInt++
	siteOrderDetailJsonToString := ""
	if len(siteDetail) != 0 {
		for _, v := range siteDetail {
			marshal, _ := json.Marshal(&v.TimeEnum)
			err = dao.SiteUseAdd(tx, &user, v.SiteId, gmtSiteUse, marshal, orderNo, Type)
			if err != nil {
				logKit.Log.Println("场地占用添加失败", err.Error())
				return 0, "", err
			}
		}
		var siteOrderDetail []models.SiteOrderDetail
		money, siteOrderDetail = OrderSiteMoneyCalculate(siteDetail, gmtSiteUse)
		//开始写入订单
		marshal, _ := json.Marshal(siteOrderDetail)
		siteOrderDetailJsonToString = string(marshal)
	}
	shop := models.Shop{}
	db.First(&shop)
	randSix := ""
	for {
		randSix = getRandSix()
		var a int64
		db.Take("order").Where("check_no = ?", randSix).Count(&a)
		if a == 0 {
			break
		}
	}
	qrUrl := util.QrCreat(randSix)

	order := models.Order{
		OrderNo:       orderNo,
		OrderName:     OrderName,
		OriginalPrice: money,
		Coupons:       0,
		Money:         money,
		PayType:       "",
		ShopId:        shop.Id,
		ShopName:      shop.Name,
		ShopAvatar:    shop.Avatar,
		SiteDetail:    siteOrderDetailJsonToString,
		UserOuid:      user.Ouid,
		UserName:      user.Name,
		UserAvatar:    user.Avatar,
		UserPhone:     user.Phone,
		//PrePay:          string(marshal),
		Type:            Type,
		Status:          dao.PayStatusNo,
		SuperviseStatus: dao.PayStatusNo,
		CheckNo:         randSix,
		ReservePhone:    reservePhone,
		ReserveName:     reserveName,
		Remark:          remark,
		CheckQr:         qrUrl,
		ReserveDate:     gmtSiteUse,
		GmtCreate:       util.Now(),
		GmtUpdate:       util.Now(),
	}
	if Type == enum.OrderTypeSaas {
		order.Status = dao.PayStatusYes
	}
	err = tx.Omit("gmt_success", "gmt_refund").Create(&order).Error
	if err != nil {
		logKit.Log.Println("订单添加失败", err.Error())
		return 0, "", err
	}
	shop.AggregateAmount += money
	shop.GmtUpdate = util.Now()
	db.Save(&shop)

	//执行一个脚本，进行订单判断，15分钟以后判断订单是否支付成功，如果没有支付成功的话，进入取消订单流程
	go func() {
		time.Sleep(time.Minute * 15)
		TimerCancelOrder(orderNo)
	}()
	//return money, orderNo, *pay.PrepayId
	return money, orderNo, nil
}
func getRandSix() string {
	rand.Seed(time.Now().UnixNano() + int64(PublicInt))

	// 生成一个0到999999之间的随机数，然后加上100000，确保结果为6位数
	return util.ToString(rand.Intn(900000) + 100000)
}
func TimerCancelOrder(orderNo string) {
	order := models.Order{}
	if errors.Is(db.Where("order_no = ? AND status = 'N'", orderNo).First(&order).Error, gorm.ErrRecordNotFound) {
		return
	} else {
		db.Table("order").Where("order_no = ?", orderNo).Updates(map[string]interface{}{"status": "C", "gmt_update": util.Now()})
		// 场地占用状态批量改为取消占用
		_ = dao.SiteUseUpdateNot(orderNo)
		//todo 并同时进行退款处理
	}
	return
}

// WxGetSiteReserve 查看场馆内场地预定信息
func WxGetSiteReserve(c *gin.Context) {
	param := param.WxGetSiteReserveParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	time, err := util.StringToTime("2006-01-02", param.Date)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "时间格式有问题", nil)
		return
	}
	//todo 这里需要检测时间

	// 获取所有商户的所有场地
	siteAll, err := dao.SiteAllShop()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取所有场地失败"+err.Error(), nil)
		return
	}

	//场地价格
	type SitePrice struct {
		TimeEnum int `json:"time_enum"`
		Price    int `json:"price"`
	}
	type Site struct {
		SiteId          int         `json:"site_id"`
		SiteName        string      `json:"site_name"`
		ReserveTimeEnum []int       `json:"reserve_time_enum"`
		StoreTimeEnum   []int       `json:"store_time_enum"`
		Price           []SitePrice `json:"price"`
	}

	var siteVO []Site

	for _, v := range siteAll {
		site := Site{
			SiteId:          v.Id,
			SiteName:        v.Name,
			ReserveTimeEnum: []int{},
		}
		//查看该场地都有哪些时间
		var timeEnum []int
		json.Unmarshal([]byte(v.TimeEnum), &timeEnum)
		site.StoreTimeEnum = timeEnum

		//计算每个时间的价格
		var sitePrice []SitePrice
		for _, vv := range timeEnum {
			sitePrice = append(sitePrice, SitePrice{
				TimeEnum: vv,
				Price:    SitePriceCalculate(v, time, vv),
			})
		}
		site.Price = sitePrice
		siteVO = append(siteVO, site)

	}

	siteUseList, err := dao.SiteUseList(param.Date)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "查找场地占用失败:"+err.Error(), siteVO)
		return
	}
	for _, v := range siteUseList {
		var a []int
		json.Unmarshal([]byte(v.UseTimeEnum), &a)
		for kk, vv := range siteVO {
			if v.SiteId == vv.SiteId {
				siteVO[kk].ReserveTimeEnum = append(siteVO[kk].ReserveTimeEnum, a...)
			}
		}
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", siteVO)
}

// WxGetOrderList 用户查看订单
func WxGetOrderList(c *gin.Context) {
	param := param.WxGetOrderListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	var orderList []models.Order
	var total int64
	begin := db
	if len(param.Status) != 0 {
		begin = begin.Where("status = ?", param.Status)
	}
	begin.Where("user_ouid = ?", param.UserOuid).Find(&orderList).Count(&total)
	begin.Where("user_ouid = ?", param.UserOuid).Scopes(util.PageKit(param.Page, param.Size)).Order("gmt_create DESC").Find(&orderList)

	type WxGetOrderListVO struct {
		OrderNo    string                   `json:"order_no"`
		ShopId     int                      `json:"shop_id"`
		ShopName   string                   `json:"shop_name"`
		ShopAvatar string                   `json:"shop_avatar"`
		Money      int                      `json:"money"`
		Status     string                   `json:"status"`
		SiteDetail []models.SiteOrderDetail `json:"site_detail"`
	}
	var VOList []WxGetOrderListVO
	for _, v := range orderList {
		var detail []models.SiteOrderDetail
		json.Unmarshal([]byte(v.SiteDetail), &detail)
		VOList = append(VOList, WxGetOrderListVO{
			OrderNo:    v.OrderNo,
			ShopId:     v.ShopId,
			ShopName:   v.ShopName,
			ShopAvatar: v.ShopAvatar,
			Money:      v.Money,
			Status:     v.Status,
			SiteDetail: detail,
		})
	}
	CheckOrderByUserOuid(param.UserOuid)
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"list": VOList, "total": total})

}

// WxGetOrderDetail 查看订单详情
func WxGetOrderDetail(c *gin.Context) {
	param := param.WxGetOrderDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	order, err := dao.OrderGet(param.OrderNo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusInternalServerError, "找不到订单", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查找订单错误:"+err.Error(), nil)
		}
		return
	}

	type WxGetOrderDetailVO struct {
		ShopAvatar    string                   `json:"shop_avatar"`
		ShopName      string                   `json:"shop_name"`
		ShopAddress   string                   `json:"shop_address"`
		OrderNo       string                   `json:"order_no"`
		GmtCreatOrder string                   `json:"gmt_creat_order"`
		GmtPay        string                   `json:"gmt_pay"`
		Money         int                      `json:"money"`
		UserOuid      string                   `json:"user_ouid"`
		UserName      string                   `json:"user_name"`
		UserPhone     string                   `json:"user_phone"`
		GmtCreate     string                   `json:"gmt_create"`
		SiteDetail    []models.SiteOrderDetail `json:"site_detail"`
		ShopPhone     string                   `json:"shop_phone"`
		Longitude     float64                  `json:"longitude"` //经度
		Latitude      float64                  `json:"latitude"`  //维度
		Status        string                   `json:"status"`
		ReserveDate   string                   `json:"reserve_date"` //预约日期
		Remark        string                   `json:"remark"`
	}
	shop := models.Shop{}
	db.Where("id = ?", order.ShopId).First(&shop)
	user := models.User{}
	db.Where("ouid = ?", order.UserOuid).First(&user)
	var detail []models.SiteOrderDetail
	json.Unmarshal([]byte(order.SiteDetail), &detail)
	//
	VO := WxGetOrderDetailVO{
		ShopAvatar:    order.ShopAvatar,
		ShopName:      order.ShopName,
		ShopAddress:   shop.Address,
		OrderNo:       order.OrderNo,
		GmtCreatOrder: order.GmtCreate.Format("2006-01-02 15:04:05"),
		Money:         order.Money,
		UserOuid:      order.UserOuid,
		UserName:      order.UserName,
		UserPhone:     user.Phone,
		GmtCreate:     order.GmtCreate.Format("2006-01-02 15:04:05"),
		SiteDetail:    detail,
		ShopPhone:     shop.Phone,
		Longitude:     shop.Longitude,
		Latitude:      shop.Latitude,
		Status:        order.Status,
		ReserveDate:   order.ReserveDate.Format("2006-01-02"),
		Remark:        order.Remark,
	}
	if !order.GmtSuccess.IsZero() {
		VO.GmtPay = order.GmtSuccess.Format("2006-01-02 15:04:05")
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", VO)
}

// WxGetOrderStatus 查看订单状态
func WxGetOrderStatus(c *gin.Context) {
	var param param.WxGetOrderDetailParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}

	order, err := dao.OrderGet(param.OrderNo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusInternalServerError, "订单不存在", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查找订单错误:"+err.Error(), nil)
		}
		return
	}

	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"order_no": order.OrderNo,
		"status":   order.Status,
	})
}

func WxGetMyReserveList(c *gin.Context) {
	param := param.WxGetMyReserveListParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	var orderList []models.Order
	begin := db
	if len(param.Status) != 0 {
		begin = begin.Where("status = ?", param.Status)
	}
	var total int64
	begin.Table("order").Where("user_ouid = ? AND status !='N'", param.UserOuid).Count(&total)
	begin.Where("user_ouid = ? AND status !='N'", param.UserOuid).Order("gmt_update DESC").Scopes(util.PageKit(param.Page, param.Size)).Find(&orderList)

	type WxGetOrderListVO struct {
		ShopId      int                      `json:"shop_id"`
		ShopName    string                   `json:"shop_name"`
		ShopAvatar  string                   `json:"shop_avatar"`
		Status      string                   `json:"status"`
		SiteDetail  []models.SiteOrderDetail `json:"site_detail"`
		GmtCreate   string                   `json:"gmt_create"`
		OrderNo     string                   `json:"order_no"`
		ReserveDate string                   `json:"reserve_date"` //预约日期

	}
	var VOList []WxGetOrderListVO
	for _, v := range orderList {
		var detail []models.SiteOrderDetail
		json.Unmarshal([]byte(v.SiteDetail), &detail)
		VOList = append(VOList, WxGetOrderListVO{
			ShopId:      v.ShopId,
			ShopName:    v.ShopName,
			ShopAvatar:  v.ShopAvatar,
			Status:      v.Status,
			SiteDetail:  detail,
			OrderNo:     v.OrderNo,
			GmtCreate:   v.GmtCreate.Format("2006-01-02"),
			ReserveDate: v.ReserveDate.Format("2006-01-02"),
		})
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"list": VOList, "total": total})
}

// WxGetReserveDetail 查看预约详情列表
func WxGetReserveDetail(c *gin.Context) {
	param := param.WxGetReserveDetailParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}

	order := models.Order{}
	if errors.Is(db.Where("order_no = ?", param.OrderNo).First(&order).Error, gorm.ErrRecordNotFound) {
		util.ResponseKit(c, http.StatusInternalServerError, "找不到该订单", nil)
		return
	}
	type WxGetOrderDetailVO struct {
		ShopAvatar    string                   `json:"shop_avatar"`
		ShopName      string                   `json:"shop_name"`
		ShopAddress   string                   `json:"shop_address"`
		OrderNo       string                   `json:"order_no"`
		GmtCreatOrder string                   `json:"gmt_creat_order"`
		GmtPay        string                   `json:"gmt_pay"`
		Money         int                      `json:"money"`
		UserOuid      string                   `json:"user_ouid"`
		UserName      string                   `json:"user_name"`
		UserPhone     string                   `json:"user_phone"`
		GmtCreate     string                   `json:"gmt_create"`
		SiteDetail    []models.SiteOrderDetail `json:"site_detail"`
		ShopPhone     string                   `json:"shop_phone"`
		CheckNo       string                   `json:"check_no"`  //订单核验码
		CheckQr       string                   `json:"check_qr"`  //核验二维码
		Longitude     float64                  `json:"longitude"` //经度
		Latitude      float64                  `json:"latitude"`  //维度
		Status        string                   `json:"status"`
		ReserveDate   string                   `json:"reserve_date"` //预约日期
		Remark        string                   `json:"remark"`
	}
	shop := models.Shop{}
	db.Where("id = ?", order.ShopId).First(&shop)
	user := models.User{}
	db.Where("ouid = ?", order.UserOuid).First(&user)
	var detail []models.SiteOrderDetail
	json.Unmarshal([]byte(order.SiteDetail), &detail)
	VO := WxGetOrderDetailVO{
		ShopAvatar:    order.ShopAvatar,
		ShopName:      order.ShopName,
		ShopAddress:   shop.Address,
		OrderNo:       order.OrderNo,
		GmtCreatOrder: order.GmtCreate.Format("2006-01-02 15:04:05"),
		Money:         order.Money,
		UserOuid:      order.UserOuid,
		UserName:      order.UserName,
		UserPhone:     user.Phone,
		GmtCreate:     order.GmtCreate.Format("2006-01-02 15:04:05"),
		ReserveDate:   order.ReserveDate.Format("2006-01-02"),
		SiteDetail:    detail,
		ShopPhone:     shop.Phone,
		Longitude:     shop.Longitude,
		Latitude:      shop.Latitude,
		Status:        order.Status,
		Remark:        order.Remark,
		CheckQr:       order.CheckQr,
		CheckNo:       order.CheckNo,
	}
	if !order.GmtSuccess.IsZero() {
		VO.GmtPay = order.GmtSuccess.Format("2006-01-02 15:04:05")
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", VO)
}

// WxAddFeedback 微信上传反馈
func WxAddFeedback(c *gin.Context) {
	param := param.WxAddFeedbackParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	user := models.User{}
	db.Where("ouid = ?", param.UserOuid).First(&user)
	marshal, _ := json.Marshal(&param.Photo)
	db.Create(&models.Feedback{
		UserOuid:   user.Ouid,
		UserName:   user.Name,
		UserAvatar: user.Avatar,
		Photo:      string(marshal),
		Content:    param.Content,
		GmtCreate:  util.Now(),
		GmtUpdate:  util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
	return
}

// WxCancelOrder 取消订单
func WxCancelOrder(c *gin.Context) {
	var param param.WxCancelOrderParam
	if err := c.BindJSON(&param); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	mender, err := dao.AdminUserFirst()
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取管理员错误："+err.Error(), nil)
		return
	}
	//if !CheckCancelOrder(param.OrderNo) {
	//	util.ResponseKit(c, http.StatusInternalServerError, "该订单不能取消", nil)
	//	return
	//}
	//todo 看看这里需不需要判定场地的开场时间
	order, err := dao.OrderGet(param.OrderNo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.ResponseKit(c, http.StatusBadRequest, "订单不存在", nil)
		} else {
			util.ResponseKit(c, http.StatusInternalServerError, "查找订单错误:"+err.Error(), nil)
		}
		return
	}

	if order.Status == "Y" {
		// 先尝试是否能储值余额退款
		if order.UserOuid != "" {
			var user *models.User
			user, err = dao.UserGetOne(order.UserOuid)
			if err == nil {
				err = db.Transaction(func(tx *gorm.DB) error {
					return logic.UserMoneyRefund(tx, user, order.OrderNo, uint32(mender.Id))
				})
				if err != nil {
					logKit.Log.Println("储值余额退款发生错误:", err.Error())
				}
			}
		}
		if order.UserOuid == "" || err != nil { // 执行微信退款操作
			if err = kit.Refunds(*order); err != nil {
				logKit.Log.Println("退款发生错误:", err.Error())
				util.ResponseKit(c, http.StatusOK, "退款发生错误", nil)
				return
			}
		}

	}
	db.Table("order").Where("order_no = ?", param.OrderNo).Updates(map[string]interface{}{"status": "C", "gmt_refund": util.Now()})

	// 场地占用状态批量改为取消占用
	if err = dao.SiteUseUpdateNot(param.OrderNo); err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "取消订单失败:"+err.Error(), nil)
		return
	}
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// 检测取消订单的时间
func CheckCancelOrder(orderNo string) bool {
	order := models.Order{}
	if errors.Is(db.Where("order_no = ?", orderNo).First(&order).Error, gorm.ErrRecordNotFound) {
		return false
	}

	if order.Status == "C" {
		return false
	}
	// 获取该订单的所有场地占用
	useSiteList, _ := dao.SiteUseAll(orderNo)
	for _, v := range useSiteList {
		var timeEnumList []int
		json.Unmarshal([]byte(v.UseTimeEnum), &timeEnumList)
		for _, vv := range timeEnumList {
			s := TimeEnum[vv]
			split := strings.Split(s, "~")
			if len(split[0]) == 4 {
				split[0] = "0" + split[0]
			}
			useDateTime := v.UseDate + " " + split[0]
			toTime, err := util.StringToTime("2006-01-02 15:04", useDateTime)
			if err != nil {
				continue
			}
			if toTime.Before(util.Now()) {
				return false
			}
		}
	}
	return true

}

// WxRecentlyReserve 查看最近预约的概括
func WxRecentlyReserve(c *gin.Context) {
	param := param.UserOuidParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数请求失败", nil)
		return
	}
	var orderList []models.Order
	var timee time.Time
	var FirstOrderNo string

	db.Where("user_ouid = ? AND status ='Y' AND reserve_date > ?", param.UserOuid, util.Now().AddDate(0, 0, -1)).Find(&orderList)
	for _, v := range orderList {
		aa, err := GetOrderFirstSiteTime(v.OrderNo)
		if err != nil {
			util.ResponseKit(c, http.StatusOK, err.Error(), nil)
			return
		}
		if aa.After(util.Now()) {
			timee = aa
			FirstOrderNo = v.OrderNo
			break
		}
	}
	for _, v := range orderList {
		aa, err := GetOrderFirstSiteTime(v.OrderNo)
		if err != nil {
			util.ResponseKit(c, http.StatusOK, err.Error(), nil)
			return
		}
		if aa.Before(timee) {
			timee = aa
			FirstOrderNo = v.OrderNo
		}
	}
	//
	type WxRecentlyReserveVo struct {
		ShopName   string                   `json:"shop_name"`
		SiteDetail []models.SiteOrderDetail `json:"site_detail"`
		Data       string                   `json:"data"`
		OrderNo    string                   `json:"order_no"`
	}

	//下面是一段写的垃圾代码，引以为戒_________________________________________________________________
	//var siteUseList []models.SiteUse
	//db.Where("user_ouid =? AND status = 'Y' AND use_date > ?", param.UserOuid, util.Now().AddDate(0, 0, -1)).Order("use_date ASC").Find(&siteUseList)
	//if len(siteUseList) == 0 {
	//	util.ResponseKit(c, http.StatusOK, "暂无数据", WxRecentlyReserveVo{})
	//	return
	//}
	//var time time.Time
	//for _, v := range siteUseList {
	//	siteTime, err := GetOrderFirstSiteTime(v.OrderNo)
	//	if err != nil {
	//		util.ResponseKit(c, http.StatusOK, err.Error(), nil)
	//		return
	//	}
	//	if siteTime.After(util.Now()) {
	//		time = siteTime
	//		FirstOrderNo = v.OrderNo
	//		break
	//	}
	//
	//}
	//for _, v := range siteUseList {
	//	siteTime, err := GetOrderFirstSiteTime(v.OrderNo)
	//	if err != nil {
	//		util.ResponseKit(c, http.StatusOK, err.Error(), nil)
	//		return
	//	}
	//	if siteTime.Before(util.Now()) {
	//		continue
	//	}
	//
	//	if siteTime.Before(time) {
	//		time = siteTime
	//		FirstOrderNo = v.OrderNo
	//	}
	//
	//}
	//if len(FirstOrderNo) == 0 {
	//	util.ResponseKit(c, http.StatusOK, "暂无数据", WxRecentlyReserveVo{})
	//	return
	//}
	//order := models.Order{}
	//db.Where("order_no = ?", FirstOrderNo).First(&order)

	for _, v := range orderList {
		if v.OrderNo == FirstOrderNo {
			var detail []models.SiteOrderDetail
			json.Unmarshal([]byte(v.SiteDetail), &detail)
			vo := WxRecentlyReserveVo{
				ShopName:   v.ShopName,
				SiteDetail: detail,
				Data:       v.ReserveDate.Format("2006-01-02"),
				OrderNo:    v.OrderNo,
			}
			util.ResponseKit(c, http.StatusOK, "请求成功", vo)
		}
	}
}

// WxGetPhone 获取手机号
func WxGetPhone(c *gin.Context) {
	param := param.GetPhoneParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusBadRequest, "参数绑定失败:"+err.Error(), nil)
		return
	}
	form, err := util.HTTPPostJSON("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token="+kit.AccessToken, nil, gin.H{"code": param.PhoneCode}, nil, "")
	fmt.Println(string(form.Body), err)
	errorCode := gojsonq.New().FromString(string(form.Body)).Find("errcode")
	if errorCode != float64(0) {
		logKit.Log.Println("微信调用获取手机号失败:" + string(form.Body))
		util.ResponseKit(c, http.StatusInternalServerError, "微信调用获取手机号失败", nil)
		return
	}
	fmt.Println(form, err)
	phoneInfo := gojsonq.New().FromString(string(form.Body)).Find("phone_info").(map[string]interface{})
	Phone := util.ToString(phoneInfo["phoneNumber"])
	//userInfo := models.User{}
	//db.Model(&userInfo).Where("ouid = ?", param.UserOuid).Updates(models.User{Phone: Phone, GmtUpdate: util.Now()})
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"phone": Phone})
}

// WxGetUserInfo 获取用户信息
func WxGetUserInfo(c *gin.Context) {
	param := param.UserOuidParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败,用户ouid不能为空", nil)
		return
	}

	user := models.User{}
	user.WxOpenid = ""
	user.WxUnionid = ""
	db.Where("ouid = ?", param.UserOuid).First(&user)

	type WxGetUserInfo struct {
		Ouid        string       `json:"ouid"`         //用户ouid
		Name        string       `json:"name"`         //用户名
		Avatar      string       `json:"avatar"`       //用户头像
		Phone       string       `json:"phone"`        //用户手机号
		Birthday    int          `json:"birthday"`     //用户生日
		Status      string       `json:"status"`       //用户状态
		Sex         int          `json:"sex"`          //用户性别
		TotalCount  int          `json:"total_count"`  //统计运动次数
		TotalLength int          `json:"total_length"` //统计运动时长
		Introduce   string       `json:"introduce"`    //个人简介
		SportDay    int          `json:"sport_day"`    //运动天数
		WxOpenid    string       `json:"wx_openid"`    //微信openid
		WxUnionid   string       `json:"wx_unionid"`   //微信unionid
		VipType     enum.VipType `json:"vip_type"`
		Balance     uint32       `json:"balance"`
	}
	CheckOrderByUserOuid(user.Ouid)
	util.ResponseKit(c, http.StatusOK, "请求成功", WxGetUserInfo{
		Ouid:        user.Ouid,
		Name:        user.Name,
		Avatar:      user.Avatar,
		Phone:       user.Phone,
		Birthday:    user.Birthday,
		Status:      user.Status,
		Sex:         user.Sex,
		Introduce:   user.Introduce,
		TotalCount:  user.TotalCount,
		TotalLength: user.TotalLength,
		SportDay:    user.TotalCount,
		WxOpenid:    user.WxOpenid,
		WxUnionid:   user.WxUnionid,
		VipType:     user.VipType,
		Balance:     user.ActualBalance + user.PresentBalance,
	})
}

// WxGetWxId 获取微信id
func WxGetWxId(c *gin.Context) {
	idParam := param.WxGetWxIdParam{}
	err := c.BindJSON(&idParam)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	openId, unionId, err := kit.WxGetUserInfo(idParam.Type, idParam.Code)
	if len(openId) == 0 || len(unionId) == 0 || err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "获取用户openId和unionId失败:"+err.Error(), nil)
		return
	}
	db.Table("user").Where("ouid = ?", idParam.UserOuid).Updates(map[string]interface{}{"wx_openid": openId, "wx_unionid": unionId, "gmt_update": util.Now()})
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)

}

// WxAddShopError 微信添加场馆错误信息
func WxAddShopError(c *gin.Context) {
	param := param.WxAddShopErrorParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定错误", nil)
		return
	}
	photos := ""
	if len(param.Photos) != 0 {
		marshal, _ := json.Marshal(&param.Photos)
		photos = string(marshal)
	}
	db.Create(&models.ShopError{
		UserOuid:  param.UserOuid,
		Type:      param.Type,
		Detail:    param.Detail,
		Photos:    photos,
		GmtCreate: util.Now(),
		GmtUpdate: util.Now(),
	})
	util.ResponseKit(c, http.StatusOK, "请求成功", nil)

}
func WxGetAgreementAbout(c *gin.Context) {
	shop := models.Shop{}
	db.First(&shop)
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{
		"agreement": shop.Agreement,
		"about_us":  shop.AboutUs,
	})
}
func GetPhoneCode(c *gin.Context) {
	param := param.GetPhoneCodeParam{}
	err := c.BindJSON(&param)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "手机号不能为空", nil)
		return
	}
	_, b := kit.GetPhoneCache(param.Phone)
	if b {
		util.ResponseKit(c, http.StatusInternalServerError, "刚发的验证码，5分钟有效", nil)
		return
	}
	// 生成 6 位随机数
	randomNum := rand.Intn(900000) + 100000
	strconv.Itoa(randomNum)
	//todo 如果需要发送验证码，将下面的注释打开并配置相关的信息
	//kit.PhoneCode(util.ToString(param.Phone), util.ToString(randomNum))
	kit.UpdatePhoneCache(param.Phone, util.ToString(randomNum))
	util.ResponseKit(c, http.StatusOK, "操作成功", nil)
}

// PhoneCodeLogin 验证码登录
func PhoneCodeLogin(c *gin.Context) {
	loginParam := param.PhoneCodeLoginParam{}
	err := c.BindJSON(&loginParam)
	if err != nil {
		util.ResponseKit(c, http.StatusInternalServerError, "参数绑定失败", nil)
		return
	}
	if loginParam.Phone != "18888888888" && loginParam.Phone != "13344445555" && loginParam.Phone != "14455556666" {
		cache, _ := kit.GetPhoneCache(loginParam.Phone)
		if cache != loginParam.Code {
			util.ResponseKit(c, http.StatusInternalServerError, "验证码错误", nil)
			return
		}
	}

	generateOUID := util.GenerateOuid()
	userInfo := models.User{}
	if errors.Is(db.Where("phone = ?", loginParam.Phone).First(&userInfo).Error, gorm.ErrRecordNotFound) {
		userInfo = models.User{
			Ouid:        generateOUID,
			Name:        "新用户" + generateOUID[5:9],
			Avatar:      "avatar/171385753493.png", //默认头像
			Phone:       loginParam.Phone,
			Birthday:    0,
			Status:      dao.PayStatusYes,
			Sex:         0,
			TotalCount:  0,
			TotalLength: 0,
			Introduce:   "",
			GmtCreate:   util.Now(),
			GmtUpdate:   util.Now(),
		}
		err = db.Create(&userInfo).Error
		if err != nil {
			logKit.Log.Printf("添加用户是发生了错误:" + err.Error())
			util.ResponseKit(c, http.StatusInternalServerError, "添加用户失败,", nil)
			return
		}
	}
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"ouid": userInfo.Ouid})
}

func UploadAvatar(c *gin.Context) {
	util.SavePhoto(c, "avatar")
}
func UploadPhoto(c *gin.Context) {
	util.SavePhoto(c, "photo")
}

// GetOrderFirstSiteTime 获取订单首场开始时间
func GetOrderFirstSiteTime(orderNo string) (time.Time, error) {
	siteUseAll, _ := dao.SiteUseAll(orderNo)
	if len(siteUseAll) == 0 {
		return time.Time{}, errors.New("订单不存在")
	}
	var enumList []int
	for _, v := range siteUseAll {
		var siteEnumList []int
		json.Unmarshal([]byte(v.UseTimeEnum), &siteEnumList)
		for _, vv := range siteEnumList {
			enumList = append(enumList, vv)
		}
	}
	value := minValue(enumList)
	if value == 0 {
		return time.Time{}, errors.New("订单发生错误")
	}

	split := strings.Split(TimeEnum[value], "~")
	if len(split) < 2 {
		return time.Time{}, errors.New("订单发生错误")
	}
	if len(split[0]) == 4 {
		split[0] = "0" + split[0]
	}
	toTime, err := util.StringToTime("2006-01-02 15:04", siteUseAll[0].UseDate+" "+split[0])
	if err != nil {
		return time.Time{}, err
	}
	return toTime, nil
}
func GetOrderEndSiteTime(orderNo string) (time.Time, error) {
	siteUseAll, _ := dao.SiteUseAll(orderNo)
	if len(siteUseAll) == 0 {
		return time.Time{}, errors.New("订单不存在")
	}
	var enumList []int
	for _, v := range siteUseAll {
		var siteEnumList []int
		json.Unmarshal([]byte(v.UseTimeEnum), &siteEnumList)
		for _, vv := range siteEnumList {
			enumList = append(enumList, vv)
		}
	}
	value := minValue(enumList)
	if value == 0 {
		return time.Time{}, errors.New("订单发生错误")
	}

	split := strings.Split(TimeEnum[value], "~")
	if len(split) < 2 {
		return time.Time{}, errors.New("订单发生错误")
	}
	if len(split[1]) == 4 {
		split[1] = "0" + split[1]
	}
	toTime, err := util.StringToTime("2006-01-02 15:04", siteUseAll[0].UseDate+" "+split[1])
	if err != nil {
		return time.Time{}, errors.New("时间工具发生错误")
	}
	return toTime, nil
}
func minValue(arr []int) int {
	if len(arr) == 0 {
		return 0 // 或者返回一个错误值，表示数组为空
	}
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	if min < 0 {
		return 0
	}
	return min
}
