package routers

import (
	"github.com/csDeng/blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() *gin.Engine {
	r = gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 全局路由
	globalRoutes()

	// 普通用户
	commonRoutes()

	// 管理员的接口
	adminRoutes()

	return r
}
