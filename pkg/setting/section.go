package setting

import "time"

type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppConfig struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	UploadSavePath  string
}

type DatabaseConfig struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
