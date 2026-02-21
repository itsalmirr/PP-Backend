package repositories

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/internal/services"
	"ppgroup.ppgroup.com/internal/util"
)

type ImageHandler struct {
	imageService services.ImageUploader
}

func NewImageHandler(imageService services.ImageUploader) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (h *ImageHandler) UploadImages(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := c.Request.MultipartForm.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images provided"})
		return
	}

	var uploadedURLs []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer file.Close()

		// Validate by MIME type (magic bytes), not extension
		valid, err := util.ValidateImageFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate file"})
			return
		}
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid file type: %s. Accepted: JPEG, PNG, GIF, WebP", fileHeader.Filename),
			})
			return
		}

		url, err := h.imageService.UploadImage(c.Request.Context(), file, fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}

		uploadedURLs = append(uploadedURLs, url)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Images uploaded successfully",
		"urls":    uploadedURLs,
	})
}
