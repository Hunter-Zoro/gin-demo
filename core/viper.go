package core

import (
	"flag"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/Internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func InitializeViper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")

		flag.Parse()

		if config == "" {
			if configEnv := os.Getenv(Internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = Internal.ConfigDefaultFile

				case gin.ReleaseMode:
					config = Internal.ConfigReleaseFile
				case gin.TestMode:
					config = Internal.ConfigTestFile
				}
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\\n", gin.Mode(), config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\\n", gin.Mode(), configEnv)

			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}
	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conifg file：%s \n", err))
	}
	vip.WatchConfig()

	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed:", e.Name)
		if err = vip.Unmarshal(&global.DEMO_CONFIG); err != nil {
			fmt.Println("Unmarshal err:", err)
		}
	})
	if err = vip.Unmarshal(&global.DEMO_CONFIG); err != nil {
		panic(err)
	}
	fmt.Println("====1-viper====:viper init config success")
	return vip
}
