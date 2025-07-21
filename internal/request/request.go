package request

import (
	"mime/multipart"
)

type UpdateTokoRequest struct {
	Nama   			string  				`form:"nama" binding:"required" validate:"required,min=2,max=100"`
	Alamat 			string  				`form:"alamat" binding:"required" validate:"required,min=2,max=500"`
	Gambar     		*multipart.FileHeader 	`form:"gambar"`
	NomorToko  		string					`form:"nomor_toko" binding:"required" validate:"required,min=10,max=15"`
	AturanPajak 	float64					`form:"aturan_pajak" binding:"required" validate:"required,min=10,max=15"`
}