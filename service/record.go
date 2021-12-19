package service

import (
	"fmt"
	"log"
	"qiniu/model"
	"time"
)

func UpdateRecord(url string) bool {

	cursor, err := DB.Prepare("insert into tb_image_upload(url,updated)value (?,?)")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer cursor.Close()
	current := time.Now().Format("2006-01-01 05:04:05 PM")
	result, err := cursor.Exec(url, current)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if id, err := result.LastInsertId(); err == nil {
		fmt.Println("insert id ", id)
		return true
	} else {
		return false
	}
}

type HIST []model.History

func GetRecord() HIST {

	rows, err := DB.Query("select id,url,updated from tb_image_upload order by id desc limit 50")
	if err != nil {
		fmt.Println("查询失败")
		log.Fatal(err)
	}

	var hist HIST

	for rows.Next() {

		var h model.History
		err = rows.Scan(&h.Id, &h.Url, &h.Updated)
		if err != nil {
			fmt.Println("读取数据出错")
		}
		hist = append(hist, h)

	}
	return hist
}
