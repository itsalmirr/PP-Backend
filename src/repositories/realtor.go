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

// CreateRealtorRepository creates a new realtor record in the database.
// It first checks if a realtor with the given email or phone already exists.
// If such a realtor exists, it returns an error.
// If not, it creates a new realtor record with the provided data.
// The operation is performed within a transaction to ensure atomicity.
//
// Parameters:
//   - data: CreateRealtorInput containing the details of the realtor to be created.
//
// Returns:
//   - error: An error if the realtor already exists or if the creation fails, otherwise nil.
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
		return errors.New("failed to create realtor")
	}
	tx.Commit()
	return nil
}
