package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type ImageService struct {
	cloudinary *cloudinary.Cloudinary
}

func NewImageService(cloudName, apiKey, apiSecret string) *ImageService {
	cld, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	return &ImageService{
		cloudinary: cld,
	}
}

func (s *ImageService) UploadImage(ctx context.Context, file multipart.File, filename string) (string, error) {
	result, err := s.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: generatePublicID(filename),
		Folder:   "real-estate-listings",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	return result.SecureURL, nil
}

func generatePublicID(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	return fmt.Sprintf("%s_%d", name, time.Now().Unix())
}
