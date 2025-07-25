package cloudinary

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"time"

	"e-commerce-go/pkg"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"strings"
)

func UploadImage(file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	originalName := file.Filename
	
	ext := filepath.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)
	
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleanName := reg.ReplaceAllString(nameWithoutExt, "_")
	cleanName = strings.Trim(cleanName, "_")

	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), cleanName, ext)
	publicID := strings.Split(filename, ".")[0]

	_, err = pkg.Cloud.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder:   folder,
		PublicID: publicID,
	})
	if err != nil {
		return "", err
	}

	return filename, nil
}

func DeleteImage(publicID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	publicID = strings.Split(publicID, ".")[0]

	_, err := pkg.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return err
	}

	return nil
}