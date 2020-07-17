package initialize

import (
	"gin-vue-admin/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-spring/go-spring/starter-mysql-gorm"
	"github.com/go-spring/go-spring/spring-boot"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {
	SpringBoot.Config(func(db *gorm.DB, maxIdleConns, maxOpenConns int, logMode bool) {
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(maxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(maxOpenConns)
		global.GVA_DB.LogMode(logMode)
		DBTables()
	}, "1:${db.max-idle-conns}", "2:${db.max-open-conns}", "3:${db.log-mode}")
}
