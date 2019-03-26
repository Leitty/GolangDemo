package main

import (
	"Gin/learnGin/golangDemo/eureka"
	"Gin/learnGin/golangDemo/models"
	"Gin/learnGin/golangDemo/pkg/gredis"
	"Gin/learnGin/golangDemo/pkg/logging"
	"Gin/learnGin/golangDemo/pkg/setting"
	"Gin/learnGin/golangDemo/routers"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf("%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	setting.Setup()
	models.Setup()
	err := gredis.Setup()
	if err !=nil {
		logging.Warn("Fail to connect to redis")
	}

	go startWebServer()

	eureka.Register()
	defer eureka.DeRegister()
	go eureka.StartHeartbeat()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func startWebServer() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}