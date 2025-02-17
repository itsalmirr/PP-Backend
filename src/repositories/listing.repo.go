package repositories

import (
	"errors"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
)

// CreateListingRepository creates a new listing record in the database.
// It first checks if a listing with the given title or address already exists.
// If such a listing exists, it returns an error.
// If not, it creates a new listing record with the provided data.
// The operation is performed within a transaction to ensure atomicity.
//
// Parameters:
//   - data: CreateListingInput containing the details of the listing to be created.
//
// Returns:
//   - error: An error if the listing already exists or if the creation fails, otherwise nil.
func CreateListingRepo(data models.Listing) error {
	var existingListing models.Listing
	if err := config.DB.Where("title = ? OR address = ?", data.Title, data.Address).First(&existingListing).Error; err == nil {
		return errors.New("listing with the given title or address already exists")

	}

	// Create a new listing
	listing := models.Listing{
		Title:          data.Title,
		Address:        data.Address,
		City:           data.City,
		State:          data.State,
		ZipCode:        data.ZipCode,
		Description:    data.Description,
		Price:          data.Price,
		Bedroom:        data.Bedroom,
		Bathroom:       data.Bathroom,
		Garage:         data.Garage,
		Sqft:           data.Sqft,
		TypeOfProperty: data.TypeOfProperty,
		LotSize:        data.LotSize,
		Pool:           data.Pool,
		YearBuilt:      data.YearBuilt,
		Media:          data.Media,
		Status:         data.Status,
		RealtorID:      data.RealtorID,
	}

	// Start a new transaction
	tx := config.DB.Begin()
	if err := tx.Create(&listing).Error; err != nil {
		return errors.New("failed to create listing")
	}
	tx.Commit()
	return nil
}

// GetListingsRepo retrieves a paginated list of property listings from the database.
//
// The function performs two database operations:
// 1. Gets the total count of all listings in the database
// 2. Retrieves a specific page of listings based on the provided parameters
//
// Parameters:
//   - page: The page number to retrieve (1-based indexing)
//   - limit: The maximum number of listings to return per page
//
// Returns:
//   - []models.Listing: A slice of Listing models containing the paginated results
//   - int64: The total number of listings in the database (before pagination)
//   - error: An error if the database operations fail, nil otherwise
//
// Example usage:
//
//	listings, total, err := GetListingsRepo(1, 10) // Get first page with 10 items
//
// Note: The function uses zero-based offset pagination internally but accepts
// one-based page numbers for better usability.
func GetListingsRepo(page, limit int) ([]models.Listing, int64, error) {
	var listings []models.Listing
	var total int64

	// Calculate the offset based on the page number and limit
	// For example: page 1 with limit 10 = offset 0, page 2 = offset 10
	offset := (page - 1) * limit

	// Get the total count of listings first
	if err := config.DB.Model(&models.Listing{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve the paginated listings using offset and limit
	result := config.DB.Model(&models.Listing{}).Offset(offset).Limit(limit).Find(&listings)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return listings, total, nil
}
