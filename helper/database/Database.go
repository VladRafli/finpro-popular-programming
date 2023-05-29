package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(192.168.1.11)/finpro-ppt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}

func Disconnect(db *gorm.DB) {
	sqlDb, err := db.DB()

	if err != nil {
		panic("Failed to disconnect from database!")
	}

	sqlDb.Close()
}
