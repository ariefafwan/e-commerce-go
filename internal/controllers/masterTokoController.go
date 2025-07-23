package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/internal/request"
	"fmt"
	"net/http"

	"e-commerce-go/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		mimeType := req.Gambar.Header.Get("Content-Type")
		if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
			helpers.Error(c, http.StatusBadRequest, nil, "File logo harus berupa gambar JPEG/JPG atau PNG")
			return
		}
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
		NomorToko:   req.NomorToko,
		AturanPajak: req.AturanPajak,
	}
	
	if err := pkg.DB.Transaction(func(tx *gorm.DB) error {
		return mt.Repo.Update(id, updatedModel)
	}); err != nil {
		helpers.DeleteImage(existing.Gambar)
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal mengupdate data")
		return
	}
	
	helpers.Success(c, http.StatusOK, nil, "Success")
}