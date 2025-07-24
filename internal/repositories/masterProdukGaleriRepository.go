package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterProdukGaleriRepository interface {
	GetAll(id string) ([]dto.MasterProdukGaleriResponse, error)
	GetByID(id string) (*dto.MasterProdukGaleriResponse, error)
	Create(galeri *models.MasterProdukGaleri) error
	Update(galeri *models.MasterProdukGaleri) error
	Delete(id string) error
}

type masterProdukGaleriRepo struct {
	db *gorm.DB
}

func NewMasterProdukGaleriRepository(db *gorm.DB) MasterProdukGaleriRepository {
	return &masterProdukGaleriRepo{db}
}

func (m *masterProdukGaleriRepo) GetAll(id string) ([]dto.MasterProdukGaleriResponse, error) {
	var galeri []models.MasterProdukGaleri
	err := m.db.Find(&galeri, "id_produk = ?", id).Error
	if err != nil {
		return nil, err
	}
	var galeriResponse []dto.MasterProdukGaleriResponse
	err = copier.Copy(&galeriResponse, &galeri)
	if err != nil {
		return nil, err
	}
	return galeriResponse, err
}

func (m *masterProdukGaleriRepo) GetByID(id string) (*dto.MasterProdukGaleriResponse, error) {
	var galeri models.MasterProdukGaleri
	err := m.db.First(&galeri, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	var galeriResponse dto.MasterProdukGaleriResponse
	err = copier.Copy(&galeriResponse, &galeri)
	if err != nil {
		return nil, err
	}
	return &galeriResponse, err
}

func (m *masterProdukGaleriRepo) Create(galeri *models.MasterProdukGaleri) error {
	var maxUrutan uint8
	if galeri.Urutan == 0 {
		m.db.Model(&models.MasterProdukGaleri{}).
				Where("id_produk = ?", galeri.IDProduk).
				Select("COALESCE(MAX(urutan), 0)").
				Scan(&maxUrutan)
		galeri.Urutan = maxUrutan + 1
	}
	
	return m.db.Create(galeri).Error
}

func (m *masterProdukGaleriRepo) Update(galeri *models.MasterProdukGaleri) error {
	var data models.MasterProdukGaleri
	err := m.db.First(&data, "id = ?", galeri.ID).Error
	if err != nil {
		return err
	}

	if data.Urutan != galeri.Urutan {
		var count int64
		m.db.Model(&models.MasterProdukGaleri{}).Where("id_produk = ? AND id != ? AND urutan = ?", galeri.IDProduk, galeri.ID, galeri.Urutan).Count(&count)
		if count > 0 {
			return errors.New("urutan sudah digunakan")
		}
	}
	return m.db.Model(&models.MasterProdukGaleri{}).Where("id = ?", galeri.ID).Updates(galeri).Error
}

var ErrorProdukGaleri = errors.New("produk harus memiliki setidaknya 1 gambar, silahkan update jika ingin merubah gambar ini")

func (m *masterProdukGaleriRepo) Delete(id string) error {
	var count int64
	var data models.MasterProdukGaleri
	err := m.db.First(&data, "id = ?", id).Error
	if err != nil {
		return err
	}
	m.db.Model(&models.MasterProdukGaleri{}).Where("id_produk = ? AND id != ?", data.IDProduk, data.ID).Count(&count)
	if count < 1 {
		return ErrorProdukGaleri
	}
	
	return m.db.Delete(&models.MasterProdukGaleri{}, "id = ?", id).Error
}