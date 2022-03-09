package controllers

import (
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

func DeleteImageByBT(ctx *gin.Context) {
	// 按钮删除图片
	if ctx.Request.Method == "POST" {

		url := ctx.PostForm("identifier")
		//fmt.Println(url)
		var status int

		ret := service.DeleteImage(url)
		if ret {
			status = 1
		}
		service.DeleteImageRecord(url)
		ctx.JSON(http.StatusOK, gin.H{"data": status, "status": "great"})
		return
	}
}
