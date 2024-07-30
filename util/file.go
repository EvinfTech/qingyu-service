package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

var photoInt = 0

func SavePhoto(c *gin.Context, savePath string) {
	saveOld := savePath
	savePath = "static/" + savePath + "/"
	file, err := c.FormFile("file")
	customName, _ := c.GetPostForm("custom_name")
	fmt.Println(customName)
	if err != nil {
		ResponseKit(c, 10032, "上传发生错误:"+err.Error(), nil)
		return
	}
	fileMD5 := GetFileMD5(file)
	PathExists(savePath)
	fileName := file.Filename
	filePath := ""
	photoResponseName := ""
	if customName == "1" {
		filePath = savePath + fileName
		photoResponseName = fileName

	} else {
		fileNameWithSuffix := path.Base(fileName)
		fileType := path.Ext(fileNameWithSuffix)
		//看需求是否前面增加路径
		//filePath = savePath + fileMD5 + fileType
		filePath = fileMD5 + fileType
		filePath = savePath + ToString(Now().Unix()) + ToString(photoInt) + fileType
		photoResponseName = ToString(Now().Unix()) + ToString(photoInt) + fileType

	}

	//PublicInt += 1
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		ResponseKit(c, 10032, "文件保存发生错误:"+err.Error(), nil)
		return
	}
	ResponseKit(c, http.StatusOK, "上传成功", saveOld+"/"+photoResponseName)
}
func GetFileMD5(file *multipart.FileHeader) string {
	open, _ := file.Open()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, open); err != nil {
		fmt.Println("Copy", err)
		return ""
	}
	has := md5hash.Sum(nil)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	countSplit := strings.Split(path, "/")
	countSplit = countSplit[:len(countSplit)-1]
	path = ""
	for index, _ := range countSplit {
		path += countSplit[index] + "/"
	}
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	os.MkdirAll(path, 777)
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func FileInit() {
	PathExists("static/avatar/")
}
