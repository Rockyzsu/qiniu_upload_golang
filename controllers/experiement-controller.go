package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DynamicAddNode(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addNode.html", nil)
}
