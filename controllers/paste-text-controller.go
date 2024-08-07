package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
	"strconv"
	"strings"
)

//临时粘贴板
func CopyTextIndex(ctx *gin.Context) {
	//上传　前台
	ctx.HTML(http.StatusOK, "paste.html", nil)
}

func HandlerCopyContent(ctx *gin.Context) {
	// 上传文本后台
	result := ctx.PostForm("result")
	result = strings.Trim(result, "")
	ret := service.UpdateContent(result)
	if ret {
		ctx.JSON(http.StatusOK, gin.H{"result": "passed", "code": 0})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"result": "failed", "code": 1})
	}
}

//临时粘贴板
func GetHistoryText(ctx *gin.Context) {
	//　获取历史记录
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "listtext.html", nil)
	} else {
		pwd := ctx.PostForm("password")
		count := ctx.PostForm("count")
		if pwd == service.AuthPassword {
			count_int, err := strconv.Atoi(count)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ret_data": nil,
				})
				return
			}
			ret := service.GetTextHistory(count_int)
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": ret,
			})

		} else {
			fmt.Println("error")
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": nil,
			})
		}
	}
}
