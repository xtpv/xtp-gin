package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

const MaxLifeTime = 3 * time.Minute

func initDB() error {
	cfg := DatabaseConfig
	dsn := cfg.UserName + ":" + cfg.Password + "@tcp(" + cfg.Host + ")/" + cfg.DBName + "?" + cfg.Charset
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(MaxLifeTime)
	DB = db
	return nil
}
