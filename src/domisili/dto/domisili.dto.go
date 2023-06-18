package dto

type CreateDomisiliDto struct {
	Alamat string `json:"alamat" validate:"required"`
}

type ReadDomisiliDto struct {
	Alamat string `json:"alamat"`
}

type UpdateDomisiliDto struct {
	Alamat string `json:"alamat"`
}