package kit

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"net/http"
	"qiyu/conf"
	logKit "qiyu/logger"
	"qiyu/util"
	"time"
)

var AccessToken string

// GetAccessToken 定制配置自己的微信appid和secret
func GetAccessToken() {
	wxUrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + conf.Config.WxAppid + "&secret=" + conf.Config.WxSecret
	request, _ := http.NewRequest("GET", wxUrl, nil)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("获取微信小程序token失败:" + err.Error())
	}
	type res struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var response res
	respBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &response)
	AccessToken = response.AccessToken
	logKit.Log.Println("更新了access_token:" + AccessToken)
}

// AccessTokenInit 30分钟更新一次
func AccessTokenInit() {
	for true {
		GetAccessToken()
		time.Sleep(30 * time.Minute)

	}
}

func WxGetUserInfo(urlType string, code string) (openId string, unionId string, err error) {
	var url string
	if urlType == "A" {
		url = "https://api.weixin.qq.com/sns/jscode2session?appid=" + conf.Config.WxAppid + "&secret=" + conf.Config.WxSecret + "&js_code=" + code + "&grant_type=authorization_code"
	} else {
		url = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + conf.Config.WxAppid + "&secret=" + conf.Config.WxSecret + "&code=" + code + "&grant_type=authorization_code"
	}
	form, err := util.HTTPGet(url, nil, nil, nil, "")
	fmt.Println(string(form.Body), err)
	errorCode := gojsonq.New().FromString(string(form.Body)).Find("errcode")
	if errorCode != 0 && errorCode != nil {
		logKit.Log.Println("微信调用获取用户信息失败:" + string(form.Body))
		return "", "", errors.New(string(form.Body))
	}
	openId = util.ToString(gojsonq.New().FromString(string(form.Body)).Find("openid"))
	unionId = util.ToString(gojsonq.New().FromString(string(form.Body)).Find("unionid"))
	return
}
