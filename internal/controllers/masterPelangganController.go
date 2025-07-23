package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MasterPelangganController struct {
	Repo repositories.MasterPelangganRepository
}

func NewMasterPelangganController(repo repositories.MasterPelangganRepository) *MasterPelangganController {
	return &MasterPelangganController{Repo: repo}
}

func (mp *MasterPelangganController) GetAll(c *gin.Context) {
	data, err := mp.Repo.GetAll()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp *MasterPelangganController) GetByID(c *gin.Context) {
	id := c.Param("id")
	data, err := mp.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mp * MasterPelangganController) Create(c *gin.Context) {
	var req request.CreatePelangganRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	data := models.MasterPelanggan{
		IDUser: uuid.MustParse(req.IDUser),
		NamaLengkap: req.NamaLengkap,
		NamaPanggilan: req.NamaPanggilan,
		Phone: req.Phone,
	}

	if err := mp.Repo.Create(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterPelangganController) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := mp.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	var req request.UpdatePelangganRequest
	req.ID = id
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	data := models.MasterPelanggan{
		ID: existing.ID,
		IDUser: uuid.MustParse(req.IDUser),
		NamaLengkap: req.NamaLengkap,
		NamaPanggilan: req.NamaPanggilan,
		Phone: req.Phone,
	}

	if err := mp.Repo.Update(&data); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (mp *MasterPelangganController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := mp.Repo.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}