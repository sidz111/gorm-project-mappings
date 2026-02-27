package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db_username := "root"
	db_pass := "root"
	db_host := "localhost"
	db_port := 3303
	db_name := "gorm_mappinga"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_username, db_pass, db_host, db_port, db_name)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = database
	return nil
}
