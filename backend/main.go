package main

import (
	"fmt"
	"net/http"

	"github.com/csDeng/blog/models"
	"github.com/csDeng/blog/pkg/logging"
	"github.com/csDeng/blog/pkg/setting"
	"github.com/csDeng/blog/routers"
)

func main() {

	// 初始化配置
	setting.Setup()

	// 初始化数据库
	models.Setup()

	// 初始化日志
	logging.Setup()

	// 启用redis
	// gredis.Setup()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
