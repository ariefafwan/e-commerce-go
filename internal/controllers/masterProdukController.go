package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"fmt"
	"math"
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
	meta := helpers.ParseQueryParams(c)

	data, total, err := mp.RepoProduk.GetAll(meta)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(meta.Limit)))

	response := gin.H{
		"data":        data,
		"total_items": total,
		"total_pages": totalPages,
		"current_page": meta.Page,
		"limit":        meta.Limit,
	}
	
	helpers.Success(c, http.StatusOK, response, "Success")
}

func (mp *MasterProdukController) GetAllByKategori(c *gin.Context) {
	slug := c.Param("slug")
	meta := helpers.ParseQueryParams(c)

	data, total, err := mp.RepoProduk.GetAllByKategori(slug, meta)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(meta.Limit)))

	response := gin.H{
		"data":        data,
		"total_items": total,
		"total_pages": totalPages,
		"current_page": meta.Page,
		"limit":        meta.Limit,
	}
	
	helpers.Success(c, http.StatusOK, response, "Success")
}

func (mp *MasterProdukController) GetProdukNonAktif(c *gin.Context) {
	meta := helpers.ParseQueryParams(c)

	data, total, err := mp.RepoProduk.GetProdukNonAktif(meta)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(meta.Limit)))

	response := gin.H{
		"data":        data,
		"total_items": total,
		"total_pages": totalPages,
		"current_page": meta.Page,
		"limit":        meta.Limit,
	}
	
	helpers.Success(c, http.StatusOK, response, "Success")
}

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
	slug_produk := c.Param("slug")
	data, err := mp.RepoProdukVariant.GetAll(slug_produk)
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
	slug_produk := c.Param("slug")
	data, err := mp.RepoProdukGaleri.GetAll(slug_produk)
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

	filename, err := helpers.UploadImage(req.Thumbnail, "Produk")
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
		return
	}

	data := models.MasterProduk{
		Nama:       req.Nama,
		MinHarga:   req.MinHarga,
		MaxHarga:   req.MaxHarga,
		Berat:      req.Berat,
		Deskripsi: 	req.Deskripsi,
		Thumbnail: 	filename,
		Status:     "Non Aktif",
	}

	if err := mp.RepoProduk.Create(&data, req.IDKategoriProduk); err != nil {
		errc := helpers.DeleteImage(filename)
		if errc != nil {
			helpers.Error(c, http.StatusInternalServerError, "Gagal Hapus Gambar", "Failed")
			return
		}
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) CreateVariant(c *gin.Context) {
	slug_produk := c.Param("slug")
	var req request.ProdukVariantRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}
	
	produk, err := mp.RepoProduk.GetBySlug(slug_produk)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Produk Tidak Ditemukan")
		return
	}

	count, err := mp.RepoProdukVariant.CountAllByProduct(produk.ID.String())
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	if count < 1 {
		if err := mp.RepoProduk.UpdateStatus(produk.ID.String(), "Aktif"); err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
			return
		}
	}

	data := models.MasterProdukVariant{
		IDProduk: 		 produk.ID,
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
	slug_produk := c.Param("slug")
	var req request.ProdukGaleriRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	produk, err := mp.RepoProduk.GetBySlug(slug_produk)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Produk Tidak Ditemukan")
		return
	}

	if req.Gambar == nil {
		helpers.Error(c, http.StatusBadRequest, nil, "Gambar harus diisi")
		return
	}

	filename, err := helpers.UploadImage(req.Gambar, "Produk-Galeri")
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
		return
	}

	data := models.MasterProdukGaleri{
		IDProduk:   produk.ID,
		Gambar:  	filename,
	}

	if req.Urutan != nil {
		data.Urutan = *req.Urutan
	} else {
		data.Urutan = 0
	}

	if err := mp.RepoProdukGaleri.Create(&data); err != nil {
		helpers.DeleteImage(fmt.Sprintf("Produk-Galeri/%s", filename))
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) Update(c *gin.Context) {
	var req request.ProdukRequest
	slug_produk := c.Param("slug")
	existing, err := mp.RepoProduk.GetBySlug(slug_produk)
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
		filename, err := helpers.UploadImage(req.Thumbnail, "Produk")
		if err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
			return
		}
		helpers.DeleteImage(fmt.Sprintf("Produk/%s", existing.Thumbnail))
		existing.Thumbnail = filename
	}

	data := models.MasterProduk{
		ID:         existing.ID,
		Nama:       req.Nama,
		MinHarga:   req.MinHarga,
		MaxHarga:   req.MaxHarga,
		Berat:      req.Berat,
		Deskripsi: 	req.Deskripsi,
		Thumbnail: 	existing.Thumbnail,
	}

	if req.Status != nil {
		data.Status = *req.Status
	}

	if err := mp.RepoProduk.Update(&data, req.IDKategoriProduk); err != nil {
		if req.Thumbnail != nil {
			helpers.DeleteImage(fmt.Sprintf("Produk/%s", existing.Thumbnail))
		}
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterProdukController) UpdateVariant(c *gin.Context) {
	var req request.ProdukVariantRequest
	slug_produk := c.Param("slug")
	id := c.Param("id")
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	produk, err := mp.RepoProduk.GetBySlug(slug_produk)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Produk Tidak Ditemukan")
		return
	}

	data := models.MasterProdukVariant{
		ID:       		 uuid.MustParse(id),
		NamaVariant:     req.NamaVariant,
		Harga:    		 req.Harga,
		Stok:    		 req.Stok,
		IDProduk: 		 produk.ID,
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
	slug_produk := c.Param("slug")
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	produk, err := m.RepoProduk.GetBySlug(slug_produk)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Produk Tidak Ditemukan")
		return
	}

	existing, err := m.RepoProdukGaleri.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	if req.Gambar != nil {
		filename, err := helpers.UploadImage(req.Gambar, "Produk-Galeri")
		if err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
			return
		}
		helpers.DeleteImage(fmt.Sprintf("Produk-Galeri/%s", existing.Gambar))
		existing.Gambar = filename
	}
	
	data := models.MasterProdukGaleri{
		ID:       uuid.MustParse(id),
		IDProduk: produk.ID,
		Gambar:   existing.Gambar,
	}

	if req.Urutan != nil {
		data.Urutan = *req.Urutan
	} else {
		data.Urutan = existing.Urutan
	}
	
	if err := m.RepoProdukGaleri.Update(&data); err != nil {
		if req.Gambar != nil {
			helpers.DeleteImage(fmt.Sprintf("Produk-Galeri/%s", existing.Gambar))
		}
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (m *MasterProdukController) Delete(c *gin.Context) {
	slug_produk := c.Param("slug")
	existing, err := m.RepoProduk.GetBySlug(slug_produk)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}
	gambar := existing.Thumbnail
	if err := m.RepoProduk.Delete(existing.ID.String()); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	helpers.DeleteImage(fmt.Sprintf("Produk/%s", gambar))
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
	helpers.DeleteImage(fmt.Sprintf("Produk-Galeri/%s", gambar))
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