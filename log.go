package gocore

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func Info(log string) {

	log = log + "\r\n"
	var filename = "/data/code/golong" + time.Now().Format("2006-01-02") + ".log"
	var f *os.File
	var err1 error
	fmt.Println(log)
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_WRONLY, 0666) //打开文件
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		check(err1)
		_, err := f.WriteAt([]byte(log), n)
		check(err)

	} else {
		f, err1 = os.Create(filename) //创建文件
		check(err1)
		_, err := f.Write([]byte(log))
		check(err)
	}

	f.Close()
}

func PutFile(fileName string, content string) {
	log := content + "\r\n"
	var f *os.File
	var err1 error
	fmt.Println(log)
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(fileName) { //如果文件存在
		f, err1 = os.OpenFile(fileName, os.O_WRONLY, 0666) //打开文件
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		check(err1)
		_, err := f.WriteAt([]byte(log), n)
		check(err)

	} else {
		f, err1 = os.Create(fileName) //创建文件
		check(err1)
		_, err := f.Write([]byte(log))
		check(err)
	}

	f.Close()
}

func ReadFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var result []string

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		result = append(result, line)
	}

	return result
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
