package gocore

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func SaveFile(url string, fileName string, filePath string) (n int64, err error) {
	isExists := PathExists(filePath)
	if isExists == false {
		//创建多级目录和设置权限
		err = os.MkdirAll(filePath, 0777)
		return 0, err

	}

	// os.Chdir(filePath)

	out, err := os.Create(fileName)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))

	os.Rename(fileName, filePath+"/"+fileName)

	return
}

func SaveMp4(url string, fileName string, filePath string) (realFileName string, err error) {

	if fileName == "" {
		fileName = GetUniqFileName()
	}

	realFileName = fileName + ".mp4"

	SaveFile(url, realFileName, filePath)

	return realFileName, err
}

func GetUniqFileName() (name string) {
	name = strconv.FormatInt(time.Now().Unix(), 10) + time.Now().Format("20060102150405")
	return name
}

func SaveImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return

}
