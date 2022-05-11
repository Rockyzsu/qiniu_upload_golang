package webmaster

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
)

//百度站长 资源提交

func Get() string {
	//url := fmt.Sprintf("https://api.github.com/users/%s", username)
	URL := "http://30daydo.com"

	client := resty.New()

	// 最简单的方法
	resp, err := client.R().
		SetHeader("Referer", "www.30daydo.com").
		SetHeader("User-Agent", "Go Spider from qiniu").
		Get(URL)

	if err != nil {
		return ""
	}
	//fmt.Println(resp)
	return resp.String()
}

func Parse(html string) []string {
	bodyIO := bytes.NewReader([]byte(html))
	doc, err := libxml2.ParseHTMLReader(bodyIO)
	if err != nil {
		return []string{}
	}
	defer doc.Free()

	// 这里写xpath 逻辑 返回url列表
	nodes := xpath.NodeList(doc.Find(`//div[@class="aw-common-list"]/div[@class="aw-item article"]`))

	//fmt.Printf("nodes type: %T,len: %d\n\n", nodes, len(nodes))
	var articles []string
	fmt.Println(len(nodes))
	for i := 0; i < len(nodes); i++ {
		fmt.Println(xpath.String(nodes[i].Find(`.//div/h4/a/@href`)))
		urlDetail := xpath.String(nodes[i].Find(`..//div/h4/a/@href`))
		articles = append(articles, urlDetail)
	}

	return articles
}

func Spider() {
	content := Get()
	Parse(content)
}
