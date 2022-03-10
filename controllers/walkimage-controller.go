package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
	"strconv"
)

const PAGENUM = 20

func WalkImages(ctx *gin.Context) {
	//获取普通文本
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	if ctx.Request.Method == "GET" {
		imageList := service.Walkaround()
		total_count := len(imageList) / PAGENUM
		ctx.HTML(http.StatusOK, "totalImages.html", gin.H{"TotalImage": imageList[:PAGENUM], "length": total_count})

	} else {
		page := ctx.PostForm("page")
		page_int, _ := strconv.Atoi(page)
		imageList := service.Walkaround()
		ret_image_list := imageList[PAGENUM*page_int : PAGENUM*(page_int+1)]
		ctx.JSON(http.StatusOK, gin.H{"TotalImage": ret_image_list})
	}

}
