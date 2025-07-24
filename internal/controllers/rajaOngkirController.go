package controllers

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RajaOngkirController struct {
	Repo repositories.RajaOngkirRepository
}

func NewRajaOngkirController(repo repositories.RajaOngkirRepository) RajaOngkirController {
	return RajaOngkirController{Repo: repo}
}

func (r *RajaOngkirController) GetProvinsi(c *gin.Context) {
	response, err := r.Repo.GetProvince()
	if err != nil {
		helpers.Error(c, http.StatusBadRequest, err, "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, response, "Success")
}

func (r *RajaOngkirController) GetKota(c *gin.Context) {
	id := c.Param("id")
	response, err := r.Repo.GetCity(id)
	if err != nil {
		helpers.Error(c, http.StatusBadRequest, err, "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, response, "Success")
}

func (r * RajaOngkirController) GetKecamatan(c *gin.Context) {
	id := c.Param("id")
	response, err := r.Repo.GetDistrict(id)
	if err != nil {
		helpers.Error(c, http.StatusBadRequest, err, "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, response, "Success")
}