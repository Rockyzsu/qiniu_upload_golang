package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
	"strings"
)

//上传记录
func ListHistory(ctx *gin.Context) {
	hist := service.GetImageRecord()
	ctx.HTML(http.StatusOK, "imageList.html", gin.H{"history": hist})
}

func ListHistorys(ctx *gin.Context) {
	hist := service.GetImageRecord()
	ctx.HTML(http.StatusOK, "images.html", gin.H{"history": hist})
}

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
	ret := service.GetTextHistory()
	ctx.HTML(http.StatusOK, "listtext.html", gin.H{"texts": ret})
}
