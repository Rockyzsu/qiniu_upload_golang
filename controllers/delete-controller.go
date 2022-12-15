package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
)

func DeleteImage(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "delete.html", nil)
		return
	}

	if ctx.Request.Method == "POST" {

		url := ctx.PostForm("url")
		ret := service.DeleteImage(url)
		var status int
		if ret {
			status = 1
		}
		service.DeleteImageRecord(url)
		ctx.JSON(http.StatusOK, gin.H{"data": status, "status": "great"})
		return
	}
}

func DeleteImageById(ctx *gin.Context) {
	// 按钮删除图片
	if ctx.Request.Method == "POST" {

		id := ctx.PostForm("id")
		fmt.Println(id)
		var status int

		url, ret := service.DeleteImageRecordById(id)
		if ret {
			status = 1
		}
		fmt.Println(url)
		//service.DeleteImageRecord(url)
		ret = service.DeleteImage(url)
		if ret {
			ctx.JSON(http.StatusOK, gin.H{"data": status, "status": 1})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": status, "status": 0})

		}
		return
	}
}
