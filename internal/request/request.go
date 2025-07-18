package request

type CreateBioskopRequest struct {
	Nama   string  `json:"nama" validate:"required"`
	Lokasi string  `json:"lokasi" validate:"required"`
	Rating float64 `json:"rating" validate:"required,gte=0,lte=10"`
}

type UpdateBioskopRequest struct {
	Nama   string  `json:"nama" validate:"required"`
	Lokasi string  `json:"lokasi" validate:"required"`
	Rating float64 `json:"rating" validate:"required,gte=0,lte=10"`
}