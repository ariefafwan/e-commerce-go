package pkg

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cloud *cloudinary.Cloudinary

func InitCloudinary() {
	cld, err := cloudinary.NewFromParams(
		GetEnv("CLOUDINARY_CLOUD_NAME", "sanbercode"),
		GetEnv("CLOUDINARY_API_KEY", "123456789"),
		GetEnv("CLOUDINARY_API_SECRET", "123456789"),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	Cloud = cld
}