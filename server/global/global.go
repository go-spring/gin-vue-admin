package global

import (
	"gin-vue-admin/config"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *oplogging.Logger
)
