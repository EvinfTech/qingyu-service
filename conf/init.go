package conf

import (
	"github.com/goccy/go-json"
	"os"
)

type config struct {
	AlibabaCloudAccessKeyId     string `json:"alibaba_cloud_access_key_id"`
	AlibabaCloudAccessKeySecret string `json:"alibaba_cloud_access_key_secret"`
	SignName                    string `json:"sign_name"`
	WxAppid                     string `json:"wx_appid"`
	WxSecret                    string `json:"wx_secret"`
	WxNotifyUrl                 string `json:"wx_notify_url"`
	MchID                       string `json:"mch_id"`                        // 商户号
	MchCertificateSerialNumber  string `json:"mch_certificate_serial_number"` // 商户证书序列号
	MchAPIv3Key                 string `json:"mch_APIv3_key"`                 // 商户APIv3密钥
}

var Config config

func Init() {
	data, _ := os.ReadFile("conf.json")

	err := json.Unmarshal(data, &Config)
	if err != nil {
		panic("配置文件conf.json解码失败:" + err.Error())
	}
}
