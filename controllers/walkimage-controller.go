package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
	"strconv"
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

	ctx.HTML(http.StatusOK, "totalImages.html", gin.H{"TotalImage": imageList[:20]})

}

func NextImagePage(ctx *gin.Context) {

	page := ctx.PostForm("page")
	page_int, _ := strconv.Atoi(page)
	imageList := service.Walkaround()
	fmt.Println("len of list ", len(imageList))
	//fmt.Printf("%v", imageList)
	ret_image_list := imageList[20*page_int : 20*(page_int+1)]
	fmt.Println("Next page")
	fmt.Println(ret_image_list)
	ctx.JSON(http.StatusOK, gin.H{"TotalImage": ret_image_list})
}
