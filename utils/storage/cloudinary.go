package storage

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	godotenv.Load()
	cldSecret := os.Getenv("API_SECRET")
	cldName := os.Getenv("API_NAME")
	cldKey := os.Getenv("API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := SetupCloudinary()
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.URL, nil
}
