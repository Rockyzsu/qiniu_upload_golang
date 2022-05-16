package webmaster

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
)

// 获取wordpress的链接

func parseWP(html string) []string {
	bodyIO := bytes.NewReader([]byte(html))
	doc, err := libxml2.ParseHTMLReader(bodyIO)
	if err != nil {
		return []string{}
	}
	defer doc.Free()

	// 这里写xpath 逻辑 返回url列表
	nodes := xpath.NodeList(doc.Find(`//div[@class="content"]/article`))

	var articles []string
	fmt.Println(len(nodes))
	for i := 0; i < len(nodes); i++ {
		fmt.Println(xpath.String(nodes[i].Find(`.//header/h2/a/@href`)))
		urlDetail := xpath.String(nodes[i].Find(`.//header/h2/a/@href`))
		articles = append(articles, urlDetail)
	}

	return articles
}

func GetPushUrlsWP(page int) []string {
	// 只获取首页
	url := fmt.Sprintf(HOST_CONFIG.Wp_host, page)
	content := get(url)
	return parseWP(content)
}

func _PushWP(urlList []string) (bool, int) {
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
			Post(HOST_CONFIG.Wp_Baidu_host_url)

		if err != nil {
			fmt.Println(err)
			return false, 0
		}
		last_remain = res.Remain

	}
	fmt.Println(last_remain)
	return true, last_remain
}

func PushWP(totalPage int) (bool, int) {
	// 扫描页数
	//totalPage := 2
	_remain := 0
	result := true
	for i := 1; i <= totalPage; i++ {
		urls := GetPushUrlsWP(i)
		status, remain := _PushWP(urls)
		_remain = remain
		if status {
		} else {
			result = false
		}
	}
	return result, _remain
}
