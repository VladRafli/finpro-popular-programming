package database

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        string         `gorm:"primaryKey" faker:"unique,uuid_digit"`
	Email     string         `gorm:"unique" validation:"required,email"`
	Password  string         `faker:"-" validation:"required"`
	Name      string         `faker:"name" validation:"required"`
	Role      string         `faker:"oneof:admin,staff,user" validation:"required"`
	CreatedAt time.Time      `gorm:"autoCreateTime:mili" faker:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}
