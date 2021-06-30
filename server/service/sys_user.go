package service

import (
	"errors"

	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-utils"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func init() {
	SpringBoot.RegisterBean(new(SysUserService))
}

type SysUserService struct {
	Db *gorm.DB `autowire:""`
}

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func (service *SysUserService) Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册
	notRegister := service.Db.Where("username = ?", u.Username).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = SpringUtils.MD5(u.Password)
		u.UUID = uuid.NewV4()
		err = service.Db.Create(&u).Error
	}
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func (service *SysUserService) Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = SpringUtils.MD5(u.Password)
	err = service.Db.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser

func (service *SysUserService) ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	// TODO:后期修改jwt+password模式
	u.Password = SpringUtils.MD5(u.Password)
	err = service.Db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", SpringUtils.MD5(newPassword)).Error
	return err, u
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func (service *SysUserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := service.Db
	var userList []model.SysUser
	err = db.Find(&userList).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func (service *SysUserService) SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = service.Db.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func (service *SysUserService) DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = service.Db.Where("id = ?", id).Delete(&user).Error
	return err
}

// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser

func (service *SysUserService) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = service.Db.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}
