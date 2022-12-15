package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
	"strconv"
)

//上传记录
func ListImageHistory(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "imageList.html", nil)
	} else {
		pwd := ctx.PostForm("password")
		count := ctx.PostForm("count")
		if pwd == service.AuthPassword {
			count_int, err := strconv.Atoi(count)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ret_data": nil,
					"status":   0,
				})
				return
			}
			hist := service.GetImageRecord(count_int)

			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": hist,
				"status":   1,
			})

		} else {
			fmt.Println("error")
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": nil,
				"status":   0,
			})
		}

	}
}

func ListImageHistoryPageV2(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "imageListV2.html", nil)
	} else {
		pwd := ctx.PostForm("password")
		count := ctx.PostForm("count")
		page := ctx.PostForm("page")
		fmt.Println("page---")
		fmt.Println(page)
		var page_index int
		if pwd == service.AuthPassword {
			count_int, err := strconv.Atoi(count)
			if page == "" {
				page_index = service.TotalCount()
				fmt.Println("set page to ", page_index)
			} else {
				page_index, _ = strconv.Atoi(page)
			}
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ret_data": nil,
					"status":   0,
				})
				return
			}

			hist := service.GetImageRecordByIndex(page_index, count_int)
			fmt.Println(hist)
			total_count := service.TotalCount()
			//fmt.Println(total_count)
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data":    hist,
				"status":      1,
				"total_count": total_count,
			})

		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": nil,
				"status":   0,
			})
		}

	}
}

func ListImageHistoryV2(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "imageListV2.html", nil)
	} else {
		pwd := ctx.PostForm("password")
		count := ctx.PostForm("count")
		if pwd == service.AuthPassword {
			count_int, err := strconv.Atoi(count)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ret_data": nil,
					"status":   0,
				})
				return
			}
			hist := service.GetImageRecord(count_int)

			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": hist,
				"status":   1,
			})

		} else {
			fmt.Println("error")
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data": nil,
				"status":   0,
			})
		}

	}
}

func ListHistorys(ctx *gin.Context) {
	hist := service.GetImageRecord(50)
	ctx.HTML(http.StatusOK, "images.html", gin.H{"history": hist})
}
