package dto

type CreatePendudukDto struct {
	NIK             string `json:"nik" validate:"required"`
	Nama            string `json:"nama" validate:"required"`
	TempatLahir     string `json:"tempat_lahir" validate:"required"`
	TanggalLahir    string `json:"tanggal_lahir" validate:"required"`
	Agama           string `json:"agama" validate:"required"`
	Pekerjaan       string `json:"pekerjaan" validate:"required"`
	Pendidikan      string `json:"pendidikan" validate:"required"`
	StatusKawin     string `json:"status_kawin" validate:"required"`
	StatusHubungan  string `json:"status_hubungan" validate:"required"`
	Kewarganegaraan string `json:"kewarga_neg" validate:"required"`
}

type ReadPendudukDto struct {
	NIK             string `json:"nik"`
	Nama            string `json:"nama"`
	TempatLahir     string `json:"tempat_lahir"`
	TanggalLahir    string `json:"tanggal_lahir"`
	Agama           string `json:"agama"`
	Pekerjaan       string `json:"pekerjaan"`
	Pendidikan      string `json:"pendidikan"`
	StatusKawin     string `json:"status_kawin"`
	StatusHubungan  string `json:"status_hubungan"`
	Kewarganegaraan string `json:"kewarganegaraan"`
}

type UpdatePendudukDto struct {
	Nama            string `json:"nama"`
	TempatLahir     string `json:"tempat_lahir"`
	TanggalLahir    string `json:"tanggal_lahir"`
	Agama           string `json:"agama"`
	Pekerjaan       string `json:"pekerjaan"`
	Pendidikan      string `json:"pendidikan"`
	StatusKawin     string `json:"status_kawin"`
	StatusHubungan  string `json:"status_hubungan"`
	Kewarganegaraan string `json:"kewarganegaraan"`
}
