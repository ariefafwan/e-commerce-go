package request

import (
	"mime/multipart"
)

type UpdateTokoRequest struct {
	Nama   			string  				`form:"nama" validate:"required,min=2,max=100"`
	Alamat 			string  				`form:"alamat" validate:"required,min=2,max=500"`
	Gambar     		*multipart.FileHeader 	`form:"gambar"`
	NomorToko  		string					`form:"nomor_toko" validate:"required,min=10,max=15"`
	AturanPajak 	float64					`form:"aturan_pajak" validate:"required,min=10,max=15"`
}