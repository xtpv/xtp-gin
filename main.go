package main

import (
	"github.com/gin-gonic/gin"
	G "github.com/xtp2217866847/xtp-gin/global"
	"github.com/xtp2217866847/xtp-gin/internal/routers"
	"log"
	"net/http"
)

func init() {
	err := G.InitResource()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
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
