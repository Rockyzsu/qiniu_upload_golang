package webmaster

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
	"strings"
)

//百度站长 资源提交

func ParseBD(html string) []string {
	bodyIO := bytes.NewReader([]byte(html))
	doc, err := libxml2.ParseHTMLReader(bodyIO)
	if err != nil {
		return []string{}
	}
	defer doc.Free()

	// 这里写xpath 逻辑 返回url列表
	nodes := xpath.NodeList(doc.Find(`//div[@class="aw-common-list"]/div[@class="aw-item article"]`))

	var articles []string
	//fmt.Println("articles string array")
	//fmt.Println(articles)
	//fmt.Println(len(nodes))
	for i := 0; i < len(nodes); i++ {
		fmt.Println(xpath.String(nodes[i].Find(`.//div/h4/a/@href`)))
		urlDetail := xpath.String(nodes[i].Find(`.//div/h4/a/@href`))

		articles = append(articles, urlDetail)
	}
	//fmt.Println("Articles ")
	//fmt.Println(articles)
	//fmt.Println("Len of article")
	//fmt.Println(len(articles))
	return articles
}

func GetPushUrlsBD(page int) []string {
	URL := fmt.Sprintf(HOST_CONFIG.Host, page)
	content := get(URL)
	return ParseBD(content)
}

func PushBaidu(urlList []string, status int) (bool, int) {
	client := resty.New()
	var last_remain int
	var BD_URL string
	if status == 1 {
		BD_URL = HOST_CONFIG.BAIDU_URL
	} else if status == 2 {
		BD_URL = strings.Replace(HOST_CONFIG.BAIDU_URL, "site=30daydo.com", "site=www.30daydo.com", -1)
	} else {
		fmt.Println("参数出错")
		return false, 0
	}
	fmt.Println(BD_URL)
	for _, url := range urlList {
		var res *Response
		fmt.Println(url)
		if status == 2 {
			url = strings.Replace(url, "http://", "http://www.", -1)
		}
		fmt.Println(url)
		fmt.Println(" ")
		_, err := client.R().
			SetHeader("Content-Type", "text/plain").
			SetHeader("Host", "data.zz.baidu.com").
			SetHeader("User-Agent", "curl/7.12.1").
			SetBody(url).
			SetResult(&res).
			Post(BD_URL)

		if err != nil {
			fmt.Println(err)
			return false, 0
		}
		last_remain = res.Remain

	}
	fmt.Println(last_remain)
	return true, last_remain
}

func PushProcess(totalPage int, value int) (bool, int) {
	// 扫描页数
	result := true
	_remain := 0
	for i := 1; i <= totalPage; i++ {
		urls := GetPushUrlsBD(i)
		status, remain := PushBaidu(urls, value)
		_remain = remain
		if status {
		} else {
			result = false
		}
	}
	return result, _remain
}
