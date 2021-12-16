package main

import (
	"fmt"
	"net/http"

	"github.com/csDeng/go-gin-example/models"
	"github.com/csDeng/go-gin-example/pkg/logging"
	"github.com/csDeng/go-gin-example/pkg/setting"
	"github.com/csDeng/go-gin-example/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
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
