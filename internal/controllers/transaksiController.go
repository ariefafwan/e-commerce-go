package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransaksiController struct {
	RepoTransaksi repositories.TransaksiRepository
	RepoTransaksiItem repositories.TransaksiItemRepository
}

func NewTransaksiController(repoTransaksi repositories.TransaksiRepository, 
							repoTransaksiItem repositories.TransaksiItemRepository) *TransaksiController {
	return &TransaksiController{RepoTransaksi : repoTransaksi,
							RepoTransaksiItem : repoTransaksiItem}
}

func (t *TransaksiController) GetAll(c *gin.Context) {
	meta := helpers.ParseQueryParams(c)

	data, total, err := t.RepoTransaksi.GetAll(meta)
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

func (t *TransaksiController) GetAllItem(c *gin.Context) {
	id_transaksi := c.Param("id")

	data, err := t.RepoTransaksiItem.GetAll(id_transaksi)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	
	helpers.Success(c, http.StatusOK, data, "Success")
}

func (t *TransaksiController) GetAllByPelanggan(c *gin.Context) {
	id_pelanggan := c.Param("id")
	meta := helpers.ParseQueryParams(c)

	data, total, err := t.RepoTransaksi.GetAllByPelanggan(id_pelanggan, meta)
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

func (t *TransaksiController) GetByID(c *gin.Context) {
	id := c.Param("id")
	data, err := t.RepoTransaksi.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	helpers.Success(c, http.StatusOK, data, "Success")
}

func (t *TransaksiController) KalkulasiTransaksi(c *gin.Context) {
	var req request.KalkulasiTransaksiRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Input tidak valid")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	data, err := t.RepoTransaksi.KalkulasiTransaksi(req.IDItems, req.IDAlamatPelanggan)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}