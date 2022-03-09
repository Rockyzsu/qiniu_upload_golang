package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
)

func NextPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "nextpage.html", nil)
}

func WalkImages(ctx *gin.Context) {
	//获取普通文本
	// 获取文件(注意这个地方的file要和html模板中的name一致)

	imageList := service.Walkaround()
	fmt.Println("len of list ", len(imageList))
	//fmt.Printf("%v", imageList)
	ctx.HTML(http.StatusOK, "totalImages.html", gin.H{"TotalImage": imageList})

}
