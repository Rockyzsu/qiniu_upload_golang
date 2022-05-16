package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type CloudHost struct {
	Host              string `json:"host"`
	BAIDU_URL         string `json:"BAIDU_URL"`
	Wp_host           string `json:"wp_Host"`
	Wp_Baidu_host_url string `json:"wp_Baidu_host_url"`
}

func JsonParseHost(filename string) (CloudHost, error) {
	v := CloudHost{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return v, err
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Println(err)
		return v, err
	}
	return v, nil
}

func GetConfigURLFmt() CloudHost {

	host, err := JsonParseHost("service/web.json")
	if err != nil {
		panic("读取service/web.json出错，请创建web.json配置文件")
	}
	return host
}
