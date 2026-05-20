package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
	"strconv"
)

const PAGENUM = 20

func WalkImages(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "imageList.html", nil)
	} else {
		pwd := ctx.PostForm("password")
		if pwd == service.AuthPassword {
			page := ctx.PostForm("page")
			page_int, _ := strconv.Atoi(page)
			imageList := service.Walkaround()
			ret_image_list := imageList[PAGENUM*page_int : PAGENUM*(page_int+1)]
			ctx.JSON(http.StatusOK, gin.H{"TotalImage": ret_image_list})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": nil,
				"status":   0,
			})
		}
	}
}
