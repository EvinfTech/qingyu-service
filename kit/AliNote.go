package kit

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	teaUtil "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"qiyu/conf"
	"strings"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

/**
* 使用STS鉴权方式初始化账号Client，推荐此方式。
* @param accessKeyId
* @param accessKeySecret
* @param securityToken
* @return Client
* @throws Exception
 */
func CreateClientWithSTS(accessKeyId *string, accessKeySecret *string, securityToken *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
		// 必填，您的 Security Token
		SecurityToken: securityToken,
		// 必填，表明使用 STS 方式
		Type: tea.String("sts"),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func PhoneCode(phone string, code string) (_err error) {
	client, _err := CreateClient(tea.String(conf.Config.AlibabaCloudAccessKeyId), tea.String(conf.Config.AlibabaCloudAccessKeySecret))
	if _err != nil {
		return _err
	}
	//todo 这里需要将TemplateCode修改为自己的短信摸板
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(conf.Config.SignName),
		TemplateCode:  tea.String("SMS_465319325"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{'code':'" + code + "'}"),
	}
	runtime := &teaUtil.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		fmt.Println("res:", res)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println("错误信息", tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = teaUtil.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
