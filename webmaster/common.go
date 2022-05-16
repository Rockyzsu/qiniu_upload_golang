package webmaster

import (
	"github.com/go-resty/resty/v2"
	"qiniu/service"
)

//公共函数

var HOST_CONFIG = service.GetConfigURLFmt()

func get(url string) string {

	client := resty.New()

	// 最简单的方法
	resp, err := client.R().
		SetHeader("User-Agent", "Go Spider").
		Get(url)

	if err != nil {
		return ""
	}
	return resp.String()
}
