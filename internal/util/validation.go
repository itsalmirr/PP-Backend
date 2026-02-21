package util

import (
	"io"
	"mime/multipart"
	"net/http"
)

var validImageMIMETypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
	"image/webp": true,
}

// ValidateImageFile checks the actual MIME type of an uploaded file
// by reading the first 512 bytes (magic bytes), rather than trusting
// the file extension which can be spoofed.
func ValidateImageFile(file multipart.File) (bool, error) {
	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return false, err
	}
	if n == 0 {
		return false, nil
	}

	// Reset file pointer for subsequent reads
	if _, err := file.Seek(0, 0); err != nil {
		return false, err
	}

	mimeType := http.DetectContentType(buf[:n])
	return validImageMIMETypes[mimeType], nil
}
