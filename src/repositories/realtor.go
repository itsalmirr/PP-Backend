package repositories

import (
	"errors"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
)

type CreateRealtorInput struct {
	FullName    string `json:"full_name" binding:"required"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required"`
	IsMVP       bool   `json:"is_mvp"`
}

func CreateRealtorRepository(data CreateRealtorInput) error {
	// check if realtor already exists
	var existingRealtor models.Realtor
	if err := config.DB.Where("email = ? OR phone = ?", data.Email, data.Phone).
		First(&existingRealtor).Error; err == nil {
		return errors.New("realtor with the given email or phone already exists")
	}

	// Create a new realtor
	realtor := models.Realtor{
		FullName:    data.FullName,
		Photo:       data.Photo,
		Description: data.Description,
		Phone:       data.Phone,
		Email:       data.Email,
		IsMVP:       data.IsMVP,
	}

	// Start a new transaction
	tx := config.DB.Begin()
	if err := tx.Create(&realtor).Error; err != nil {
		return errors.New("Failed to create realtor")
	}
	tx.Commit()
	return nil
}
