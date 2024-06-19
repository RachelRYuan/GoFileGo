package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// UploadFile uploads a file to Cloudinary and returns the upload result.
func UploadFile(file multipart.File, fileType string) (*uploader.UploadResult, error) {
	cld, err := cloudinary.NewFromURL(GodotEnv("CLOUDINARY_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %w", err)
	}

	randString := GenerateRandomString(6)
	fileName := fmt.Sprintf("uploads/%s", randString)
	fmt.Println("File type:", fileType)

	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: fileName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}
	return result, nil
}

// GetFileUrl retrieves the URL of a file stored in Cloudinary using its public ID.
func GetFileUrl(fileId string) string {
	cld, err := cloudinary.NewFromURL(GodotEnv("CLOUDINARY_URL"))
	if err != nil {
		logError(fmt.Errorf("failed to initialize Cloudinary: %w", err))
		return ""
	}

	result, err := cld.Admin.Asset(context.Background(), admin.AssetParams{
		PublicID: fileId,
	})
	if err != nil {
		logError(fmt.Errorf("failed to get file URL: %w", err))
		return ""
	}
	return result.URL
}

// logError logs the error to the console.
func logError(err error) {
	fmt.Println("Error:", err)
}
