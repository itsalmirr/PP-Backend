package util

import (
	"bytes"
	"io"
	"testing"
)

// fakeFile wraps a bytes.Reader to implement multipart.File
type fakeFile struct {
	*bytes.Reader
}

func (f *fakeFile) Close() error { return nil }

func newFakeFile(data []byte) *fakeFile {
	return &fakeFile{Reader: bytes.NewReader(data)}
}

func (f *fakeFile) Seek(offset int64, whence int) (int64, error) {
	return f.Reader.Seek(offset, whence)
}

func (f *fakeFile) ReadAt(p []byte, off int64) (n int, err error) {
	return f.Reader.ReadAt(p, off)
}

// JPEG magic bytes
var jpegHeader = []byte{0xFF, 0xD8, 0xFF, 0xE0}

// PNG magic bytes
var pngHeader = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

// GIF magic bytes
var gifHeader = []byte("GIF89a")

func paddedFile(header []byte) *fakeFile {
	data := make([]byte, 512)
	copy(data, header)
	return newFakeFile(data)
}

func TestValidateImageFile_JPEG(t *testing.T) {
	file := paddedFile(jpegHeader)
	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !valid {
		t.Error("expected JPEG to be valid")
	}

	// Verify file pointer was reset
	pos, _ := file.Seek(0, io.SeekCurrent)
	if pos != 0 {
		t.Errorf("expected file pointer at 0, got %d", pos)
	}
}

func TestValidateImageFile_PNG(t *testing.T) {
	file := paddedFile(pngHeader)
	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !valid {
		t.Error("expected PNG to be valid")
	}
}

func TestValidateImageFile_GIF(t *testing.T) {
	file := paddedFile(gifHeader)
	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !valid {
		t.Error("expected GIF to be valid")
	}
}

func TestValidateImageFile_TextFile(t *testing.T) {
	data := make([]byte, 512)
	copy(data, []byte("this is a text file, not an image"))
	file := newFakeFile(data)

	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if valid {
		t.Error("expected text file to be invalid")
	}
}

func TestValidateImageFile_ExecutableFile(t *testing.T) {
	// ELF header
	data := make([]byte, 512)
	copy(data, []byte{0x7F, 0x45, 0x4C, 0x46})
	file := newFakeFile(data)

	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if valid {
		t.Error("expected executable to be invalid")
	}
}

func TestValidateImageFile_SmallFile(t *testing.T) {
	// File smaller than 512 bytes — Read returns io.EOF with valid data
	file := newFakeFile(jpegHeader) // only 4 bytes
	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error for small file: %v", err)
	}
	if !valid {
		t.Error("expected small JPEG to be valid")
	}
}

func TestValidateImageFile_EmptyFile(t *testing.T) {
	file := newFakeFile([]byte{})
	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error for empty file: %v", err)
	}
	if valid {
		t.Error("expected empty file to be invalid")
	}
}

func TestValidateImageFile_SpoofedExtension(t *testing.T) {
	// A file named "image.jpg" but actually a text file
	data := make([]byte, 512)
	copy(data, []byte("#!/bin/bash\nrm -rf /"))
	file := newFakeFile(data)

	valid, err := ValidateImageFile(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if valid {
		t.Error("expected spoofed file to be invalid (text content, not image)")
	}
}
