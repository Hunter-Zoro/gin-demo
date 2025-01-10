package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DEMO_CONFIG config.Configuration
	DEMO_VIPER  *viper.Viper
	LOG         *zap.Logger
	DB          *gorm.DB
)
