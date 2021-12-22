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
	"log"
	"qiniu/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	// 配置加载模板路径
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	//  第一个参数是url的路径,第二个是本地的路径
	router.Static("/upload", "upload")

	// 将文件映射为url
	router.StaticFile("/check", "./upload/img.jpg")

	// 渲染模板
	router.GET("/", controllers.Index)
	router.POST("/", controllers.UploadPage)
	router.GET("/list", controllers.ListHistory)
	router.GET("/copy", controllers.CopyTextIndex)
	router.POST("/copycontent", controllers.HandlerCopyContent)
	router.GET("/paste", controllers.GetHistoryText)

	log.Fatalln(router.Run(":8080"))
}
