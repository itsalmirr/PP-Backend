package repositories

import (
	"errors"

	"backend.com/go-backend/app/config"
	"backend.com/go-backend/app/models"
	"gorm.io/gorm"
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

// GetRealtorRepository retrieves a realtor record from the database based on the provided email.
// It returns a models.Realtor object and an error if any occurred during the query.
// If the realtor is not found, it returns an error indicating "user not found".
// If there is any other error during the query, it returns an error indicating "failed to get user".
func GetRealtorRepository(email string) (models.Realtor, error) {
	var realtor models.Realtor
	if err := config.DB.Select("id", "email", "full_name", "is_mvp").Where("email = ?", email).First(&realtor).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return realtor, errors.New("realtor not found")
		}
		return realtor, errors.New("failed to get realtor")
	}
	return realtor, nil
}
