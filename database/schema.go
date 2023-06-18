package database

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        string         `gorm:"primaryKey;notNull" faker:"unique,uuid_digit"`
	Email     string         `gorm:"unique" validation:"required,email"`
	Password  string         `faker:"-" validation:"required"`
	Name      string         `faker:"name" validation:"required"`
	Role      string         `faker:"oneof:admin,staff,user" validation:"required"`
	CreatedAt time.Time      `gorm:"autoCreateTime:mili" faker:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}

type RT struct {
	ID        string         `gorm:"primaryKey;notNull" faker:"unique,uuid_digit"`
	IDRW      string         `faker:"unique,uuid_digit"`
	NoRT      string         `gorm:"unique" faker:"oneof:1,2,3,4,5,6,7,8,9,10"`
	Nama      string         `faker:"name"`
	Alamat    string         `faker:"address"`
	Penduduk  []Penduduk     `gorm:"foreignKey:IDRT"`
	CreatedAt time.Time      `gorm:"autoCreateTime:mili" faker:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}

type RW struct {
	ID        string         `gorm:"primaryKey;notNull" faker:"unique,uuid_digit"`
	NoRW      string         `gorm:"unique" faker:"oneof:1,2,3,4,5,6,7,8,9,10"`
	Nama      string         `faker:"name"`
	Alamat    string         `faker:"address"`
	RT        []RT           `gorm:"foreignKey:IDRW"`
	CreatedAt time.Time      `gorm:"autoCreateTime:mili" faker:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}

type Penduduk struct {
	NIK             string         `gorm:"primaryKey;notNull" faker:"unique,uuid_digit"`
	IDRT            string         `faker:"unique,uuid_digit"`
	IDDomisili      string         `faker:"unique,uuid_digit"`
	Nama            string         `faker:"name"`
	TempatLahir     string         `faker:"city"`
	TanggalLahir    string         `faker:"date"`
	Agama           string         `faker:"oneof:islam,kristen,katolik,hindu,budha,konghucu"`
	Pekerjaan       string         `faker:"oneof:petani,nelayan,pns,tni,polri,swasta,wiraswasta"`
	Pendidikan      string         `faker:"oneof:sd,smp,sma,d1,d2,d3,s1,s2,s3"`
	StatusKawin     string         `faker:"oneof:kawin,belum_kawin"`
	StatusHubungan  string         `faker:"oneof:ayah,ibu,anak,suami,istri"`
	Kewarganegaraan string         `faker:"oneof:wni,wna"`
	CreatedAt       time.Time      `gorm:"autoCreateTime:mili" faker:"-"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime:mili" faker:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" faker:"-"`
}

type Domisili struct {
	ID        string         `gorm:"primaryKey;notNull" faker:"unique,uuid_digit"`
	Penduduk  []Penduduk     `gorm:"foreignKey:IDDomisili"`
	Alamat    string         `faker:"address"`
	CreatedAt time.Time      `gorm:"autoCreateTime:mili" faker:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}
