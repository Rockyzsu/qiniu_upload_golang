/*
 * #Date: 2021-11-23 01:58:55
 * #Author: Rocky Chen
 * #LastEditors: Rocky Chen
 * #LastEditTime: 2021-11-23 02:01:29
 * #FilePath: \qiniu_web_gin\test.go
 * #Description:
 */
package main

import (
	"fmt"
	"log"
	"qiniu/service"
)

func main() {
	var location="C:\\git\\qiniu_web_gin\\upload\\20211108130019.png"
	var source="webupload"
	url,err:=service.UploadImg(location,source)
	if err!=nil{
		log.Fatal("Error")
	}else{
		fmt.Println(url)
	}

}
