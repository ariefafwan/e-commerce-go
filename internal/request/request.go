package request

import (
	"e-commerce-go/internal/models"
	"mime/multipart"
)

type UpdateTokoRequest struct {
	Nama   			string  				`form:"nama" validate:"required,min=2,max=100"`
	Alamat 			string  				`form:"alamat" validate:"required,min=2,max=500"`
	Gambar     		*multipart.FileHeader 	`form:"gambar"`
	NomorToko  		string					`form:"nomor_toko" validate:"required,min=10,max=15"`
	AturanPajak 	float64					`form:"aturan_pajak" validate:"required,min=10,max=15"`
}

type CreatePelangganRequest struct {
	IDUser 			string  				`form:"id_user" validate:"required,fk_exists=users:id,unique=master_pelanggan:id_user"`
	NamaLengkap  	string  				`form:"nama_lengkap" validate:"required,min=2,max=500"`
	NamaPanggilan 	string  				`form:"nama_panggilan" validate:"required,min=2,max=500"`
	Phone 			string  				`form:"phone" validate:"required,min=10,max=15"`
}

type UpdatePelangganRequest struct {
	ID 				string  				`form:"id" validate:"required"`
	IDUser 			string  				`form:"id_user" validate:"required,fk_exists=users:id,unique_except=master_pelanggan:id_user:id"`
	NamaLengkap  	string  				`form:"nama_lengkap" validate:"required,min=2,max=500"`
	NamaPanggilan 	string  				`form:"nama_panggilan" validate:"required,min=2,max=500"`
	Phone 			string  				`form:"phone" validate:"required,min=10,max=15"`
}

type AlamatPelangganRequest struct {
	IDPelanggan  	string  				`form:"id_pelanggan" validate:"required,fk_exists=master_pelanggan:id"`
	AlamatLengkap 	string  				`form:"alamat_lengkap" validate:"required,min=2,max=500"`
	KodePos 		string  				`form:"kode_pos" validate:"required,min=2,max=500"`
	Kota 			string  				`form:"kota" validate:"required,min=2,max=500"`
	Negara 			string  				`form:"negara" validate:"required,min=2,max=500"`
	NomorPenerima 	string  				`form:"nomor_penerima" validate:"required,min=2,max=500"`
	NamaPenerima 	string  				`form:"nama_penerima" validate:"required,min=2,max=500"`
}

type CreateKategoriProdukRequest struct {
	Nama 			string  				`form:"nama" validate:"required,min=2,max=500,unique=master_kategori_produk:nama"`
	IDParent 		*string  				`form:"id_parent" validate:"omitempty,fk_exists=master_kategori_produk:id"`
	Urutan 			*uint8					`form:"urutan" validate:"omitempty,unique=master_kategori_produk:urutan"`
}

type UpdateKategoriProdukRequest struct {
	ID 				string  				`form:"id" validate:"required"`
	Nama 			string  				`form:"nama" validate:"required,min=2,max=500,unique_except=master_kategori_produk:nama:id"`
	IDParent 		*string  				`form:"id_parent" validate:"omitempty,fk_exists=master_kategori_produk:id"`
	Urutan 			*uint8					`form:"urutan" validate:"omitempty,unique_except=master_kategori_produk:urutan:id"`
}

type ProdukRequest struct {
	Nama      	string					`form:"nama" validate:"required,min=2,max=100"`
	Thumbnail 	*multipart.FileHeader 	`form:"thumbnail"`
	MinHarga  	float64					`form:"min_harga" validate:"required"`
	MaxHarga  	float64					`form:"max_harga" validate:"required"`
	Deskripsi 	string					`form:"deskripsi" validate:"required"`
	Status    	*models.StatusProduk	`form:"status" validate:"required"`
	IDKategoriProduk []string			`form:"id_kategori_produk" validate:"required,fk_exists=master_kategori_produk:id"`
}

type ProdukVariantRequest struct {
	NamaVariant string					`form:"nama_variant" validate:"required,min=2,max=100"`
	IDProduk 	string					`form:"id_produk" validate:"required,fk_exists=master_produk:id"`
	Harga 		float64					`form:"harga" validate:"required"`
	Stok 		int						`form:"stok" validate:"required"`
}

type ProdukGaleriRequest struct {
	IDProduk 	string					`form:"id_produk" validate:"required,fk_exists=master_produk:id"`
	Gambar 		*multipart.FileHeader 	`form:"gambar" validate:"required"`
	Urutan 		*uint8					`form:"urutan" validate:"omitempty,unique_except=master_produk_galeri:urutan:id"`
}