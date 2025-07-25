package request

import (
	"e-commerce-go/internal/models"
	"mime/multipart"
)

type UpdateTokoRequest struct {
	Nama   			string  				`form:"nama" validate:"required,min=2,max=100"`
	Alamat 			string  				`form:"alamat" validate:"required,min=2,max=500"`
	IDProvinsi 		string  				`form:"id_provinsi" validate:"required,min=2,max=500,fk_exists=master_provinsi:id"`
	IDKota 			string  				`form:"id_kota" validate:"required,min=2,max=500,fk_exists=master_kota:id"`
	IDKecamatan 	string  				`form:"id_kecamatan" validate:"required,min=2,max=500,fk_exists=master_kecamatan:id"`
	Gambar     		*multipart.FileHeader 	`form:"gambar"`
	NoTelp  		string					`form:"no_telp" validate:"required,min=10,max=15"`
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
	Label 			string  				`form:"label" validate:"required,min=2,max=500"`
	AlamatLengkap 	string  				`form:"alamat_lengkap" validate:"required,min=2,max=500"`
	KodePos 		string  				`form:"kode_pos" validate:"required,min=2,max=500"`
	IDProvinsi 		string  				`form:"id_provinsi" validate:"required,min=2,max=500,fk_exists=master_provinsi:id"`
	IDKota 			string  				`form:"id_kota" validate:"required,min=2,max=500,fk_exists=master_kota:id"`
	IDKecamatan 	string  				`form:"id_kecamatan" validate:"required,min=2,max=500,fk_exists=master_kecamatan:id"`
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
	Status    	*models.StatusProduk	`form:"status"`
	Berat     	float64					`form:"berat" validate:"required"`
	IDKategoriProduk []string			`form:"id_kategori_produk" validate:"required"`
}

type ProdukVariantRequest struct {
	NamaVariant string					`form:"nama_variant" validate:"required,min=2,max=100"`
	Harga 		float64					`form:"harga" validate:"required"`
	Stok 		int						`form:"stok" validate:"required"`
}

type ProdukGaleriRequest struct {
	Gambar 		*multipart.FileHeader 	`form:"gambar" validate:"required"`
	Urutan 		*uint8					`form:"urutan" validate:"omitempty,unique_except=master_produk_galeri:urutan:id"`
}

type UpdateKeranjangRequest struct {
	IDPelanggan  	string  			`form:"id_pelanggan" validate:"required,fk_exists=master_pelanggan:id"`
	BerlakuSampai   *string 			`form:"berlaku_sampai" validate:"required,date_format,future_date"`
}

type CreateKeranjangItemRequest struct {
	IDPelanggan  	string  			`form:"id_pelanggan" validate:"required,fk_exists=master_pelanggan:id"`
	IDVariantProduk string  			`form:"id_variant_produk" validate:"required,fk_exists=master_produk_variant:id"`
	Quantity     	int					`form:"quantity" validate:"required,min=1"`
}

type UpdateKeranjangItemRequest struct {
	IDKeranjang  	string  			`form:"id_keranjang" validate:"required,fk_exists=transaksi_keranjang:id"`
	IDVariantProduk string  			`form:"id_variant_produk" validate:"required,fk_exists=master_produk_variant:id"`
	Quantity     	int					`form:"quantity" validate:"required,min=1"`
}

type KalkulasiTransaksiRequest struct {
	IDItems 				[]string  			`form:"id_items" validate:"required"`
	IDAlamatPelanggan 		*string  			`form:"id_alamat_pelanggan" validate:"omitempty,fk_exists=master_alamat_pelanggan:id"`
}