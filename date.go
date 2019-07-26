package gocore

import (
	"strconv"
	"strings"
)

func StrToSecond(time string) (num int) {

	times := strings.Split(time, ":")
	num = 0

	if len(times) == 3 {
		num1, _ := strconv.Atoi(times[0])
		num2, _ := strconv.Atoi(times[1])
		num3, _ := strconv.Atoi(times[2])
		num = num1*60*60 + num2*60 + num3
	} else {
		num1, _ := strconv.Atoi(times[0])
		num2, _ := strconv.Atoi(times[1])
		num = num1*60 + num2
	}
	return num
}
