package util

import (
	"github.com/skip2/go-qrcode"
	"log"
)

var publicInt = 0

func QrCreat(s string) string {
	// 生成二维码
	Now().Unix()
	err := qrcode.WriteFile(s, qrcode.Medium, 256, "./static/qr/"+ToString(Now().Unix())+ToString(publicInt)+".png")
	if err != nil {
		log.Fatal(err)
	}
	return "/qr/" + ToString(Now().Unix()) + ToString(publicInt) + ".png"
}
