package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.FileUploadController)).Init(func(c *v1.FileUploadController) {

		r := SpringBoot.Route("/fileUploadAndDownload")

		r.POST("/upload", SpringGin.Gin(c.UploadFile))
		r.POST("/getFileList", SpringGin.Gin(c.GetFileList))
		r.POST("/deleteFile", SpringGin.Gin(c.DeleteFile))
		r.POST("/breakpointContinue", SpringGin.Gin(c.BreakpointContinue))
		r.GET("/findFile", SpringGin.Gin(c.FindFile))
		r.POST("/breakpointContinueFinish", SpringGin.Gin(c.BreakpointContinueFinish))
		r.POST("/removeChunk", SpringGin.Gin(c.RemoveChunk))
	})
}
