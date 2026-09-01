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
	"net/http"
	"qiniu/router"
	"qiniu/service"
	"time"
)

func main() {
	routerObj := router.Router()
	srv := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", service.Port),
		Handler:           routerObj,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
		IdleTimeout:       120 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())
}
