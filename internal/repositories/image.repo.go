package repositories

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/internal/services"
)

type ImageHandler struct {
	imageService *services.ImageService // or *services.S3Service
}

func NewImageHandler(imageService *services.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (h *ImageHandler) UploadImages(c *gin.Context) {
	// Parse multipart form
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
		// Validate file type
		if !isValidImageType(fileHeader.Filename) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid file type: %s", fileHeader.Filename),
			})
			return
		}

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer file.Close()

		// Upload to cloud storage
		url, err := h.imageService.UploadImage(c.Request.Context(), file, fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		uploadedURLs = append(uploadedURLs, url)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Images uploaded successfully",
		"urls":    uploadedURLs,
	})
}

func isValidImageType(filename string) bool {
	validTypes := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	filename = strings.ToLower(filename)

	for _, ext := range validTypes {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}
