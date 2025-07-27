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

func (t *TransaksiController) Create(c *gin.Context) {
	var req request.CreateTransaksiRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Input tidak valid")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	paymentResponse, err := t.RepoTransaksi.Create(req.IDItems, req.IDAlamatPelanggan, req.Layanan, req.Notes)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, paymentResponse, "Transaksi berhasil dibuat. Silakan lakukan pembayaran pada link yang tersedia.")
}

func (t *TransaksiController) UpdateStatus(c *gin.Context) {
	id_transaksi := c.Param("id")
	var req request.UpdateStatusTransaksiRequest
	if err := c.ShouldBind(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Input tidak valid")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
		return
	}

	err := t.RepoTransaksi.UpdateStatus(id_transaksi, req.Status)
	if err != nil {
		helpers.Error(c, http.StatusUnprocessableEntity, err.Error(), "Failed")
		return
	}

	helpers.Success(c, http.StatusOK, nil, "Success")
}

func (t *TransaksiController) MidtransCallback(c *gin.Context) {
	var callbackData struct {
		OrderID           string `json:"order_id"`
		TransactionStatus string `json:"transaction_status"`
		FraudStatus       string `json:"fraud_status"`
		StatusCode        string `json:"status_code"`
		GrossAmount       string `json:"gross_amount"`
		PaymentType       string `json:"payment_type"`
		TransactionTime   string `json:"transaction_time"`
		SignatureKey      string `json:"signature_key"`
	}

	if err := c.ShouldBindJSON(&callbackData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	if !helpers.VerifySignatureMidtrans(callbackData.OrderID,
		callbackData.StatusCode, callbackData.GrossAmount, callbackData.SignatureKey) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	// callback
	err := t.RepoTransaksi.HandlePaymentCallback(
		callbackData.OrderID,
		callbackData.TransactionStatus,
		callbackData.FraudStatus,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	// semua response di serderhanakan karena untuk midtrans
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}