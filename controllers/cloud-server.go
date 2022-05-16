package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerManager(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "cloud.html", nil)
}
