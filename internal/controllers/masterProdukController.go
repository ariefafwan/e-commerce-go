package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MasterProdukController struct {
	RepoProduk repositories.MasterProdukRepository
	RepoProdukVariant repositories.MasterProdukVariantRepository
	RepoProdukGaleri repositories.MasterProdukGaleriRepository
}

func NewMasterProdukController(repo repositories.MasterProdukRepository, 
					repoProdukVariant repositories.MasterProdukVariantRepository, 
					repoProdukGaleri repositories.MasterProdukGaleriRepository) MasterProdukController {
	return MasterProdukController{RepoProduk: repo, 
						RepoProdukVariant: repoProdukVariant, 
						RepoProdukGaleri: repoProdukGaleri}
}

func (mp *MasterProdukController) GetAll(c *gin.Context) {
	data, err := mp.RepoProduk.GetAll()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterProdukController) GetAllByKategori(c *gin.Context) {
	slug := c.Param("slug")
	data, err := mp.RepoProduk.GetAllByKategori(slug)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

// func (mp *MasterProdukController) GetByID(c *gin.Context) {
// 	id := c.Param("id")
// 	data, err := mp.RepoProduk.GetByID(id)
// 	if err != nil {
// 		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
// 		return
// 	}

// 	helpers.Success(c, http.StatusOK, data, "Success")
// }

func (mp *MasterProdukController) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	data, err := mp.RepoProduk.GetBySlug(slug)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterProdukController) GetAllVariant(c *gin.Context) {
	id := c.Param("id")
	data, err := mp.RepoProdukVariant.GetAll(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterProdukController) GetVariantByID(c *gin.Context) {
	id := c.Param("id")
	data, err := mp.RepoProdukVariant.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterProdukController) GetAllGaleri(c *gin.Context) {
	id := c.Param("id")
	data, err := mp.RepoProdukGaleri.GetAll(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterProdukController) GetGaleriByID(c *gin.Context) {
	id := c.Param("id")
	data, err := mp.RepoProdukGaleri.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterProdukController) Create(c *gin.Context) {
	var req request.ProdukRequest
	if req.Thumbnail != nil {
		helpers.Error(c, http.StatusBadRequest, nil, "Gambar harus diisi")
	}

	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	mimeType := req.Thumbnail.Header.Get("Content-Type")
	if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
		helpers.Error(c, http.StatusBadRequest, nil, "File logo harus berupa gambar JPEG/JPG atau PNG")
		return
	}
	filename, err := helpers.UploadImage(req.Thumbnail, "Toko")
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
		return
	}

	data := models.MasterProduk{
		Nama:       req.Nama,
		MinHarga:   req.MinHarga,
		MaxHarga:   req.MaxHarga,
		Deskripsi: 	req.Deskripsi,
		Thumbnail: 	filename,
		Status:     "Non Aktif",
	}

	if err := mp.RepoProduk.Create(&data, req.IDKategoriProduk); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) CreateVariant(c *gin.Context) {
	var req request.ProdukVariantRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	count, err := mp.RepoProdukVariant.CountAllByProduct(req.IDProduk)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	if count < 1 {
		if err := mp.RepoProduk.UpdateStatus(req.IDProduk, "Aktif"); err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
			return
		}
	}

	data := models.MasterProdukVariant{
		IDProduk: 		 uuid.MustParse(req.IDProduk),
		NamaVariant:     req.NamaVariant,
		Harga:    		 req.Harga,
		Stok:    		 req.Stok,
	}

	if err := mp.RepoProdukVariant.Create(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) CreateGaleri(c *gin.Context) {
	var req request.ProdukGaleriRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	if req.Gambar == nil {
		helpers.Error(c, http.StatusBadRequest, nil, "Gambar harus diisi")
		return
	}

	mimeType := req.Gambar.Header.Get("Content-Type")
	if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
		helpers.Error(c, http.StatusBadRequest, nil, "File logo harus berupa gambar JPEG/JPG atau PNG")
		return
	}
	filename, err := helpers.UploadImage(req.Gambar, "Produk")
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
		return
	}

	data := models.MasterProdukGaleri{
		IDProduk:   uuid.MustParse(req.IDProduk),
		Gambar:  filename,
	}

	if req.Urutan != nil {
		data.Urutan = *req.Urutan
	} else {
		data.Urutan = 0
	}

	if err := mp.RepoProdukGaleri.Create(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) Update(c *gin.Context) {
	var req request.ProdukRequest
	id := c.Param("id")
	existing, err := mp.RepoProduk.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}
	
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	if req.Thumbnail != nil {
		mimeType := req.Thumbnail.Header.Get("Content-Type")
		if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
			helpers.Error(c, http.StatusBadRequest, nil, "File logo harus berupa gambar JPEG/JPG atau PNG")
			return
		}
		filename, err := helpers.UploadImage(req.Thumbnail, "Toko")
		if err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
			return
		}
		helpers.DeleteImage(fmt.Sprintf("Toko/%s", existing.Thumbnail))
		existing.Thumbnail = filename
	}

	data := models.MasterProduk{
		ID:         uuid.MustParse(id),
		Nama:       req.Nama,
		MinHarga:   req.MinHarga,
		MaxHarga:   req.MaxHarga,
		Deskripsi: 	req.Deskripsi,
		Thumbnail: 	existing.Thumbnail,
	}

	if req.Status != nil {
		data.Status = *req.Status
	}

	if err := mp.RepoProduk.Update(&data, req.IDKategoriProduk); err != nil {
		helpers.DeleteImage(existing.Thumbnail)
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) UpdateVariant(c *gin.Context) {
	var req request.ProdukVariantRequest
	id := c.Param("id")
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	data := models.MasterProdukVariant{
		ID:       uuid.MustParse(id),
		NamaVariant:     req.NamaVariant,
		Harga:    		 req.Harga,
		Stok:    		 req.Stok,
		IDProduk: 		 uuid.MustParse(req.IDProduk),
	}

	if err := mp.RepoProdukVariant.Update(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (m *MasterProdukController) UpdateGaleri(c *gin.Context) {
	var req request.ProdukGaleriRequest
	id := c.Param("id")
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	existing, err := m.RepoProdukGaleri.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	if req.Gambar != nil {
		mimeType := req.Gambar.Header.Get("Content-Type")
		if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
			helpers.Error(c, http.StatusBadRequest, nil, "File logo harus berupa gambar JPEG/JPG atau PNG")
			return
		}
		filename, err := helpers.UploadImage(req.Gambar, "Produk")
		if err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
			return
		}
		helpers.DeleteImage(fmt.Sprintf("Produk/%s", existing.Gambar))
		existing.Gambar = filename
	}
	
	data := models.MasterProdukGaleri{
		ID:       uuid.MustParse(id),
		IDProduk: uuid.MustParse(req.IDProduk),
		Gambar:   existing.Gambar,
	}

	if req.Urutan != nil {
		data.Urutan = *req.Urutan
	} else {
		data.Urutan = existing.Urutan
	}
	
	if err := m.RepoProdukGaleri.Update(&data); err != nil {
		if req.Gambar != nil {
			helpers.DeleteImage(existing.Gambar)
		}
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (m *MasterProdukController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := m.RepoProduk.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (m *MasterProdukController) DeleteProdukVariant(c *gin.Context) {
	id := c.Param("id")
	if err := m.RepoProdukVariant.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (m *MasterProdukController) DeleteGaleri(c *gin.Context) {
	id := c.Param("id")

	existing, err := m.RepoProdukGaleri.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}
	gambar := existing.Gambar
	if err := m.RepoProdukGaleri.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	helpers.DeleteImage(gambar)
	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (m *MasterProdukController) DeleteVariant(c *gin.Context) {
	id := c.Param("id")
	if err := m.RepoProdukVariant.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}