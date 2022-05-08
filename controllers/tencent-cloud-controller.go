package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"net/http"
	"qiniu/service"
)

func TencentUpdateCDN(ctx *gin.Context) {
	// 更新cdn数据
	result := updateCDN()
	ctx.JSON(http.StatusOK, gin.H{"cdn result": result})
}

func updateCDN() bool {
	secret := service.NewTencentUserInfo()
	credential := common.NewCredential(
		secret.AccessKey,
		secret.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	client, _ := cdn.NewClient(credential, "", cpf)

	request := cdn.NewPurgePathCacheRequest()
	request.Paths = common.StringPtrs(secret.Host)
	request.FlushType = common.StringPtr("flush")

	response, err := client.PurgePathCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	if err != nil {
		//panic(err)
		return false
	}
	fmt.Printf("%s", response.ToJsonString())
	return true
}
