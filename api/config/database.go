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
func ConnectDB() {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		GetEnvString("MYSQL_USER", "root"),
		GetEnvString("MYSQL_PASS", "password"),
		GetEnvString("MYSQL_HOST", "127.0.0.1"),
		GetEnvString("MYSQL_PORT", "3306"),
		GetEnvString("MYSQL_DATABASE", "app_name"),
	)
	var err error

	DB, err = gorm.Open("mysql", connect)

	if err != nil {
		panic(err.Error())
	}
}
