package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	db, err := gorm.Open("mysql", connect)

	fmt.Printf("==========%#v\n", connect)

	if err != nil {
		panic(err.Error())
	}
	return db
}
