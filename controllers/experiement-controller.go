package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu/service"
)

func DynamicAddNode(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addNode.html", nil)
}

func CountInterface(ctx *gin.Context) {
	count := service.TotalCount()
	fmt.Print(count)
	ctx.JSON(http.StatusOK, gin.H{"count": count})
}
