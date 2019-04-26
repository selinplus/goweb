package main

import (
	"fmt"
	"github.com/selinplus/goweb/models"
	"github.com/selinplus/goweb/pkg/logging"
	"github.com/selinplus/goweb/pkg/setting"
	"github.com/selinplus/goweb/routers"
	"net/http"
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
