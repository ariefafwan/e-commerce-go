package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type TransaksiKeranjangRepository interface {
	GetAll(q QueryParams) ([]dto.TransaksiKeranjangResponse, int64, error)
	GetAllByPelanggan(id_pelanggan string) ([]dto.TransaksiKeranjangResponse, error)
	GetByID(id string) (*dto.TransaksiKeranjangResponse, error)
	Create(id_pelanggan string, item models.TransaksiKeranjangItem) error
	Update(keranjang *models.TransaksiKeranjang) error
	Delete(id string) error
}

type transaksiKeranjangRepo struct {
	db *gorm.DB
}

func NewTransaksiKeranjangRepository(db *gorm.DB) TransaksiKeranjangRepository {
	return &transaksiKeranjangRepo{db}
}

func (t *transaksiKeranjangRepo) GetAll(q QueryParams) ([]dto.TransaksiKeranjangResponse, int64, error) {
	var keranjangs []models.TransaksiKeranjang
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := t.db.Model(&models.TransaksiKeranjang{}).Preload("DataPelanggan.DataUser").Preload("DataItems.DataProduk").Preload("DataItems.DataVariant")

	if q.Search != "" {
		query = query.Where("id_pelanggan IN (?)",
			t.db.Table("master_pelanggan").Select("id").Where("nama_lengkap ILIKE ?", "%"+q.Search+"%").Or("nama_panggilan ILIKE ?", "%"+q.Search+"%"),
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&keranjangs).Error
	if err != nil {
		return nil, 0, err
	}

	var keranjangResponse []dto.TransaksiKeranjangResponse
	if err = copier.Copy(&keranjangResponse, &keranjangs); err != nil {
		return nil, 0, err
	}
	return keranjangResponse, total, err
}

func (t *transaksiKeranjangRepo) GetAllByPelanggan(id_pelanggan string) ([]dto.TransaksiKeranjangResponse, error) {
	var keranjangs []models.TransaksiKeranjang
	var count int64
	err := t.db.Model(&models.MasterPelanggan{}).Where("id = ?", id_pelanggan).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("pelanggan tidak ditemukan")
	}

	err = t.db.Preload("DataPelanggan.DataUser").Preload("DataItems.DataProduk").Preload("DataItems.DataVariant").Where("id_pelanggan = ?", id_pelanggan).Find(&keranjangs).Error
	if err != nil {
		return nil, err
	}

	var keranjangResponse []dto.TransaksiKeranjangResponse
	if err = copier.Copy(&keranjangResponse, &keranjangs); err != nil {
		return nil, err
	}
	return keranjangResponse, err
}

func (t *transaksiKeranjangRepo) GetByID(id string) (*dto.TransaksiKeranjangResponse, error) {
	var keranjang models.TransaksiKeranjang
	err := t.db.Preload("DataPelanggan.DataUser").Preload("DataItems.DataProduk").Preload("DataItems.DataVariant").First(&keranjang, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	var keranjangResponse dto.TransaksiKeranjangResponse
	if err = copier.Copy(&keranjangResponse, &keranjang); err != nil {
		return nil, err
	}
	return &keranjangResponse, err
}

func (t *transaksiKeranjangRepo) Create(id_pelanggan string, item models.TransaksiKeranjangItem) error {
	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var data models.MasterProdukVariant
	err := tx.First(&data, "id = ?", item.IDVariantProduk).Error
	if err != nil {
		tx.Rollback()
		return errors.New("produk tidak ditemukan")
	}
	
	if data.Stok < item.Quantity {
		tx.Rollback()
		return errors.New("stok produk tidak mencukupi")
	}

	var count int64
	err = tx.Model(&models.TransaksiKeranjang{}).Where("id_pelanggan = ?", id_pelanggan).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var keranjang models.TransaksiKeranjang
	if count > 0 {
		err = tx.First(&keranjang, "id_pelanggan = ?", id_pelanggan).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		keranjang = models.TransaksiKeranjang{
			IDPelanggan: uuid.MustParse(id_pelanggan),
			BerlakuSampai: time.Now().Add(time.Hour *7 * 24),
		}
		err = tx.Create(&keranjang).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Model(&models.MasterProdukVariant{}).Where("id = ?", item.IDVariantProduk).Update("stok", gorm.Expr("stok - ?", item.Quantity)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	item.IDProduk = data.IDProduk
	item.IDKeranjang = keranjang.ID
	err = tx.Create(&item).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (t *transaksiKeranjangRepo) Update(keranjang *models.TransaksiKeranjang) error {
	return t.db.Model(&models.TransaksiKeranjang{}).Where("id = ?", keranjang.ID).Updates(keranjang).Error
}

func (t *transaksiKeranjangRepo) Delete(id string) error {
	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&models.TransaksiKeranjangItem{}).Where("id_keranjang = ?", id).Delete(&models.TransaksiKeranjangItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	err := tx.Delete(&models.TransaksiKeranjang{}, "id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
