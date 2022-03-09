package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"qiniu/service"
)

//上传页面 index
func UploadPage(ctx *gin.Context) {
	//获取普通文本
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	file, err := ctx.FormFile("file")
	//fmt.Printf("%v\n",file)
	if file == nil {
		log.Println("文件没有选中")
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "请选择文件",
			"url":     "",
		})
		return
	}
	if err != nil {
		log.Println("获取数据失败")
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
		status := service.InsertImageRecord(urlpath)
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

}
