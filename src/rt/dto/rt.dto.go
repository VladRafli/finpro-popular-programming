package dto

type CreateRtDto struct {
	NoRT  string `json:"no_rt" validate:"required"`
	Nama  string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
}

type ReadRtDto struct {
	NoRT  string `json:"no_rt"`
	Nama  string `json:"nama"`
	Alamat string `json:"alamat"`
}

type UpdateRtDto struct {
	NoRT  string `json:"no_rt"`
	Nama  string `json:"nama"`
	Alamat string `json:"alamat"`
}