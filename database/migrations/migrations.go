package migrations

import (
	"my_kelurahan/database"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&database.Users{})
}
