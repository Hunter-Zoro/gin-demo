package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/spf13/viper"
)

var (
	DEMO_CONFIG config.Configuration
	DEMO_VIPER  *viper.Viper
)
