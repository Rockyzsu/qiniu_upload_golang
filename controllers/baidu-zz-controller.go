package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/webmaster"
	"strconv"
	"strings"
)

// 百度站长

func BaiduSite1(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {

		value := ctx.PostForm("value")
		fmt.Println("current is ", value)
		int_value, err := strconv.Atoi(value)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":  1,
				"res":   "",
				"count": 0,
			})
			return
		}
		res, count := webmaster.PushProcess(1, int_value)
		ctx.JSON(http.StatusOK, gin.H{
			"code":  0,
			"res":   res,
			"count": count,
		})
	}
}

func BaiduSite2(ctx *gin.Context) {

	res, count := webmaster.PushWP(1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"res":   res,
		"count": count,
	})
}

func BaiduSiteResource(ctx *gin.Context) {

	url := ctx.PostForm("url")
	url = strings.Trim(url, "")
	res, count := webmaster.PushBaiduQMT(url)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"res":   res,
		"count": count,
	})
}
