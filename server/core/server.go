package core

import (
	"errors"
	"net/http"

	"gin-vue-admin/global"
	"gin-vue-admin/initialize"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
	"github.com/go-spring/go-spring/starter-web"
)

func init() {
	SpringBoot.RegisterBeanFn(func(config WebStarter.WebServerConfig) SpringWeb.WebContainer {
		c := SpringGin.NewContainer(SpringWeb.ContainerConfig{Port: config.Port})
		c.SetEnableSwagger(false)
		return c
	})
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	for _, info := range Router.Routes() {
		switch info.Method {
		case http.MethodGet:
			SpringBoot.GET(info.Path, SpringGin.Gin(info.HandlerFunc))
		case http.MethodPost:
			SpringBoot.POST(info.Path, SpringGin.Gin(info.HandlerFunc))
		case http.MethodDelete:
			SpringBoot.DELETE(info.Path, SpringGin.Gin(info.HandlerFunc))
		case http.MethodPut:
			SpringBoot.PUT(info.Path, SpringGin.Gin(info.HandlerFunc))
		case http.MethodHead:
			SpringBoot.HEAD(info.Path, SpringGin.Gin(info.HandlerFunc))
		default:
			panic(errors.New("unsupported http method " + info.Method))
		}
	}

	// 插件安装 暂时只是后台功能 添加model 添加路由 添加对数据库的操作  详细插件测试模板可看https://github.com/piexlmax/gvaplug  此处不建议投入生产
	//err := initialize.InstallPlug(global.GVA_DB, Router, gvaplug.GvaPlug{})
	//if err != nil {
	//	panic(fmt.Sprintf("插件安装失败： %v", err))
	//}
	// end 插件描述

	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	//s := &http.Server{
	//	Addr:           address,
	//	Handler:        Router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	//time.Sleep(10 * time.Microsecond)
	//global.GVA_LOG.Debug("server run success on ", address)

	//fmt.Printf(`欢迎使用 Gin-Vue-Admin
	//默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	//默认前端文件运行地址:http://127.0.0.1:8080`, s.Addr)

	// global.GVA_LOG.Error(s.ListenAndServe())
	SpringBoot.RunApplication("./")
}
