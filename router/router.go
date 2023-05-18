package router

import (
	"github.com/gin-gonic/gin"
	"qiniu/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()
	// 配置加载模板路径

	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLGlob("templates/css/*")
	//router.LoadHTMLGlob("templates/js/*")
	//router.LoadHTMLGlob("templates/img/*")
	//router.LoadHTMLGlob("templates/font-aweome-4.6.3/*")

	router.Static("/static", "static")

	//  第一个参数是url的路径,第二个是本地的路径
	router.Static("/upload", "upload")

	// 将文件映射为url
	router.StaticFile("/check", "./upload/img.jpg")

	// 图片
	router.GET("/", controllers.Index)
	router.POST("/", controllers.UploadPage) //上传

	router.GET("/list", controllers.ListImageHistory)  //显示url
	router.POST("/list", controllers.ListImageHistory) //显示url

	router.GET("/l", controllers.ListImageHistory)  //快捷方式
	router.POST("/l", controllers.ListImageHistory) //快捷方式

	router.GET("/l2", controllers.ListImageHistoryV2)  //快捷方式
	router.POST("/l2", controllers.ListImageHistoryV2) //快捷方式

	router.GET("/qs", controllers.ListImageHistoryQS)  //券商
	router.POST("/qs", controllers.ListImageHistoryQS) //券商

	router.GET("/jump", controllers.ListImageHistoryPageV2)  //快捷方式
	router.POST("/jump", controllers.ListImageHistoryPageV2) //快捷方式

	router.GET("/ll", controllers.ListHistorys) //快捷方式
	// 文本
	router.GET("/copy", controllers.CopyTextIndex)
	router.GET("/c", controllers.CopyTextIndex)
	router.POST("/copycontent", controllers.HandlerCopyContent)
	router.GET("/paste", controllers.GetHistoryText)
	router.POST("/paste", controllers.GetHistoryText)
	router.GET("/p", controllers.GetHistoryText)
	router.POST("/p", controllers.GetHistoryText)

	//删除 需要权限
	router.GET("/d", controllers.DeleteImage) // 删除
	router.POST("/delimage", controllers.DeleteImage)
	router.POST("/image/remove_id", controllers.DeleteImageById)

	// 获取所有图片
	router.GET("/w", controllers.WalkImages)
	router.GET("/walk", controllers.WalkImages)
	router.POST("/image/next", controllers.WalkImages)

	//router.GET("/cloud", controllers.TencentUpdateCDN)

	router.GET("/cloud", controllers.ServerManager)
	router.POST("/update-site1", controllers.BaiduSite1)
	router.POST("/update-site2", controllers.BaiduSite2)
	router.POST("/update-cdn", controllers.TencentUpdateCDN)

	// 测试用
	router.GET("/node", controllers.DynamicAddNode)
	router.GET("/count", controllers.CountInterface)

	return router
}
