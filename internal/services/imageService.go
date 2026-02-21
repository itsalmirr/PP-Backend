package services

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// ImageUploader defines the interface for image upload operations.
type ImageUploader interface {
	UploadImage(ctx context.Context, file multipart.File, filename string) (string, error)
}

type ImageService struct {
	cloudinary *cloudinary.Cloudinary
}

func NewImageService(cloudName, apiKey, apiSecret string) (*ImageService, error) {
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return nil, errors.New("cloudinary credentials are required: cloudName, apiKey, apiSecret")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("initializing cloudinary: %w", err)
	}

	return &ImageService{
		cloudinary: cld,
	}, nil
}

func (s *ImageService) UploadImage(ctx context.Context, file multipart.File, filename string) (string, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return "", fmt.Errorf("resetting file pointer: %w", err)
	}

	publicID := generatePublicID(filename)

	result, err := s.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: publicID,
		Folder:   "real-estate-listings",
	})
	if err != nil {
		return "", fmt.Errorf("uploading image to cloudinary: %w", err)
	}

	if result == nil || result.SecureURL == "" {
		return "", fmt.Errorf("cloudinary upload failed: invalid response")
	}

	return result.SecureURL, nil
}

func generatePublicID(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	return fmt.Sprintf("%s_%d", name, time.Now().Unix())
}
