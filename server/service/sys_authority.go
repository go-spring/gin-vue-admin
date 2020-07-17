package service

import (
	"errors"
	"strconv"

	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"

	"github.com/go-spring/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
)

func init() {
	SpringBoot.RegisterBean(new(SysAuthorityService))
}

type SysAuthorityService struct {
	SysMenuService   *SysMenuService   `autowire:""`
	SysCasbinService *SysCasbinService `autowire:""`
	Db               *gorm.DB          `autowire:""`
}

// @title    CreateAuthority
// @description   创建一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority

func (service *SysAuthorityService) CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	notHas := service.Db.Where("authority_id = ?", auth.AuthorityId).Find(&authorityBox).RecordNotFound()
	if !notHas {
		return errors.New("存在相同角色id"), auth
	}
	err = service.Db.Create(&auth).Error
	return err, auth
}

// @title    CopyAuthority
// @description   复制一个角色
// @auth                     （2020/04/05  20:22）
// @param     copyInfo        response.SysAuthorityCopyResponse
// @return                    error
// @return    authority       model.SysAuthority

func (service *SysAuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	notHas := service.Db.Where("authority_id = ?", copyInfo.Authority.AuthorityId).Find(&authorityBox).RecordNotFound()
	if !notHas {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []model.SysAuthority{}
	err, menus := service.SysMenuService.GetMenuAuthority(copyInfo.OldAuthorityId)
	var baseMenu []model.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = service.Db.Create(&copyInfo.Authority).Error

	paths := service.SysCasbinService.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = service.SysCasbinService.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = service.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

// @title    UpdateAuthority
// @description   更改一个角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @return    authority       model.SysAuthority

func (service *SysAuthorityService) UpdateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	err = service.Db.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

// @title    DeleteAuthority
// @description   删除角色
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// 删除角色

func (service *SysAuthorityService) DeleteAuthority(auth *model.SysAuthority) (err error) {
	err = service.Db.Where("authority_id = ?", auth.AuthorityId).Find(&model.SysUser{}).Error
	if err == nil {
		err = errors.New("此角色有用户正在使用禁止删除")
		return
	}
	err = service.Db.Where("parent_id = ?", auth.AuthorityId).Find(&model.SysAuthority{}).Error
	if err == nil {
		err = errors.New("此角色存在子角色不允许删除")
		return
	}
	db := service.Db.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth)
	if len(auth.SysBaseMenus) > 0 {
		err = db.Association("SysBaseMenus").Delete(auth.SysBaseMenus).Error
	} else {
		err = db.Error
	}
	service.SysCasbinService.ClearCasbin(0, auth.AuthorityId)
	return err
}

// @title    GetInfoList
// @description   删除文件切片记录
// @auth                     （2020/04/05  20:22）
// @param     info            request.PaveInfo
// @return                    error
// 分页获取数据

func (service *SysAuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := service.Db
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = service.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

// @title    GetAuthorityInfo
// @description   获取所有角色信息
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error
// @param     authority       model.SysAuthority

func (service *SysAuthorityService) GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = service.Db.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// @title    SetDataAuthority
// @description   设置角色资源权限
// @auth                     （2020/04/05  20:22）
// @param     auth            model.SysAuthority
// @return                    error

func (service *SysAuthorityService) SetDataAuthority(auth model.SysAuthority) error {
	var s model.SysAuthority
	service.Db.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := service.Db.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId).Error
	return err
}

// @title    SetMenuAuthority
// @description   菜单与角色绑定
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error

func (service *SysAuthorityService) SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	service.Db.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := service.Db.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus).Error
	return err
}

// @title    findChildrenAuthority
// @description   查询子角色
// @auth                     （2020/04/05  20:22）
// @param     auth            *model.SysAuthority
// @return                    error

func (service *SysAuthorityService) findChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = service.Db.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = service.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
