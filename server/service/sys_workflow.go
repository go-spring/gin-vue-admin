package service

import (
	"gin-vue-admin/model"

	"github.com/go-spring/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
)

func init() {
	SpringBoot.RegisterBean(new(SysWorkflowService))
}

type SysWorkflowService struct {
	Db *gorm.DB `autowire:""`
}

// @title    Create
// @description   create a workflow, 创建工作流
// @auth                     （2020/04/05  20:22）
// @param     wk              model.SysWorkflow
// @return                    error

func (service *SysWorkflowService) Create(wk model.SysWorkflow) error {
	err := service.Db.Create(&wk).Error
	return err
}
