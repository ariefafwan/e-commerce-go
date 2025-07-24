package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MasterTokoController struct {
	Repo repositories.MasterTokoRepository
}

func NewMasterTokoController(repo repositories.MasterTokoRepository) MasterTokoController {
	return MasterTokoController{Repo: repo}
}

func (mt *MasterTokoController) GetToko(c *gin.Context) {
	data, err := mt.Repo.GetToko()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, data, "Success")
}

func (mt *MasterTokoController) UpdateToko(c *gin.Context) {
	id := c.Param("id")
	existing, err := mt.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Data Tidak Ditemukan")
		return
	}

	var req request.UpdateTokoRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Failed")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
        helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
        return
    }

	if req.Gambar != nil {
		filename, err := helpers.UploadImage(req.Gambar, "Toko")
		if err != nil {
			helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal Saat Menyimpan Gambar")
			return
		}
		helpers.DeleteImage(fmt.Sprintf("Toko/%s", existing.Gambar))
		existing.Gambar = filename
	}

	var updatedModel = models.MasterToko{
		Nama:        req.Nama,
		Alamat:      req.Alamat,
		Gambar:      existing.Gambar,
		NoTelp:   	 req.NoTelp,
		IDProvinsi:    req.IDProvinsi,
		IDKota:        req.IDKota,
		IDKecamatan:   req.IDKecamatan,
		AturanPajak: req.AturanPajak,
	}
	
	if err := mt.Repo.Update(id, updatedModel); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal mengupdate data")
		helpers.DeleteImage(fmt.Sprintf("Toko/%s", existing.Gambar))
		return
	}
	
	helpers.Success(c, http.StatusOK, nil, "Success")
}