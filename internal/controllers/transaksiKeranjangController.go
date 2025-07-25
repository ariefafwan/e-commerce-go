package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransaksiKeranjangController struct {
	RepoKeranjang repositories.TransaksiKeranjangRepository
	RepoItem repositories.TransaksiKeranjangItemRepository
	RepoProductVariant repositories.MasterProdukVariantRepository
}

func NewTransaksiKeranjangController(
	repoKeranjang repositories.TransaksiKeranjangRepository,
	repoItem repositories.TransaksiKeranjangItemRepository,
	repoProductVariant repositories.MasterProdukVariantRepository,
) TransaksiKeranjangController {
	return TransaksiKeranjangController{
		RepoKeranjang: repoKeranjang,
		RepoItem: repoItem,
		RepoProductVariant: repoProductVariant,
	}
}

func (t *TransaksiKeranjangController) GetAllKeranjang(c *gin.Context) {
	meta := helpers.ParseQueryParams(c)

	data, total, err := t.RepoKeranjang.GetAll(meta)
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

func (t *TransaksiKeranjangController) GetAllByPelanggan(c *gin.Context) {
	id_pelanggan := c.Param("id")
	data, err := t.RepoKeranjang.GetAllByPelanggan(id_pelanggan)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, data, "Success")
}

func (t *TransaksiKeranjangController) GetByID(c *gin.Context) {
	id := c.Param("id")
	data, err := t.RepoKeranjang.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, data, "Success")
}

func (t* TransaksiKeranjangController) Create(c *gin.Context) {
	var req request.CreateKeranjangItemRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	data := models.TransaksiKeranjangItem{
		IDVariantProduk: uuid.MustParse(req.IDVariantProduk),
		Quantity: req.Quantity,
	}

	if err := t.RepoKeranjang.Create(req.IDPelanggan, data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (t *TransaksiKeranjangController) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := t.RepoKeranjang.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	var req request.UpdateKeranjangRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	if req.BerlakuSampai != nil {
		waktu, err := time.Parse("2006-01-02", *req.BerlakuSampai)
		if err != nil {
			helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Validasi gagal")
			return
		}
		existing.BerlakuSampai = waktu
	}

	data := models.TransaksiKeranjang{
		ID: existing.ID,
		IDPelanggan: uuid.MustParse(req.IDPelanggan),
		BerlakuSampai: existing.BerlakuSampai,
	}

	if err := t.RepoKeranjang.Update(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (t *TransaksiKeranjangController) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	existing, err := t.RepoItem.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	var req request.UpdateKeranjangItemRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if req.IDVariantProduk != existing.IDVariantProduk.String() {
		helpers.Error(c, http.StatusUnprocessableEntity, "ID Variant Produk Tidak Boleh Berubah", "Validasi gagal")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	data := models.TransaksiKeranjangItem{
		ID: existing.ID,
		IDKeranjang: existing.IDKeranjang,
		IDProduk: existing.IDProduk,
		IDVariantProduk: uuid.MustParse(req.IDVariantProduk),
		Quantity: req.Quantity,
	}

	if err := t.RepoItem.Update(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (t *TransaksiKeranjangController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := t.RepoKeranjang.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (t *TransaksiKeranjangController) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	existing, err := t.RepoItem.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}
	if err := t.RepoItem.Delete(existing.IDKeranjang.String(), id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, nil, "Success")
}

