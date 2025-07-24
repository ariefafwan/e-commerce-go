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

type MasterKategoriProdukController struct {
	Repo repositories.MasterKategoriProdukRepository
}

func NewMasterKategoriProdukController(repo repositories.MasterKategoriProdukRepository) *MasterKategoriProdukController {
	return &MasterKategoriProdukController{Repo: repo}
}

func (mkp *MasterKategoriProdukController) GetAll(c *gin.Context) {
	meta := helpers.ParseQueryParams(c)

	data, total, err := mkp.Repo.GetAll(meta)
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

func (mkp *MasterKategoriProdukController) GetByID(c *gin.Context) {
	data, err := mkp.Repo.GetByID(c.Param("id"))
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}
	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mkp *MasterKategoriProdukController) Create(c *gin.Context) {
	var req request.CreateKategoriProdukRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if err := request.ValidateStruct(req); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err, "Validasi gagal")
		return
	}

	data := models.MasterKategoriProduk{
		Nama: req.Nama,
	}

	if req.IDParent != nil {
		var idParent uuid.UUID
		idParent, _ = uuid.Parse(*req.IDParent)
		data.IDParent = &idParent
	}

	if err := mkp.Repo.Create(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mkp *MasterKategoriProdukController) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := mkp.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	var req request.UpdateKategoriProdukRequest
	req.ID = id
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if err := request.ValidateStruct(req); err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err, "Validasi gagal")
		return
	}

	if req.IDParent != nil && id == *req.IDParent {
		helpers.Error(c, http.StatusUnprocessableEntity, "ID Parent Tidak Boleh Sama Dengan ID Kategori", "Validasi gagal")
		return
	}

	data := models.MasterKategoriProduk{
		ID:     existing.ID,
		Nama:   req.Nama,
		Urutan: existing.Urutan,
	}

	if req.IDParent != nil {
		var idParent uuid.UUID
		idParent, _ = uuid.Parse(*req.IDParent)
		data.IDParent = &idParent
	}

	if req.Urutan != nil {
		data.Urutan = *req.Urutan
	}

	if err := mkp.Repo.Update(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mkp *MasterKategoriProdukController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := mkp.Repo.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}