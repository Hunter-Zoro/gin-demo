package core

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFile("/", "./static/dist/index.html")
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	r.Static("/public", "./static")
	r.Static("/storage", "./storage/app/public")

	apiGroup := r.Group("/api")
	router.SetApiGroupRouter(apiGroup)
	return r
}

func RunServer() {
	r := setupRouter()
	r.Run(":" + global.DEMO_CONFIG.App.Port)
}
