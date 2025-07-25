package request

import (
	"e-commerce-go/pkg"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	// Register custom validator: fk_exists
	validate.RegisterValidation("fk_exists", func(fl validator.FieldLevel) bool {
		param := fl.Param() // "users|id"
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			return false
		}
		table, column := parts[0], parts[1]
		value := fl.Field().Interface()
		if value == nil || value == "" {
			return true
		}

		query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column)
		var count int64
		db := pkg.DB
		err := db.Raw(query, value).Scan(&count).Error
		if err != nil || count == 0 {
			return false
		}
		return true
	})

	// Register custom validator: unique
	validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		param := fl.Param() // "users|email"
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			return false
		}
		table, column := parts[0], parts[1]
		value := fl.Field().Interface()
		if value == nil || value == "" {
			return true
		}

		query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column)
		var count int64
		db := pkg.DB
		err := db.Raw(query, value).Scan(&count).Error
		if err != nil || count > 0 {
			return false
		}
		return true
	})

	// Register custom validator: unique_except
	validate.RegisterValidation("unique_except", func(fld validator.FieldLevel) bool {
		params := fld.Param()
		parts := strings.Split(params, ":")
		if len(parts) != 3 {
			return false
		}
		table, column, id := parts[0], parts[1], parts[2]
		value := fld.Field().Interface()

		parent := fld.Parent()
		excludeID := parent.FieldByNameFunc(func(name string) bool {
			return strings.EqualFold(name, id)
		})
		if !excludeID.IsValid() {
			return false
		}
		
		if value == nil || value == "" {
			return false
		}

		var count int64
		db := pkg.DB
		err := db.Table(table).Where(column + " = ? AND id != ?", value, excludeID.Interface()).Scan(&count).Error
		if err != nil || count > 0 {
			return false
		}
		return true
	})

	validate.RegisterValidation("date_format", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		if dateStr == "" {
			return true
		}
		
		_, err := time.Parse("02-01-2006", dateStr)
		return err == nil
	})
}

func ValidateStruct(data interface{}) map[string]string {	
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errors[field] = field + " wajib diisi"
		case "min":
			errors[field] = field + " minimal sepanjang " + e.Param()
		case "max":
			errors[field] = field + " maksimal sepanjang " + e.Param()
		case "fk_exists":
			errors[field] = field + " tidak ditemukan di database"
		case "unique":
			errors[field] = field + " sudah digunakan"
		case "unique_except":
			errors[field] = field + " sudah digunakan"
		case "date_format":
			errors[field] = field + " harus berformat DD-MM-YYYY (contoh: 27-09-2025)"
		default:
			errors[field] = "Field " + strings.ToLower(field) + " tidak valid"
		}
	}
	return errors
}
