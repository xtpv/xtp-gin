package global

import (
	"github.com/xtpv/xtp-gin/pkg/setting"
	"time"
)

var (
	ServerConfig   *setting.ServerConfig
	AppConfig      *setting.AppConfig
	DatabaseConfig *setting.DatabaseConfig
)

func initConfig() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &ServerConfig)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &AppConfig)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &DatabaseConfig)
	if err != nil {
		return err
	}

	ServerConfig.ReadTimeout *= time.Second
	ServerConfig.WriteTimeout *= time.Second
	return nil
}
