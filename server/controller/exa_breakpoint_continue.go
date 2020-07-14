package controller

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"gin-vue-admin/global/response"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/go-spring/go-spring-web/spring-web"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func (controller *FileUploadController) BreakpointContinue(webCtx SpringWeb.WebContext) {
	fileMd5 := webCtx.FormValue("fileMd5")
	fileName := webCtx.FormValue("fileName")
	chunkMd5 := webCtx.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(webCtx.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(webCtx.FormValue("chunkTotal"))
	FileHeader, err := webCtx.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), webCtx)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		response.FailWithMessage(err.Error(), webCtx)
		return
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if flag := utils.CheckMd5(cen, chunkMd5); !flag {
		return
	}
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		response.FailWithMessage(err.Error(), webCtx)
		return
	}
	err, pathc := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		response.FailWithMessage(err.Error(), webCtx)
		return
	}

	if err = service.CreateFileChunk(file.ID, pathc, chunkNumber); err != nil {
		response.FailWithMessage(err.Error(), webCtx)
		return
	}
	response.OkWithMessage("切片创建成功", webCtx)
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func (controller *FileUploadController) FindFile(webCtx SpringWeb.WebContext) {
	fileMd5 := webCtx.QueryParam("fileMd5")
	fileName := webCtx.QueryParam("fileName")
	chunkTotal, _ := strconv.Atoi(webCtx.QueryParam("chunkTotal"))
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		response.FailWithMessage("查找失败", webCtx)
	} else {
		response.OkWithData(resp.FileResponse{File: file}, webCtx)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func (controller *FileUploadController) BreakpointContinueFinish(webCtx SpringWeb.WebContext) {
	fileMd5 := webCtx.QueryParam("fileMd5")
	fileName := webCtx.QueryParam("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.FilePathResponse{FilePath: filePath}, fmt.Sprintf("文件创建失败：%v", err), webCtx)
	} else {
		response.OkDetailed(resp.FilePathResponse{FilePath: filePath}, "文件创建成功", webCtx)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func (controller *FileUploadController) RemoveChunk(webCtx SpringWeb.WebContext) {
	fileMd5 := webCtx.QueryParam("fileMd5")
	fileName := webCtx.QueryParam("fileName")
	filePath := webCtx.QueryParam("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.FilePathResponse{FilePath: filePath}, fmt.Sprintf("缓存切片删除失败：%v", err), webCtx)
	} else {
		response.OkDetailed(resp.FilePathResponse{FilePath: filePath}, "缓存切片删除成功", webCtx)
	}
}
