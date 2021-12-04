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
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"qiniu/service"
)

func main() {
	router := gin.Default()
	// 配置加载模板路径
	router.LoadHTMLGlob("templates/*")
	router.Static("./static", "static")
	// 渲染模板
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index_old.html", nil)
	})
	router.POST("/", func(ctx *gin.Context) {
		//获取普通文本
		// 获取文件(注意这个地方的file要和html模板中的name一致)
		file, err := ctx.FormFile("file")
		if err != nil {
			fmt.Println("获取数据失败")
			ctx.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "获取数据失败",
				"url":"",
			})
		} else {
			//保存上传文件
			filePath := filepath.Join("./upload", file.Filename)
			err = ctx.SaveUploadedFile(file, filePath)
			if err != nil {
				fmt.Println("保存失败")
			}
			var source = "webupload"
			urlpath, err := service.UploadImg(filePath, source)
			if err != nil {
				log.Fatal("Error")
			}

			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "success",
				"data":    urlpath,
			})
		}

	})


	router.Run(":80")
}
