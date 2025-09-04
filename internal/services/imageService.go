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
	// Add validation for empty credentials
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		panic("Cloudinary credentials are required: cloudName, apiKey, apiSecret")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Cloudinary: %v", err))
	}

	return &ImageService{
		cloudinary: cld,
	}
}

func (s *ImageService) UploadImage(ctx context.Context, file multipart.File, filename string) (string, error) {
	// Reset file pointer to beginning
	file.Seek(0, 0)

	publicID := generatePublicID(filename)

	result, err := s.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: publicID,
		Folder:   "real-estate-listings",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image to cloudinary: %w", err)
	}

	// Debug: Check what we got back from Cloudinary
	if result == nil {
		return "", fmt.Errorf("cloudinary returned nil result")
	}

	if result.SecureURL == "" {
		return "", fmt.Errorf("cloudinary returned empty SecureURL, PublicID: %s, result: %+v", publicID, result)
	}

	return result.SecureURL, nil
}

func generatePublicID(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	return fmt.Sprintf("%s_%d", name, time.Now().Unix())
}
