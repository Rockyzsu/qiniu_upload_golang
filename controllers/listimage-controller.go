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
	//根据个数来查询
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "imageListV2.html", nil)
	} else {
		pwd := ctx.PostForm("password")
		count := ctx.PostForm("count")
		page := ctx.PostForm("page") //所有的数据
		var page_count, page_int int

		if pwd == service.AuthPassword {
			count_int, err := strconv.Atoi(count)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ret_data": nil,
					"status":   0,
				})
				return
			}
			total_count := service.TotalCount()
			if page == "" {
				page_count = total_count
				page_int = 1

			} else {
				page_int, _ = strconv.Atoi(page)
				page_count = int((float64(total_count)/float64(count_int) - float64(page_int) + 1) * float64(count_int))
			}

			if page_count <= 0 {
				ctx.JSON(http.StatusOK, gin.H{
					//"ret_data":     reversed_hist,
					"status": -1,
					//"total_count":  total_count,
					//"current_page": page_int,
				})
				return
			}

			hist := service.GetImageRecordByCount(page_count, count_int)
			var reversed_hist service.HIST
			for i := len(hist) - 1; i >= 0; i-- {
				reversed_hist = append(reversed_hist, hist[i])
			}
			ctx.JSON(http.StatusOK, gin.H{
				"ret_data":     reversed_hist,
				"status":       1,
				"total_count":  total_count,
				"current_page": page_int,
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
