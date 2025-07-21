package helpers

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"e-commerce-go/pkg"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	_, err = pkg.Cloud.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder:   folder,
		PublicID: filename,
	})
	if err != nil {
		return "", err
	}

	return filename, nil
}

func DeleteImage(publicID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pkg.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return err
	}

	return nil
}