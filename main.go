/*
 * #Date: 2021-11-22 21:58:08
 * #Author: Rocky Chen
 * #LastEditors: Rocky Chen
 * #LastEditTime: 2021-11-23 01:29:20
 * #FilePath: \qiniu_web_gin\main.go
 * #Description: 七牛上传图片
 */

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"qiniu/router"
	"qiniu/service"
)

func main() {
	routerObj := router.Router()
	log.Fatalln(routerObj.Run(fmt.Sprintf("0.0.0.0:%d", service.Port)))
}
