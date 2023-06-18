package dto

type CreateRwDto struct {
	NoRW  string `json:"no_rw" validate:"required"`
	Nama  string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
}

type ReadRwDto struct {
	NoRW  string `json:"no_rw"`
	Nama  string `json:"nama"`
	Alamat string `json:"alamat"`
}

type UpdateRwDto struct {
	NoRW  string `json:"no_rw"`
	Nama  string `json:"nama"`
	Alamat string `json:"alamat"`
}