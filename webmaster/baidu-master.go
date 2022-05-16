package webmaster

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
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
	fmt.Println(len(nodes))
	for i := 0; i < len(nodes); i++ {
		fmt.Println(xpath.String(nodes[i].Find(`.//div/h4/a/@href`)))
		urlDetail := xpath.String(nodes[i].Find(`..//div/h4/a/@href`))
		articles = append(articles, urlDetail)
	}

	return articles
}

func GetPushUrlsBD(page int) []string {
	URL := fmt.Sprintf(HOST_CONFIG.Host, page)
	content := get(URL)
	return ParseBD(content)
}

func PushBaidu(urlList []string) (bool, int) {
	client := resty.New()
	var last_remain int
	for _, url := range urlList {
		var res *Response
		_, err := client.R().
			SetHeader("Content-Type", "text/plain").
			SetHeader("Host", "data.zz.baidu.com").
			SetHeader("User-Agent", "curl/7.12.1").
			SetBody(url).
			SetResult(&res).
			Post(HOST_CONFIG.BAIDU_URL)

		if err != nil {
			fmt.Println(err)
			return false, 0
		}
		last_remain = res.Remain

	}
	fmt.Println(last_remain)
	return true, last_remain
}

func PushProcess(totalPage int) (bool, int) {
	// 扫描页数
	result := true
	_remain := 0
	for i := 1; i <= totalPage; i++ {
		urls := GetPushUrlsBD(i)
		status, remain := PushBaidu(urls)
		_remain = remain
		if status {
		} else {
			result = false
		}
	}
	return result, _remain
}
