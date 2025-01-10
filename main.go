package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func main() {

	////创建一个默认的路由引擎
	//r := gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello  world",
	//	})
	//})
	//r.Run(":9090")
	global.DEMO_VIPER = core.InitializeViper()
	global.LOG = core.InitializeLog()
	global.LOG.Info("log init success!")

	global.DB = core.InitializeDB()
	defer func() {
		if global.DB != nil {
			db, _ := global.DB.DB()
			db.Close()
		}
	}()

	core.InitializeValidator()

	core.RunServer()
	/*	fmt.Print(global.DEMO_CONFIG)
		core.RunServer()*/

}
