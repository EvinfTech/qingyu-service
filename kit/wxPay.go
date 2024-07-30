package kit

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"qiyu/conf"
	"qiyu/dao"
	logKit "qiyu/logger"
	"qiyu/models"
	"qiyu/util"
	"time"
)

type contentType struct {
	Mchid           *string    `json:"mchid"`
	Appid           *string    `json:"appid"`
	CreateTime      *time.Time `json:"create_time"`
	OutContractCode *string    `json:"out_contract_code"`
}

var OfficialAccountsAppId string = ""
var mchPrivateKey *rsa.PrivateKey
var client *core.Client
var ctx context.Context

func WxPayInit() {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	var err error
	mchPrivateKey, err = utils.LoadPrivateKeyWithPath("./apiclient_key.pem")
	if err != nil {
		logKit.Log.Println("商户私钥加载失败:load merchant private key error")
	}

	ctx = context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(conf.Config.MchID, conf.Config.MchCertificateSerialNumber, mchPrivateKey, conf.Config.MchAPIv3Key),
	}
	client, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
		fmt.Println("商户私钥初始化失败:" + err.Error())
	}

	// 得到prepay_id，以及调起支付所需的参数和签名

}
func PrePayJsApi(userOpenId string, money int64, orderNo string, orderDescription string) (jsapi.PrepayWithRequestPaymentResponse, error) {
	if len(conf.Config.WxAppid) == 0 {
		return jsapi.PrepayWithRequestPaymentResponse{}, errors.New("请先配置微信支付")
	}
	if len(userOpenId) == 0 {
		return jsapi.PrepayWithRequestPaymentResponse{}, errors.New("用户openid为空")
	}
	svc := jsapi.JsapiApiService{Client: client}
	mm, _ := time.ParseDuration("15m")
	mm1 := util.Now().Add(mm)
	resp, rest, err := svc.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(conf.Config.WxAppid), //appid
			Mchid:       core.String(conf.Config.MchID),   //商户号
			Description: core.String(orderDescription),    //商品描述
			TimeExpire:  core.Time(mm1),
			OutTradeNo:  core.String(orderNo),                 //商户本地订单号
			Attach:      core.String("微信支付"),                  //自动逸数据说明
			NotifyUrl:   core.String(conf.Config.WxNotifyUrl), //描述
			Amount: &jsapi.Amount{
				Total: core.Int64(1),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(userOpenId),
			},
		},
	)
	fmt.Println(rest)
	if err == nil {
		fmt.Println("预支付返回来的值:", resp)
		return *resp, nil
	} else {
		logKit.Log.Println("支付发生错误", err.Error())
		return jsapi.PrepayWithRequestPaymentResponse{}, err
	}
}
func Notify(c *gin.Context) (notify.Request, error) {
	err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, conf.Config.MchCertificateSerialNumber, conf.Config.MchID, conf.Config.MchAPIv3Key)
	if err != nil {
		logKit.Log.Println("向 Mgr 注册商户的平台证书下载器失败:" + err.Error())
		return notify.Request{}, err
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(conf.Config.MchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(conf.Config.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	content := new(contentType)

	notifyReq, err := handler.ParseNotifyRequest(context.Background(), c.Request, content)
	if err != nil {
		return notify.Request{}, errors.New("解密失败:" + err.Error())
	}
	return *notifyReq, nil

}
func Refunds(order models.Order) error {

	svc := refunddomestic.RefundsApiService{Client: client}
	no := GenerateOrderNo(1)
	resp, result, err := svc.Create(ctx,
		refunddomestic.CreateRequest{
			//SubMchid:      core.String("1900000109"),
			//TransactionId: core.String("1217752501201407033233368018"),
			OutTradeNo:  core.String(order.OrderNo),
			OutRefundNo: core.String(no),
			Reason:      core.String("退款"),
			//NotifyUrl:     core.String("https://weixin.qq.com"),
			//FundsAccount:  refunddomestic.REQFUNDSACCOUNT_AVAILABLE.Ptr(),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				Refund:   core.Int64(1),
				Total:    core.Int64(1),
			},
		},
	)

	if err != nil {
		// 处理错误
		fmt.Println("我就说不行啊," + err.Error())
		return err
	} else {
		// 处理返回结果
		fmt.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
		db.Table("order").Where("id = ?", order.Id).Updates(map[string]interface{}{"status": "T", "refunds_time": util.Now(), "gmt_update": util.Now()})
		dao.PayRefund(db, &order, resp)
		return nil
	}
}
func PrePayH5(money int64, orderNo string, orderDescription string, ip string) (url string, err error) {
	//setting, err := dao.SettingGet()
	//if err != nil {
	//	return "", errors.New("获取设置失败:" + err.Error())
	//}
	mm, _ := time.ParseDuration("15m")
	mm1 := util.Now().Add(mm)
	svc := h5.H5ApiService{Client: client}
	resp, _, err := svc.Prepay(ctx,
		h5.PrepayRequest{
			Appid:       core.String(conf.Config.WxAppid), //appid
			Mchid:       core.String(conf.Config.MchID),   //商户号
			Description: core.String(orderDescription),    //商品描述
			TimeExpire:  core.Time(mm1),
			OutTradeNo:  core.String(orderNo),                 //商户本地订单号
			Attach:      core.String("微信支付"),                  //自动逸数据说明
			NotifyUrl:   core.String(conf.Config.WxNotifyUrl), //描述

			Amount: &h5.Amount{
				Total: core.Int64(1),
			},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String(ip),
				H5Info: &h5.H5Info{
					Type: core.String("iOS"),
				},
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Prepay err:%s", err)
		return "", err
	}
	return *resp.H5Url, nil
}
