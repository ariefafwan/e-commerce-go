package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"math"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MasterAlamatPelangganController struct {
	Repo repositories.MasterAlamatPelangganRepository
}

func NewMasterAlamatPelangganController(repo repositories.MasterAlamatPelangganRepository) *MasterAlamatPelangganController {
	return &MasterAlamatPelangganController{Repo: repo}
}

func (ma *MasterAlamatPelangganController) GetAll(c *gin.Context) {
	meta := helpers.ParseQueryParams(c)

	data, total, err := ma.Repo.GetAll(meta)
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

func (ma *MasterAlamatPelangganController) GetAllByPelanggan(c *gin.Context) {
	id := c.Param("id")
	meta := helpers.ParseQueryParams(c)

	data, total, err := ma.Repo.GetAllByPelanggan(id, meta)
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

func (ma *MasterAlamatPelangganController) GetByID(c *gin.Context) {
	data, err := ma.Repo.GetByID(c.Param("id"))
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	helpers.Success(c, http.StatusOK, data, "Success")
}

func (ma *MasterAlamatPelangganController) Create(c *gin.Context) {
	var req request.AlamatPelangganRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if err := request.ValidateStruct(req); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err, "Validasi gagal")
		return
	}

	data := models.MasterAlamatPelanggan{
		IDPelanggan:   uuid.MustParse(req.IDPelanggan),
		Label:         req.Label,
		AlamatLengkap: req.AlamatLengkap,
		KodePos:       req.KodePos,
		IDProvinsi:    req.IDProvinsi,
		IDKota:        req.IDKota,
		IDKecamatan:   req.IDKecamatan,
		NomorPenerima: req.NomorPenerima,
		NamaPenerima:  req.NamaPenerima,
	}

	if err := ma.Repo.Create(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (ma *MasterAlamatPelangganController) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := ma.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	var req request.AlamatPelangganRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if err := request.ValidateStruct(req); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err, "Validasi gagal")
		return
	}

	data := models.MasterAlamatPelanggan{
		ID:            existing.ID,
		IDPelanggan:   uuid.MustParse(req.IDPelanggan),
		Label:         req.Label,
		AlamatLengkap: req.AlamatLengkap,
		KodePos:       req.KodePos,
		IDProvinsi:    req.IDProvinsi,
		IDKota:        req.IDKota,
		IDKecamatan:   req.IDKecamatan,
		NomorPenerima: req.NomorPenerima,
		NamaPenerima:  req.NamaPenerima,
		IsDefault:     existing.IsDefault,
	}

	if err := ma.Repo.Update(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (ma *MasterAlamatPelangganController) SetAlamatUtama(c *gin.Context) {
	id := c.Param("id")
	if err := ma.Repo.SetAlamatUtama(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (ma *MasterAlamatPelangganController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ma.Repo.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}