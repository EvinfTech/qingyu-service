package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"time"
)

func Countersign(m map[string]interface{}) map[string]interface{} {
	unix := time.Now().Unix()

	var a []string
	for k, _ := range m {
		a = append(a, k)
	}
	sort.Strings(a)
	ss := ""
	for _, v := range a {
		if m[v] == nil || m[v] == "" {
			continue
		}
		ss += ToString(v) + ":" + ToString(m[v]) + "&"
	}
	fmt.Println("ss=" + ss)

	code := fmt.Sprintf("%x", md5.Sum([]byte(ToString(ss)+"yueqian")))
	fmt.Println("code=" + code)

	s := sha256.New()
	s.Write([]byte(code + ToString(unix)))
	res := hex.EncodeToString(s.Sum(nil))

	fmt.Println(ToString(res))
	return gin.H{"aes": res, "timestamp": unix}
}
