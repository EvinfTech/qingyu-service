package kit

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var lock sync.Mutex

func init() {
	exists, err := PathExists("../static")
	fmt.Println(exists, err)

}

func WirteFile(wireteString string, filename string) {
	lock.Lock()
	var f *os.File
	var err1 error
	if CheckFileIsExist(filename) { //如果文件存在
		err1 := os.Remove(filename)
		if err1 != nil {
			fmt.Println(err1.Error())
		}
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件存在")
	} else {
		PathExists(filename)
		f, err1 = os.Create(filename) //创建文件
		fmt.Println(err1)
		fmt.Println("文件不存在")
	}
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("写入 %d 个字节n", n)
	lock.Unlock()
}

// 判断文件是否存在，存在返回true
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
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
func Copy(from, to string) error {
	f, e := os.Stat(from)
	if e != nil {
		return e
	}
	if f.IsDir() {
		//from是文件夹，那么定义to也是文件夹
		if list, e := ioutil.ReadDir(from); e == nil {
			for _, item := range list {
				if e = Copy(filepath.Join(from, item.Name()), filepath.Join(to, item.Name())); e != nil {
					return e
				}
			}
		}
	} else {
		//from是文件，那么创建to的文件夹
		p := filepath.Dir(to)
		if _, e = os.Stat(p); e != nil {
			if e = os.MkdirAll(p, 0777); e != nil {
				return e
			}
		}
		//读取源文件
		file, e := os.Open(from)
		if e != nil {
			return e
		}
		defer file.Close()
		bufReader := bufio.NewReader(file)
		// 创建一个文件用于保存
		out, e := os.Create(to)
		if e != nil {
			return e
		}
		defer out.Close()
		// 然后将文件流和文件流对接起来
		_, e = io.Copy(out, bufReader)
	}
	return e
}
func GetFileModTime(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return int(time.Now().Unix())
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return int(time.Now().Unix())
	}
	return int(fi.ModTime().Unix())
}
