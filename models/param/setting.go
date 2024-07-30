package param

type SettingUpdateMessage struct {
	AlibabaCloudAccessKeyId     string `json:"alibaba_cloud_access_key_id" gorm:"column:alibaba_cloud_access_key_id;type:varchar(256);not null"`
	AlibabaCloudAccessKeySecret string `json:"alibaba_cloud_access_key_secret" gorm:"column:alibaba_cloud_access_key_secret;type:varchar(256);not null"`
	SignName                    string `json:"sign_name" gorm:"column:sign_name;type:varchar(256);not null"`
}

type SettingUpdateWx struct {
	WxAppid  string `json:"wx_appid" gorm:"column:wx_appid;type:varchar(256);not null"`
	WxSecret string `json:"wx_secret" gorm:"column:wx_secret;type:varchar(256);not null"`
}
