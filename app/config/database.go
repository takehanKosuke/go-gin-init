package config

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

// ConnectDB connect db
func ConnectDB(cfg *Config) *gorm.DB {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Mysql.User,
		cfg.Mysql.Pass,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database,
	)
	var err error

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
