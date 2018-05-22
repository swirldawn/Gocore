package gocore

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Md5(share string) string {

	h := md5.New()

	h.Write([]byte(share))

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果

}

func String2Byte(stringSlice []string) []byte {

	stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null

	return []byte(stringByte)
}

func GetPage(total string, size string, page string) map[string]interface{} {

	newTotal, _ := strconv.Atoi(total)
	newSize, _ := strconv.Atoi(size)
	newPage, _ := strconv.Atoi(page)

	if newSize > 40 {
		newSize = 40
	}

	var last_page int = int(math.Ceil(float64(newTotal) / float64(newSize)))

	var offset int = (newPage - 1) * newSize

	paginatorMap := make(map[string]interface{})

	paginatorMap["total"] = newTotal
	paginatorMap["per_page"] = newSize
	paginatorMap["last_page"] = last_page
	paginatorMap["current_page"] = newPage
	paginatorMap["offset"] = offset

	return paginatorMap

}

func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}

/**
 * 获取一个文件的md5值 路径为绝对路径
 */

func FileMd5(file_dir string) string {

	value := ""

	file, inerr := os.Open(file_dir)

	if inerr == nil {

		md5h := md5.New()
		io.Copy(md5h, file)

		value = hex.EncodeToString(md5h.Sum(nil)) // 输出加密结果
		file.Close()
	}

	return value
}
