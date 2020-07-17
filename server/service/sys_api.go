package service

import (
	"errors"

	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"github.com/go-spring/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
)

func init() {
	SpringBoot.RegisterBean(new(SysApiService))
}

type SysApiService struct {
	SysCasbinService *SysCasbinService `autowire:""`
	Db               *gorm.DB          `autowire:""`
}

// @title    CreateApi
// @description   create base apis, 新增基础api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @return                    error

func (service *SysApiService) CreateApi(api model.SysApi) (err error) {
	findOne := service.Db.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = service.Db.Create(&api).Error
	}
	return err
}

// @title    DeleteApi
// @description   delete a base api, 删除基础api
// @param     api             model.SysApi
// @auth                     （2020/04/05  20:22）
// @return                    error

func (service *SysApiService) DeleteApi(api model.SysApi) (err error) {
	err = service.Db.Delete(api).Error
	service.SysCasbinService.ClearCasbin(1, api.Path, api.Method)
	return err
}

// @title    GetInfoList
// @description   get apis by pagination, 分页获取数据
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     info            request.PageInfo
// @param     order           string
// @param     desc            bool
// @return    err             error
// @return    list            interface{}
// @return    total           int

func (service *SysApiService) GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := service.Db
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Find(&apiList).Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr, true).Find(&apiList).Error
		} else {
			err = db.Order("api_group", true).Find(&apiList).Error
		}
	}
	return err, apiList, total
}

// @title    GetAllApis
// @description   get all apis, 获取所有的api
// @auth                     （2020/04/05  20:22）
// @return       err          error
// @return       apis         []SysApi

func (service *SysApiService) GetAllApis() (err error, apis []model.SysApi) {
	err = service.Db.Find(&apis).Error
	return
}

// @title    GetApiById
// @description   根据id获取api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @param     id              float64
// @return                    error

func (service *SysApiService) GetApiById(id float64) (err error, api model.SysApi) {
	err = service.Db.Where("id = ?", id).First(&api).Error
	return
}

// @title    UpdateApi
// @description   update a base api, update api
// @auth                     （2020/04/05  20:22）
// @param     api             model.SysApi
// @return                    error

func (service *SysApiService) UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi

	err = service.Db.Where("id = ?", api.ID).First(&oldA).Error

	if oldA.Path != api.Path || oldA.Method != api.Method {
		flag := service.Db.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).RecordNotFound()
		if !flag {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = service.SysCasbinService.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = service.Db.Save(&api).Error
		}
	}
	return err
}
