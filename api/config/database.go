package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB のグローバル変数
var DB *gorm.DB

// ConnectDB connect db
func ConnectDB(cfg *Config) {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Mysql.User,
		cfg.Mysql.Pass,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database,
	)
	var err error

	DB, err = gorm.Open("mysql", connect)

	if err != nil {
		panic(err.Error())
	}
}
