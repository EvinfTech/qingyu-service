package util

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io/ioutil"
	logKit "qiyu/logger"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(bufio.NewReaderSize(c.Request.Body, 16384))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		if len(buf) < 500 {
			logKit.Log.Println("====收到请求====:" + c.Request.URL.Path + "请求ip:" + c.ClientIP() + "请求参数：" + string(buf))
		} else {
			logKit.Log.Println("====收到请求====:" + c.Request.URL.Path + "请求ip:" + c.ClientIP() + "请求参数有点多，就不展示了")

		}
		c.Next()
		paramMap := make(map[string]interface{})
		json.Unmarshal(buf, &paramMap)

		defer func() {
			if err := recover(); err != nil {
				logKit.Log.Printf("Panic info is: %v", err)
			}
		}()
	}
}
