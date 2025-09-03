package repositories

import "ppgroup.m0chi.com/internal/services"

type ImageHandler struct {
	imageService *services.ImageService
}

func NewImageHandler(imageService *services.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}
