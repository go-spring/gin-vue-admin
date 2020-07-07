package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {

	f := SpringBoot.Route("/fileUploadAndDownload")

	f.POST("/upload", SpringGin.Gin(v1.UploadFile))
	f.POST("/getFileList", SpringGin.Gin(v1.GetFileList))
	f.POST("/deleteFile", SpringGin.Gin(v1.DeleteFile))
	f.POST("/breakpointContinue", SpringGin.Gin(v1.BreakpointContinue))
	f.GET("/findFile", SpringGin.Gin(v1.FindFile))
	f.POST("/breakpointContinueFinish", SpringGin.Gin(v1.BreakpointContinueFinish))
	f.POST("/removeChunk", SpringGin.Gin(v1.RemoveChunk))

	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload")
	// .Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		FileUploadAndDownloadGroup.POST("/upload", v1.UploadFile)                                 // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", v1.GetFileList)                           // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", v1.DeleteFile)                             // 删除指定文件
		FileUploadAndDownloadGroup.POST("/breakpointContinue", v1.BreakpointContinue)             // 断点续传
		FileUploadAndDownloadGroup.GET("/findFile", v1.FindFile)                                  // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		FileUploadAndDownloadGroup.POST("/removeChunk", v1.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
