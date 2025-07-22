package request

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UpdateTokoRequest struct {
	Nama   			string  				`form:"nama" validate:"required,min=2,max=100"`
	Alamat 			string  				`form:"alamat" validate:"required,min=2,max=500"`
	Gambar     		*multipart.FileHeader 	`form:"gambar"`
	NomorToko  		string					`form:"nomor_toko" validate:"required,min=10,max=15"`
	AturanPajak 	float64					`form:"aturan_pajak" validate:"required,min=10,max=15"`
}

type PelangganRequest struct {
	IDUser 			uuid.UUID  				`form:"id_user" validate:"required,min=2,max=500,fk_exists=users|id"`
	NamaLengkap  	string  				`form:"nama_lengkap" validate:"required,min=2,max=500"`
	NamaPanggilan 	string  				`form:"nama_panggilan" validate:"required,min=2,max=500"`
	Alamat 			string  				`form:"alamat" validate:"required,min=2,max=500"`
	Phone 			string  				`form:"phone" validate:"required,min=10,max=15"`
}

type AlamatPelangganRequest struct {
	IDPelanggan  	uuid.UUID  				`form:"id_pelanggan" validate:"required,min=2,max=500,fk_exists=master_pelanggan|id"`
	AlamatLengkap 	string  				`form:"alamat_lengkap" validate:"required,min=2,max=500"`
	KodePos 		string  				`form:"kode_pos" validate:"required,min=2,max=500"`
	Kota 			string  				`form:"kota" validate:"required,min=2,max=500"`
	Negara 			string  				`form:"negara" validate:"required,min=2,max=500"`
	NomorPenerima 	string  				`form:"nomor_penerima" validate:"required,min=2,max=500"`
	NamaPenerima 	string  				`form:"nama_penerima" validate:"required,min=2,max=500"`
}

type KategoriProdukRequest struct {
	Nama 			string  				`form:"nama" validate:"required,min=2,max=500,unique=master_kategori_produk:nama"`
	IDParent 		*uuid.UUID  			`form:"id_parent" validate:"omitempty,fk_exists=master_kategori_produk:id"`
	Urutan 			*uint8					`form:"urutan" validate:"omitempty,unique=master_kategori_produk:urutan"`
}