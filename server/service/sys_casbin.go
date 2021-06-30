package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"

	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	"github.com/casbin/gorm-adapter"
	"github.com/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(SysCasbinService))
}

type SysCasbinService struct {
	Db *gorm.DB `autowire:""`
}

// @title    UpdateCasbin
// @description   update casbin authority, 更新casbin权限
// @auth                     （2020/04/05  20:22）
// @param     authorityId      string
// @param     casbinInfos      []CasbinInfo
// @return                     error

func (service *SysCasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	service.ClearCasbin(0, authorityId)
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			ID:          0,
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		addflag := service.AddCasbin(cm)
		if addflag == false {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

// @title    AddCasbin
// @description   add casbin authority, 添加权限
// @auth                     （2020/04/05  20:22）
// @param     cm              model.CasbinModel
// @return                    bool

func (service *SysCasbinService) AddCasbin(cm model.CasbinModel) bool {
	e := service.Casbin()
	return e.AddPolicy(cm.AuthorityId, cm.Path, cm.Method)
}

// @title    UpdateCasbinApi
// @description   update casbin apis, API更新随动
// @auth                     （2020/04/05  20:22）
// @param     oldPath          string
// @param     newPath          string
// @param     oldMethod        string
// @param     newMethod        string
// @return                     error

func (service *SysCasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	var cs []model.CasbinModel
	err := service.Db.Table("casbin_rule").Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Find(&cs).Updates(map[string]string{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// @title    GetPolicyPathByAuthorityId
// @description   get policy path by authorityId, 获取权限列表
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return                    []string

func (service *SysCasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := service.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// @title    ClearCasbin
// @description   清除匹配的权限
// @auth                     （2020/04/05  20:22）
// @param     v               int
// @param     p               string
// @return                    bool

func (service *SysCasbinService) ClearCasbin(v int, p ...string) bool {
	e := service.Casbin()
	return e.RemoveFilteredPolicy(v, p...)

}

// @title    Casbin
// @description   store to DB, 持久化到数据库  引入自定义规则
// @auth                     （2020/04/05  20:22）

func (service *SysCasbinService) Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(service.Db)
	e := casbin.NewEnforcer(SpringBoot.GetStringProperty("casbin.model-path"), a)
	e.AddFunction("ParamsMatch", service.ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// @title    ParamsMatch
// @description   customized rule, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     fullNameKey1    string
// @param     key2            string
// @return                    bool

func (service *SysCasbinService) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @title    ParamsMatchFunc
// @description   customized function, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     args            ...interface{}
// @return                    interface{}
// @return                    error

func (service *SysCasbinService) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return service.ParamsMatch(name1, name2), nil
}
