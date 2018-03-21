package models

import (
	"fmt"
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

func GetConfig(section string, option string) (value string) {
	conf := goconf.New()
	if err := conf.Parse("./.env"); err != nil {
		panic(err)
	}
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
