package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.FileUploadController)).Init(func(c *v1.FileUploadController) {

		r := SpringBoot.Route("/fileUploadAndDownload")

		r.PostMapping("/upload", c.UploadFile)
		r.PostMapping("/getFileList", c.GetFileList)
		r.PostMapping("/deleteFile", c.DeleteFile)
		r.PostMapping("/breakpointContinue", c.BreakpointContinue)
		r.GetMapping("/findFile", c.FindFile)
		r.PostMapping("/breakpointContinueFinish", c.BreakpointContinueFinish)
		r.PostMapping("/removeChunk", c.RemoveChunk)
	})
}
