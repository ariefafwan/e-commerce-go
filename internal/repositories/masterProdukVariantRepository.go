package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"

	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterProdukVariantRepository interface {
	GetAll(slug_produk string) ([]dto.MasterProdukVariantResponse, error)
	CountAllByProduct(id_produk string) (int64, error)
	GetByID(id string) (*dto.MasterProdukVariantResponse, error)
	Create(masterProdukVariant *models.MasterProdukVariant) (error)
	Update(masterProdukVariant *models.MasterProdukVariant) (error)
	Delete(id string) error
}

type masterProdukVariantRepo struct {
	db *gorm.DB
}

func NewMasterProdukVariantRepository(db *gorm.DB) MasterProdukVariantRepository {
	return &masterProdukVariantRepo{db}
}

func (m *masterProdukVariantRepo) GetAll(slug_produk string) ([]dto.MasterProdukVariantResponse, error) {
	var masterProdukVariant []models.MasterProdukVariant
	var dataProduk models.MasterProduk
	err := m.db.Model(&models.MasterProduk{}).Where("slug = ?", slug_produk).First(&dataProduk).Error
	if err != nil {
		return nil, err
	}
	err = m.db.Find(&masterProdukVariant, "id_produk = ?", dataProduk.ID).Error
	if err != nil {
		return nil, err
	}
	var masterProdukVariantResponse []dto.MasterProdukVariantResponse
	err = copier.Copy(&masterProdukVariantResponse, &masterProdukVariant)
	if err != nil {
		return nil, err
	}
	return masterProdukVariantResponse, err
}

func (m *masterProdukVariantRepo) GetByID(id string) (*dto.MasterProdukVariantResponse, error) {
	var masterProdukVariant models.MasterProdukVariant
	err := m.db.First(&masterProdukVariant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	var masterProdukVariantResponse dto.MasterProdukVariantResponse
	err = copier.Copy(&masterProdukVariantResponse, &masterProdukVariant)
	if err != nil {
		return nil, err
	}
	return &masterProdukVariantResponse, err
}

func (m *masterProdukVariantRepo) CountAllByProduct(id_produk string) (int64, error) {
	var count int64
	err := m.db.Model(&models.MasterProdukVariant{}).Where("id_produk = ?", id_produk).Count(&count).Error
	return count, err
}

func (m *masterProdukVariantRepo) Create(masterProdukVariant *models.MasterProdukVariant) (error) {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var count int64
	tx.Model(&models.MasterProdukVariant{}).Where("id_produk = ? AND nama_variant = ?", masterProdukVariant.IDProduk, masterProdukVariant.NamaVariant).Count(&count)
	if count > 0 {
		tx.Rollback()
		return errors.New("nama variant sudah ada")
	}
	var dataProduk models.MasterProduk
	err := tx.First(&dataProduk, "id = ?", masterProdukVariant.IDProduk).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if masterProdukVariant.Harga < dataProduk.MinHarga {
		return errors.New("harga tidak boleh kurang dari minimal harga produk")
	} else if masterProdukVariant.Harga > dataProduk.MaxHarga {
		return errors.New("harga tidak boleh lebih dari maximal harga produk")
	}

	err = tx.Create(&masterProdukVariant).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m *masterProdukVariantRepo) Update(masterProdukVariant *models.MasterProdukVariant) (error) {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var count int64
	tx.Model(&models.MasterProdukVariant{}).Where("id_produk = ? AND id != ? AND nama_variant = ?", masterProdukVariant.IDProduk, masterProdukVariant.ID, masterProdukVariant.NamaVariant).Count(&count)
	if count > 0 {
		tx.Rollback()
		return errors.New("nama variant sudah ada")
	}

	var dataProduk models.MasterProduk
	err := tx.First(&dataProduk, "id = ?", masterProdukVariant.IDProduk).Error
	if err != nil {
		tx.Rollback()
		return errors.New("produk tidak ditemukan")
	}
	if masterProdukVariant.Harga < dataProduk.MinHarga {
		tx.Rollback()
		return errors.New("harga tidak boleh kurang dari minimal harga produk")
	} else if masterProdukVariant.Harga > dataProduk.MaxHarga {
		tx.Rollback()
		return errors.New("harga tidak boleh lebih dari maximal harga produk")
	}

	err = tx.Model(&models.MasterProdukVariant{}).Where("id = ?", masterProdukVariant.ID).Updates(&masterProdukVariant).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

var ErrorProdukVariant = errors.New("produk harus memiliki setidaknya 1 variant, silahkan update jika ingin merubah data ini")

func (m *masterProdukVariantRepo) Delete(id string) error {
	var count int64
	var data models.MasterProdukVariant
	err := m.db.First(&data, "id = ?", id).Error
	if err != nil {
		return err
	}
	m.db.Model(&models.MasterProdukVariant{}).Where("id_produk = ? AND id != ?", data.IDProduk, data.ID).Count(&count)
	if count < 1 {
		return ErrorProdukVariant
	}
	return m.db.Delete(&models.MasterProdukVariant{}, "id = ?", id).Error
}