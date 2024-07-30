package kit

//
//import (
//	"bytes"
//	"github.com/goccy/go-json"
//	"github.com/robfig/cron"
//	"io"
//	"io/ioutil"
//	"log"
//	"mime/multipart"
//	"net/http"
//	"net/url"
//	"os"
//	"os/exec"
//	"path/filepath"
//	logKit "sport/log"
//	"sport/util"
//	"strings"
//	"time"
//)
//
////数据库自动备份
//func DbBackup() {
//	c := cron.New()
//	c.AddFunc("0 0 1 * * ?", func() {
//		format := util.Now().Format("2006-01-02")
//		logKit.Log.Println("开始执行数据库自动备份")
//		//开始备份
//
//		cmd := exec.Command("mysqldump", "-u", "root", "-pwoshimima", "--databases", "sport3", "-r", format+".sql")
//		err := cmd.Start()
//		if err != nil {
//			logKit.Log.Println("------------------------------------------------自动备份失败")
//		}
//		//加一些延时，其实0.5秒就够，但是长一点
//		time.Sleep(10 * time.Second)
//		stat, err := os.Stat("./backups/" + format)
//		if stat == nil {
//			os.MkdirAll("./backups/"+format, 777)
//		} else {
//			if !stat.IsDir() {
//				os.MkdirAll("./backups/"+format, 777)
//			}
//		}
//		if err != nil {
//			log.Println("判断文件夹出现错误:", err)
//		}
//		util.Zip(format+".sql", "./backups/"+format+"/运动服务器"+format+"db.zip")
//		os.Remove(format + ".sql")
//
//		logKit.Log.Println("自动备份成功")
//	})
//	c.Start()
//}
//
//type UploadFile struct {
//	// 表单名称
//	Name string
//	// 文件全路径
//	Filepath string
//}
//
//// 请求客户端
//var httpClient = &http.Client{}
//
////下面是发送文件一次性的一些方法
//func PostFile(reqUrl string, reqParams map[string]string, files []UploadFile, headers map[string]string) string {
//	return post(reqUrl, reqParams, "multipart/form-data", files, headers)
//}
//
//func post(reqUrl string, reqParams map[string]string, contentType string, files []UploadFile, headers map[string]string) string {
//	requestBody, realContentType := getReader(reqParams, contentType, files)
//	httpRequest, _ := http.NewRequest("POST", reqUrl, requestBody)
//	// 添加请求头
//	httpRequest.Header.Add("Content-Type", realContentType)
//	if headers != nil {
//		for k, v := range headers {
//			httpRequest.Header.Add(k, v)
//		}
//	}
//	// 发送请求
//	resp, err := httpClient.Do(httpRequest)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	response, _ := ioutil.ReadAll(resp.Body)
//	return string(response)
//}
//
//func getReader(reqParams map[string]string, contentType string, files []UploadFile) (io.Reader, string) {
//	if strings.Index(contentType, "json") > -1 {
//		bytesData, _ := json.Marshal(reqParams)
//		return bytes.NewReader(bytesData), contentType
//	} else if files != nil {
//		body := &bytes.Buffer{}
//		// 文件写入 body
//		writer := multipart.NewWriter(body)
//		for _, uploadFile := range files {
//			file, err := os.Open(uploadFile.Filepath)
//			if err != nil {
//				panic(err)
//			}
//			part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.Filepath))
//			if err != nil {
//				panic(err)
//			}
//			_, err = io.Copy(part, file)
//			file.Close()
//		}
//		// 其他参数列表写入 body
//		for k, v := range reqParams {
//			if err := writer.WriteField(k, v); err != nil {
//				panic(err)
//			}
//		}
//		if err := writer.Close(); err != nil {
//			panic(err)
//		}
//		// 上传文件需要自己专用的contentType
//		return body, writer.FormDataContentType()
//	} else {
//		urlValues := url.Values{}
//		for key, val := range reqParams {
//			urlValues.Set(key, val)
//		}
//		reqBody := urlValues.Encode()
//		return strings.NewReader(reqBody), contentType
//	}
//}
