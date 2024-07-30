package util

import (
	"github.com/gin-gonic/gin"
	logKit "qiyu/logger"
	"runtime"
	"strconv"
)

func ResponseKit(c *gin.Context, responseCode int, msg string, responseData interface{}) (code int, data interface{}) {
	if responseData == nil {
		responseData = gin.H{}
	}
	c.JSON(200, gin.H{
		"code": strconv.Itoa(responseCode),
		"msg":  msg,
		"data": responseData,
	})
	if responseCode != 200 {

		url := c.Request.URL.Path
		_, file, line, _ := runtime.Caller(1)
		logKit.Log.Println(msg + "   --调用位置:" + file + "   --行号:" + strconv.Itoa(line) + "----请求路径:" + url)
	}
	//if len(ToString(responseData)) < 300 {
	logKit.Log.Println("==请求路径==:" + c.Request.URL.Path + "=返回信息=:" + msg + "====返回数据====:" + ToString(responseData))
	//}

	return
}
