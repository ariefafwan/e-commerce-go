package helpers

import (
	"e-commerce-go/external/cloudinary"
	"log"
	"mime/multipart"

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
	return cloudinary.UploadImage(file, folder)
}

func DeleteImage(publicID string) error {
	return cloudinary.DeleteImage(publicID)
}