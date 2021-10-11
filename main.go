package main

import (
	"context"
	"github.com/gin-gonic/gin"
	G "github.com/xtpv/xtp-gin/global"
	"github.com/xtpv/xtp-gin/internal/routers"
	"log"
	"net/http"
)

func init() {
	err := G.InitResource()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/xtpv/xtp-gin
func main() {
	// 开启异步任务
	G.StartEvent(context.Background())
	err := HttpStart()
	if err != nil {
		log.Panic(err)
	}
}

// HttpStart 启动 http 服务
func HttpStart() error {
	gin.SetMode(G.ServerConfig.RunMode)
	router := routers.NewRouter()
	cfg := G.ServerConfig
	s := &http.Server{
		Addr:           ":" + cfg.HttpPort,
		Handler:        router,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}
