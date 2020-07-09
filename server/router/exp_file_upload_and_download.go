package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.FileUploadController)).Init(func(controller *v1.FileUploadController) {
		f := SpringBoot.Route("/fileUploadAndDownload")

		f.POST("/upload", SpringGin.Gin(controller.UploadFile))
		f.POST("/getFileList", SpringGin.Gin(controller.GetFileList))
		f.POST("/deleteFile", SpringGin.Gin(controller.DeleteFile))
		f.POST("/breakpointContinue", SpringGin.Gin(controller.BreakpointContinue))
		f.GET("/findFile", SpringGin.Gin(controller.FindFile))
		f.POST("/breakpointContinueFinish", SpringGin.Gin(controller.BreakpointContinueFinish))
		f.POST("/removeChunk", SpringGin.Gin(controller.RemoveChunk))
	})

}
