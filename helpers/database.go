package helpers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to postgres database
func ConnectDatabase(host string, user string, password string, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
	}

	return db
}