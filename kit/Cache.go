package kit

import (
	"github.com/patrickmn/go-cache"
	"qiyu/util"
	"time"
)

var Phone *cache.Cache

func init() {
	// 设置超时时间和清理时间
	Phone = cache.New(5*time.Minute, 1*time.Minute)
}
func UpdatePhoneCache(phone string, code string) {
	Phone.Set(phone, code, cache.DefaultExpiration)
}

// 如果超时返回false 没超时返回true
func GetPhoneCache(phone string) (string, bool) {
	code, b := Phone.Get(phone)
	if code == nil {
		return "", b
	}
	return util.ToString(code), b
}
