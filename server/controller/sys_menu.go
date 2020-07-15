package controller

import (
	"fmt"

	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(MenuController)).Init(func(c *MenuController) {

		r := SpringBoot.Route("/menu",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/getMenu", c.GetMenu)
		r.PostMapping("/getMenuList", c.GetMenuList)
		r.PostMapping("/addBaseMenu", c.AddBaseMenu)
		r.PostMapping("/getBaseMenuTree", c.GetBaseMenuTree)
		r.PostMapping("/addMenuAuthority", c.AddMenuAuthority)
		r.PostMapping("/getMenuAuthority", c.GetMenuAuthority)
		r.PostMapping("/deleteBaseMenu", c.DeleteBaseMenu)
		r.PostMapping("/updateBaseMenu", c.UpdateBaseMenu)
		r.PostMapping("/getBaseMenuById", c.GetBaseMenuById)
	})
}

type MenuController struct {
	SysMenuService *service.SysMenuService `autowire:""`
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func (controller *MenuController) GetMenu(webCtx SpringWeb.WebContext) {
	claims := webCtx.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err, menus := controller.SysMenuService.GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysMenusResponse{Menus: menus}, webCtx)
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func (controller *MenuController) GetMenuList(webCtx SpringWeb.WebContext) {
	var pageInfo request.PageInfo
	_ = webCtx.Bind(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), webCtx)
		return
	}
	err, menuList, total := controller.SysMenuService.GetInfoList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, webCtx)
	}
}

// @Tags menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "新增菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addBaseMenu [post]
func (controller *MenuController) AddBaseMenu(webCtx SpringWeb.WebContext) {
	var menu model.SysBaseMenu
	_ = webCtx.Bind(&menu)
	MenuVerify := utils.Rules{
		"Path":      {utils.NotEmpty()},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0"), "ge=0"},
	}
	MenuVerifyErr := utils.Verify(menu, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), webCtx)
		return
	}
	MetaVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		response.FailWithMessage(MetaVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysMenuService.AddBaseMenu(menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("添加成功", webCtx)
	}
}

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getBaseMenuTree [post]
func (controller *MenuController) GetBaseMenuTree(webCtx SpringWeb.WebContext) {
	err, menus := controller.SysMenuService.GetBaseMenuTree()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysBaseMenusResponse{Menus: menus}, webCtx)
	}
}

// @Tags authorityAndMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/addMenuAuthority [post]
func (controller *MenuController) AddMenuAuthority(webCtx SpringWeb.WebContext) {
	var addMenuAuthorityInfo request.AddMenuAuthorityInfo
	_ = webCtx.Bind(&addMenuAuthorityInfo)
	MenuVerify := utils.Rules{
		"AuthorityId": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(addMenuAuthorityInfo, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysMenuService.AddMenuAuthority(addMenuAuthorityInfo.Menus, addMenuAuthorityInfo.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("添加成功", webCtx)
	}
}

// @Tags authorityAndMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityIdInfo true "增加menu和角色关联关系"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func (controller *MenuController) GetMenuAuthority(webCtx SpringWeb.WebContext) {
	var authorityIdInfo request.AuthorityIdInfo
	_ = webCtx.Bind(&authorityIdInfo)
	MenuVerify := utils.Rules{
		"AuthorityId": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(authorityIdInfo, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), webCtx)
		return
	}
	err, menus := controller.SysMenuService.GetMenuAuthority(authorityIdInfo.AuthorityId)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.SysMenusResponse{Menus: menus}, fmt.Sprintf("添加失败，%v", err), webCtx)
	} else {
		response.Result(response.SUCCESS, gin.H{"menus": menus}, "获取成功", webCtx)
	}
}

// @Tags menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/deleteBaseMenu [post]
func (controller *MenuController) DeleteBaseMenu(webCtx SpringWeb.WebContext) {
	var idInfo request.GetById
	_ = webCtx.Bind(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysMenuService.DeleteBaseMenu(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), webCtx)
	} else {
		response.OkWithMessage("删除成功", webCtx)

	}
}

// @Tags menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "更新菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/updateBaseMenu [post]
func (controller *MenuController) UpdateBaseMenu(webCtx SpringWeb.WebContext) {
	var menu model.SysBaseMenu
	_ = webCtx.Bind(&menu)
	MenuVerify := utils.Rules{
		"Path":      {"notEmpty"},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0"), "ge=0"},
	}
	MenuVerifyErr := utils.Verify(menu, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), webCtx)
		return
	}
	MetaVerify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	MetaVerifyErr := utils.Verify(menu.Meta, MetaVerify)
	if MetaVerifyErr != nil {
		response.FailWithMessage(MetaVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysMenuService.UpdateBaseMenu(menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改失败：%v", err), webCtx)
	} else {
		response.OkWithMessage("修改成功", webCtx)
	}
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取菜单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuById [post]
func (controller *MenuController) GetBaseMenuById(webCtx SpringWeb.WebContext) {
	var idInfo request.GetById
	_ = webCtx.Bind(&idInfo)
	MenuVerify := utils.Rules{
		"Id": {"notEmpty"},
	}
	MenuVerifyErr := utils.Verify(idInfo, MenuVerify)
	if MenuVerifyErr != nil {
		response.FailWithMessage(MenuVerifyErr.Error(), webCtx)
		return
	}
	err, menu := controller.SysMenuService.GetBaseMenuById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败：%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysBaseMenuResponse{Menu: menu}, webCtx)
	}
}
