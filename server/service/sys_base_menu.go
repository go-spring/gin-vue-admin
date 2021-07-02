package service

import (
	"errors"
	"gin-vue-admin/model"

	"github.com/go-spring/spring-logger"
)

// @title    DeleteBaseMenu
// @description   删除基础路由
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error

func (service *SysMenuService) DeleteBaseMenu(id float64) (err error) {
	err = service.Db.Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := service.Db.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		if len(menu.SysAuthoritys) > 0 {
			err = db.Association("SysAuthoritys").Delete(menu.SysAuthoritys).Error
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// @title    UpdateBaseMenu
// @description   更新路由
// @auth                     （2020/04/05  20:22）
// @param     menu            model.SysBaseMenu
// @return    err             errorgetMenu

func (service *SysMenuService) UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort
	db := service.Db.Where("id = ?", menu.ID).Find(&oldMenu)
	if oldMenu.Name != menu.Name {
		notSame := service.Db.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.SysBaseMenu{}).RecordNotFound()
		if !notSame {
			SpringLogger.Debug("存在相同name修改失败")
			return errors.New("存在相同name修改失败")
		}
	}
	err = db.Updates(upDateMap).Error
	SpringLogger.Debug("菜单修改时候，关联菜单err:%v", err)
	return err
}

// @title    GetBaseMenuById
// @description   get current menus, 返回当前选中menu
// @auth                     （2020/04/05  20:22）
// @param     id              float64
// @return    err             error

func (service *SysMenuService) GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = service.Db.Where("id = ?", id).First(&menu).Error
	return
}
