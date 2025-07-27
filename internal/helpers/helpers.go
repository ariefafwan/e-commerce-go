package helpers

import (
	"crypto/sha512"
	"e-commerce-go/external/cloudinary"
	"e-commerce-go/internal/repositories"
	"e-commerce-go/pkg"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}

func UploadImage(file *multipart.FileHeader, folder string) (string, error) {
	mimeType := file.Header.Get("Content-Type")
	if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
		return "", errors.New("file logo harus berupa gambar .jpeg/jpg atau .png")
	}
	
	return cloudinary.UploadImage(file, folder)
}

func DeleteImage(publicID string) error {
	return cloudinary.DeleteImage(publicID)
}

func ParseQueryParams(c *gin.Context) repositories.QueryParams {
	search := c.DefaultQuery("search", "")
	sort := strings.ToLower(c.DefaultQuery("sort", "asc"))
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	return repositories.QueryParams{
		Page:   page,
		Limit:  limit,
		Search: search,
		Sort:   sort,
	}
}

func VerifySignatureMidtrans(orderID, statusCode, grossAmount, signatureKey string) bool {
	// Format: order_id+status_code+gross_amount+server_key
	serverKey := pkg.GetEnv("MIDTRANS_SERVER_KEY", "")
	payload := orderID + statusCode + grossAmount + serverKey
	
	// Hash dengan SHA512
	hash := sha512.Sum512([]byte(payload))
	expectedSignature := fmt.Sprintf("%x", hash)
	
	return expectedSignature == signatureKey
}