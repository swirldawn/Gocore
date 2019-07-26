package gocore

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Terry-Mao/goconf"
)

type TestConfig struct {
	ID     int            `goconf:"core:id"`
	Col    string         `goconf:"core:col"`
	Ignore int            `goconf:"-"`
	Arr    []string       `goconf:"core:arr:,"`
	Test   time.Duration  `goconf:"core:t_1:time"`
	Buf    int            `goconf:"core:buf:memory"`
	M      map[int]string `goconf:"core:m:,"`
}

var conf *goconf.Config

func InitConfig() {
	conf = goconf.New()
	if err := conf.Parse("./.env"); err != nil {
		panic(err)
	}
}

func GetConfig(section string, option string) (value string) {

	core := conf.Get(section)

	if core == nil {
		fmt.Println("no core section")
	}
	value, err := core.String(option)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func GetSection(section string) (value []string) {

	core := conf.Get(section)

	for a := 0; a < 100; a++ {

		val, err := core.String(strconv.Itoa(a))

		if err != nil {
			break
		}

		value = append(value, val)
	}

	return value
}
