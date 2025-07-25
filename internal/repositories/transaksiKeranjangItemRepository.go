package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"
	"math"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type TransaksiKeranjangItemRepository interface {
	GetByID(id string) (*dto.TransaksiKeranjangItemResponse, error)
	Update(item *models.TransaksiKeranjangItem) error
	Delete(id_keranjang string, id string) error
}

type transaksiKeranjangItemRepo struct {
	db *gorm.DB
}

func NewTransaksiKeranjangItemRepository(db *gorm.DB) TransaksiKeranjangItemRepository {
	return &transaksiKeranjangItemRepo{db}
}

func (t *transaksiKeranjangItemRepo) GetByID(id string) (*dto.TransaksiKeranjangItemResponse, error) {
	var data models.TransaksiKeranjangItem
	err := t.db.Preload("DataProduk").Preload("DataVariant").First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	var produkVariant dto.TransaksiKeranjangItemResponse
	if err = copier.Copy(&produkVariant, &data); err != nil {
		return nil, err
	}
	return &produkVariant, err
}

func (t *transaksiKeranjangItemRepo) Update(item *models.TransaksiKeranjangItem) error {
	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existing models.TransaksiKeranjangItem
	if err := tx.Preload("DataProduk").Preload("DataVariant").First(&existing, "id = ?", item.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	selisihQuantity := item.Quantity - existing.Quantity

    if selisihQuantity > 0 {
        updateResult := tx.Model(&models.MasterProdukVariant{}).
            Where("id = ? AND stok >= ?", item.IDVariantProduk, selisihQuantity).
            Update("stok", gorm.Expr("stok - ?", selisihQuantity))

        if updateResult.Error != nil {
            tx.Rollback()
            return updateResult.Error
        }

        if updateResult.RowsAffected == 0 {
            tx.Rollback()
            return errors.New("stok produk tidak mencukupi")
        }
    } else if selisihQuantity < 0 {
        stokDikembalikan := math.Abs(float64(selisihQuantity))
        err := tx.Model(&models.MasterProdukVariant{}).
            Where("id = ?", item.IDVariantProduk).
            Update("stok", gorm.Expr("stok + ?", stokDikembalikan)).Error
        if err != nil {
            tx.Rollback()
            return err
        }
    }

	err := tx.Model(&models.TransaksiKeranjangItem{}).Where("id = ?", item.ID).Update("quantity", item.Quantity).Error
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit().Error
}

func (t *transaksiKeranjangItemRepo) Delete(id_keranjang string, id string) error {
    tx := t.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    var item models.TransaksiKeranjangItem
    if err := tx.First(&item, "id = ?", id).Error; err != nil {
        tx.Rollback()
        return errors.New("item keranjang tidak ditemukan")
    }

    if err := tx.Model(&models.MasterProdukVariant{}).Where("id = ?", item.IDVariantProduk).Update("stok", gorm.Expr("stok + ?", item.Quantity)).Error; err != nil {
        tx.Rollback()
        return err
    }

	var count int64
    if err := tx.Model(&models.TransaksiKeranjangItem{}).Where("id != ? AND id_keranjang = ?", id, id_keranjang).Count(&count).Error; err != nil {
        tx.Rollback()
        return err
    }

	if count == 0 {
        if err := tx.Delete(&models.TransaksiKeranjang{}, "id = ?", id_keranjang).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    if err := tx.Where("id = ?", id).Delete(&models.TransaksiKeranjangItem{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit().Error
}