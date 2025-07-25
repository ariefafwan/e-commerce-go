package repositories

import (
	"e-commerce-go/external/raja_ongkir"
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type TransaksiRepository interface {
	GetAll(q QueryParams) ([]dto.TransaksiResponse, int64, error)
	GetAllByPelanggan(id_pelanggan string, q QueryParams) ([]dto.TransaksiResponse, int64, error)
	KalkulasiTransaksi(itemIDs []string, idAlamatPelanggan *string) (*dto.TransaksiResponse, error)
	GetByID(id string) (*dto.TransaksiResponse, error)
	Create(transaksi *models.Transaksi) error
	Update(transaksi *models.Transaksi) error
	UpdateStatus(id string, status string) error
}

type transaksiRepo struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepo{db}
}

func (m *transaksiRepo) GetAll(q QueryParams) ([]dto.TransaksiResponse, int64, error) {
	var data []models.Transaksi
	var total int64
	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := m.db.Model(&models.Transaksi{}).Preload("DataPelanggan").Preload("DataAlamat").Preload("DataItems")

	if q.Search != "" {
		query = query.Where("id_pelanggan IN (?)",
			m.db.Table("master_pelanggan").Select("id").Where("nama_lengkap ILIKE ?", "%"+q.Search+"%").Or("nama_panggilan ILIKE ?", "%"+q.Search+"%"),
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	var response []dto.TransaksiResponse
	if err := copier.Copy(&response, &data); err != nil {
		return nil, 0, err
	}

	return response, total, err
}

func (m *transaksiRepo) GetAllByPelanggan(id_pelanggan string,q QueryParams) ([]dto.TransaksiResponse, int64, error) {
	var data []models.Transaksi
	var total int64
	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := m.db.Model(&models.Transaksi{}).Where("id_pelanggan = ?", id_pelanggan).Preload("DataPelanggan").Preload("DataAlamat").Preload("DataItems")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	var response []dto.TransaksiResponse
	if err := copier.Copy(&response, &data); err != nil {
		return nil, 0, err
	}

	return response, total, err
}

func (m *transaksiRepo) GetByID(id string) (*dto.TransaksiResponse, error) {
	var data models.Transaksi
	err := m.db.Preload("DataPelanggan").Preload("DataAlamat").Preload("DataItems").First(&data, "id = ?", id).Error

	var response dto.TransaksiResponse
	if err := copier.Copy(&response, &data); err != nil {
		return nil, err
	}
	return &response, err
}

func (r *transaksiRepo) KalkulasiTransaksi(itemIDs []string, idAlamatPelanggan *string) (*dto.TransaksiResponse, error) {
	if len(itemIDs) == 0 {
		return nil, errors.New("tidak ada item yang dipilih untuk memulai transaksi")
	}

	var items []models.TransaksiKeranjangItem
	var idPelanggan uuid.UUID

	err := r.db.
		Preload("DataVariant.DataProduk").
		Preload("DataKeranjang.DataPelanggan").
		Where("id IN ?", itemIDs).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("tidak ada item yang valid ditemukan")
	}

	idPelanggan = items[0].DataKeranjang.IDPelanggan

	var totalHarga float64
	var beratTotal float64
	var jumlahItem int
	var responseItems []dto.TransaksiItemResponse

	for _, item := range items {
		
		subtotal := item.DataVariant.Harga * float64(item.Quantity)
		totalHarga += subtotal
		beratTotal += item.DataVariant.DataProduk.Berat * float64(item.Quantity)
		jumlahItem += item.Quantity

		var itemResponse dto.TransaksiItemResponse
		copier.Copy(&itemResponse, &item)
		itemResponse.Subtotal = subtotal
		copier.Copy(&itemResponse.DataVariant, &item.DataVariant)
		copier.Copy(&itemResponse.DataProduk, &item.DataVariant.DataProduk)

		responseItems = append(responseItems, itemResponse)
	}

	var toko models.MasterToko
	if err := r.db.First(&toko).Error; err != nil {
		return nil, errors.New("data master toko tidak ditemukan")
	}
	pajak := totalHarga * (toko.AturanPajak / 100.0)

	var alamat models.MasterAlamatPelanggan
	if idAlamatPelanggan != nil {
		if err := r.db.Preload("DataKecamatan.DataKota.DataProvinsi").First(&alamat, "id = ? AND id_pelanggan = ?", idAlamatPelanggan, idPelanggan).Error; err != nil {
			return nil, errors.New("alamat pelanggan yang dipilih tidak ditemukan atau bukan milik pelanggan ini")
		}
	} else {
		if err := r.db.Preload("DataKecamatan.DataKota.DataProvinsi").Where("is_default = ?", true).First(&alamat, "id_pelanggan = ?", idPelanggan).Error; err != nil {
			return nil, errors.New("alamat pelanggan yang dipilih tidak ditemukan atau bukan milik pelanggan ini")
		}
	}

	var pilihanOngkir []dto.PilihanOngkirResponse
	costResponse, err := raja_ongkir.CalculateShippingCost(toko.IDKecamatan, alamat.IDKecamatan, int(beratTotal))
    if err != nil {
        return nil, fmt.Errorf("gagal mengambil data ongkir: %v", err)
    }

    limitOpsi := 3
    if len(costResponse.Data) < limitOpsi {
        limitOpsi = len(costResponse.Data)
    }
    
    for i := 0; i < limitOpsi; i++ {
        opsi := costResponse.Data[i]
        pilihanOngkir = append(pilihanOngkir, dto.PilihanOngkirResponse{
            NamaLayanan: fmt.Sprintf("%s (%s)", opsi.Name, opsi.Service),
            Estimasi:    opsi.Etd,
            Harga:       opsi.Cost,
        })
    }

	response := &dto.TransaksiResponse{
		ID:                uuid.New(),
		IDPelanggan:       idPelanggan,
		IDAlamatPelanggan: alamat.ID,
		TotalHarga:        totalHarga,
		TotalOngkir:       0,
		PilihanOngkir:     &pilihanOngkir,
		JumlahItem:        int16(jumlahItem),
		BeratTotal:        beratTotal,
		Pajak:             pajak,
		GrandTotal:        totalHarga + pajak,
		Notes:             nil,
		Status:            "Pending",
		DataItems:         responseItems,
	}

	copier.Copy(&response.DataPelanggan, &items[0].DataKeranjang.DataPelanggan)
	copier.Copy(&response.DataAlamat, &alamat)

	return response, nil
}

func (m *transaksiRepo) Create(transaksi *models.Transaksi) error {
	return m.db.Create(transaksi).Error
}

func (m *transaksiRepo) Update(transaksi *models.Transaksi) error {
	return m.db.Model(&models.Transaksi{}).Where("id = ?", transaksi.ID).Updates(transaksi).Error
}

func (m *transaksiRepo) UpdateStatus(id string, status string) error {
	return m.db.Model(&models.Transaksi{}).Where("id = ?", id).Update("status", status).Error
}