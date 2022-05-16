package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/webmaster"
)

// 百度站长

func BaiduSite1(ctx *gin.Context) {

	res, count := webmaster.PushProcess(1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"res":   res,
		"count": count,
	})
}

func BaiduSite2(ctx *gin.Context) {

	res, count := webmaster.PushWP(1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"res":   res,
		"count": count,
	})
}
