package service

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"

	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(SysSystemService))
}

type SysSystemService struct {
}

// @title    GetSystemConfig
// @description   读取配置文件
// @auth                     （2020/04/05  20:22）
// @return    err             error
// @return    conf            Server
func (service *SysSystemService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// @title    SetSystemConfig
// @description   set system config, 设置配置文件
// @auth                    （2020/04/05  20:22）
// @param     system         model.System
// @return    err            error
func (service *SysSystemService) SetSystemConfig(system model.System) (err error) {
	// TODO 该功能并不完善，重新保存后的文件格式变得很乱，应该考虑其他存储方式
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}
