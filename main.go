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
	_ "github.com/go-sql-driver/mysql"
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
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/", func(ctx *gin.Context) {
		//获取普通文本
		// 获取文件(注意这个地方的file要和html模板中的name一致)
		file, err := ctx.FormFile("file")
		//fmt.Printf("%v\n",file)
		if file == nil {
			fmt.Println("文件没有选中")
			ctx.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "请选择文件",
				"url":     "",
			})
			return
		}
		if err != nil {
			fmt.Println("获取数据失败")
			ctx.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "获取数据失败",
				"url":     "",
			})
			return
		} else {
			//保存上传文件
			filePath := filepath.Join("./upload", file.Filename)
			err = ctx.SaveUploadedFile(file, filePath)
			if err != nil {
				fmt.Println("保存失败")
				ctx.JSON(http.StatusOK, gin.H{
					"code":    2,
					"message": "保存失败",
					"data":    "",
				})
				return
			}
			if service.CheckExist(file.Filename) {
				// 文件已经存在
				//fmt.Println("文件已经存在")
				ctx.JSON(http.StatusOK, gin.H{
					"code":    2,
					"message": "文件名已经存在，为避免覆盖，修改文件名",
					"data":    "",
				})
				return

			}
			urlpath, err := service.UploadImg(filePath, service.SOURCE)
			if err != nil {
				log.Fatal("Error")
			}
			//fmt.Println("上传成功")
			status := service.UpdateRecord(urlpath)
			if status {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    0,
					"message": "上传成功",
					"data":    urlpath,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "上传成功，写入记录出错",
					"data":    urlpath,
				})
			}
			return
		}

	})

	router.GET("/list", func(ctx *gin.Context) {
		hist := service.GetRecord()
		ctx.HTML(http.StatusOK, "ilist.html", gin.H{"history": hist})
	})

	router.GET("/copy", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "paste.html", nil)
	})

	router.Run(":80")
}
