package repositories

import (
	"errors"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
	"github.com/google/uuid"
)

type CreateListingInput struct {
	Title          string  `json:"title"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	ZipCode        string  `json:"zip_code"`
	Description    string  `json:"description"`
	Price          string  `json:"price"`
	Bedroom        int     `json:"bedroom"`
	Bathroom       float32 `json:"bathroom"`
	Garage         int     `json:"garage"`
	Sqft           int64   `json:"sqft"`
	TypeOfProperty string  `json:"type_of_property"`
	LotSize        int64   `json:"lot_size"`
	Pool           bool    `json:"pool"`
	YearBuilt      string  `json:"year_built"`

	PhotoMain string `json:"photo_main"`
	Photo1    string `json:"photo_1"`
	Photo2    string `json:"photo_2"`
	Photo3    string `json:"photo_3"`
	Photo4    string `json:"photo_4"`
	Photo5    string `json:"photo_5"`

	IsPublished bool      `json:"is_published"`
	RealtorID   uuid.UUID `json:"realtor_id"`
}

func CreateListingRepository(data CreateListingInput) error {
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
		PhotoMain:      data.PhotoMain,
		Photo1:         data.Photo1,
		Photo2:         data.Photo2,
		Photo3:         data.Photo3,
		Photo4:         data.Photo4,
		Photo5:         data.Photo5,
		IsPublished:    data.IsPublished,
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
