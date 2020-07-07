package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	f := SpringBoot.Route("/fileUploadAndDownload")

	f.POST("/upload", SpringGin.Gin(v1.UploadFile))
	f.POST("/getFileList", SpringGin.Gin(v1.GetFileList))
	f.POST("/deleteFile", SpringGin.Gin(v1.DeleteFile))
	f.POST("/breakpointContinue", SpringGin.Gin(v1.BreakpointContinue))
	f.GET("/findFile", SpringGin.Gin(v1.FindFile))
	f.POST("/breakpointContinueFinish", SpringGin.Gin(v1.BreakpointContinueFinish))
	f.POST("/removeChunk", SpringGin.Gin(v1.RemoveChunk))
}
