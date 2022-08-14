package service

import (
	"fmt"
	"log"
	"qiniu/model"
	"time"
)

func InsertImageRecord(url string) bool {

	cursor, err := DB.Prepare("insert into tb_image_upload(url,updated,isDeleted)value (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer cursor.Close()
	current := time.Now().Format("2006-01-02 15:04:05")
	result, err := cursor.Exec(url, current, false)

	if err != nil {
		log.Println(err)
		return false
	}
	if id, err := result.LastInsertId(); err == nil {
		log.Println("insert id ", id)
		return true
	} else {
		return false
	}
}

type HIST []model.History

func GetImageRecord(count int) HIST {

	rows, err := DB.Query(fmt.Sprintf("select id,url,updated from tb_image_upload where isDeleted=0 order by id desc limit %d", count))
	if err != nil {
		log.Println("查询url失败")
		log.Fatal(err)
	}

	var hist HIST

	for rows.Next() {

		var h model.History
		err = rows.Scan(&h.Id, &h.Url, &h.Updated)
		if err != nil {
			log.Println("读取url数据出错")
		}
		hist = append(hist, h)

	}
	return hist
}

type TEXTList []model.ContentText

func GetTextHistory(count int) TEXTList {
	rows, err := DB.Query(fmt.Sprintf("select id,text,updated from tb_content order by id desc limit %d", count))
	if err != nil {
		log.Println("查询text失败")
	}

	var texts TEXTList
	for rows.Next() {
		var text model.ContentText
		err = rows.Scan(&text.Id, &text.Text, &text.Updated)
		if err != nil {
			log.Println("读取text失败")
		}

		texts = append(texts, text)
	}

	return texts

}

func UpdateContent(text string) bool {

	cursor, err := DB.Prepare("insert into tb_content(text,updated)value (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer cursor.Close()
	current := time.Now().Format("2006-01-01 15:04:05")
	result, err := cursor.Exec(text, current)
	if err != nil {
		log.Println(err)
		return false
	}
	if id, err := result.LastInsertId(); err == nil {
		log.Println("insert id ", id)
		return true
	} else {
		return false
	}
}

func DeleteImageRecord(url string) bool {
	//删除图片后把状态修改
	cursor, err := DB.Prepare("update tb_image_upload set isDeleted=true where url=?")
	if err != nil {
		log.Println(err)
		return false
	}
	defer cursor.Close()
	result, err := cursor.Exec(url)
	if err != nil {
		log.Println(err)
		return false
	}
	if id, err := result.LastInsertId(); err == nil {
		log.Println("insert id ", id)
		return true
	} else {
		return false
	}
}
