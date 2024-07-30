package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	logKit "qiyu/logger"
	"time"
)

const (
	SECRETKEY = "alib2aba+=yu4sishid&adao6xi(xihaha" //私钥
)

type CustomClaims struct {
	UserOuid string
	jwt.StandardClaims
}
type ManagerCustomClaims struct {
	Id        int
	UserName  string
	StoreId   int
	StoreName string
	jwt.StandardClaims
}

func GetToken(ouid string) string {
	maxAge := 60 * 60 * 24 * 30
	//创建jwt
	customClaims := &CustomClaims{
		UserOuid: ouid, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		logKit.Log.Println("token生成发生错误:" + err.Error())
		return ""
	}
	return tokenString
}

func GetManagerToken(name string, storeId int, id int, storeName string) string {
	maxAge := 60 * 60 * 24 * 30
	//创建jwt
	customClaims := &ManagerCustomClaims{
		UserName:  name, //用户id
		StoreId:   storeId,
		Id:        id,
		StoreName: storeName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		logKit.Log.Println("token生成发生错误:" + err.Error())
		return ""
	}
	return tokenString
}

// 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func ParseManagerToken(tokenString string) (*ManagerCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ManagerCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*ManagerCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
func GetLoginUser(c *gin.Context) CustomClaims {
	token := c.GetHeader("Authorization")
	parseToken, _ := ParseToken(token[7:])
	return *parseToken
}
