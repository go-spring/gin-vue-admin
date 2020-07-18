package controller

import (
	"fmt"
	"strings"

	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(FileUploadController)).Init(func(c *FileUploadController) {

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

type FileUploadController struct {
	ExaFileUploadDownloadService *service.ExaFileUploadDownloadService `autowire:""`
	ExaBreakpointContinueService *service.ExaBreakpointContinueService `autowire:""`
	UploadService                *utils.UploadService                  `autowire:""`
}

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func (controller *FileUploadController) UploadFile(webCtx SpringWeb.WebContext) {
	noSave := webCtx.QueryParam("noSave") // TODO c.DefaultQuery("noSave", "0")
	if len(noSave) == 0 {
		noSave = "0"
	}
	header, err := webCtx.FormFile("file")
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), webCtx)
	} else {
		// 文件上传后拿到文件路径
		err, filePath, key := controller.UploadService.Upload(header)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), webCtx)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			var file model.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = header.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			if noSave == "0" {
				err = controller.ExaFileUploadDownloadService.Upload(file)
			}
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), webCtx)
			} else {
				response.OkDetailed(resp.ExaFileResponse{File: file}, "上传成功", webCtx)
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func (controller *FileUploadController) DeleteFile(webCtx SpringWeb.WebContext) {
	var file model.ExaFileUploadAndDownload
	_ = webCtx.Bind(&file)
	err, f := controller.ExaFileUploadDownloadService.FindFile(file.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), webCtx)
	} else {
		err = controller.UploadService.DeleteFile(f.Key)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), webCtx)

		} else {
			err = controller.ExaFileUploadDownloadService.DeleteFile(f)
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), webCtx)
			} else {
				response.OkWithMessage("删除成功", webCtx)
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取文件户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func (controller *FileUploadController) GetFileList(webCtx SpringWeb.WebContext) {
	var pageInfo request.PageInfo
	_ = webCtx.Bind(&pageInfo)
	err, list, total := controller.ExaFileUploadDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, webCtx)
	}
}
