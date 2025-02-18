package repositories

import (
	"errors"
	"fmt"
	"time"

	"backend.com/go-backend/app/config"
	"backend.com/go-backend/app/models"
)

// Listing alias for brevity
type Listing = models.Listing

// ListingQueryParams holds parameters for querying listings.
type ListingQueryParams struct {
	PageSize  int     `form:"page_size" binding:"omitempty,min=1,max=100"`
	Cursor    string  `form:"cursor"`
	SortBy    string  `form:"sort_by" binding:"omitempty,oneof=created_at price sqft"`
	SortOrder string  `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	City      string  `form:"city"`
	MinPrice  float64 `form:"min_price" binding:"omitempty,min=0"`
}

// PaginationMeta holds metadata for paginated results.
type PaginationMeta struct {
	Total   int64
	HasNext bool
	Cursor  string
}

var allowedSortFields = map[string]bool{
	"price":      true,
	"city":       true,
	"created_at": true, // Add other valid fields here
}

var allowedSortOrders = map[string]bool{
	"asc":  true,
	"desc": true,
}

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
func CreateListingRepo(data Listing) error {
	var existingListing Listing
	if err := config.DB.Where("title = ? OR address = ?", data.Title, data.Address).First(&existingListing).Error; err == nil {
		return errors.New("listing with the given title or address already exists")

	}

	// Create a new listing
	listing := Listing{
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

// GetListingsRepo retrieves listings from the database using various query parameters.
// It supports filtering by city and minimum price, sorting by a specified field and order,
// and implements cursor-based pagination. Additionally, the associated Realtor data for each listing is preloaded.
//
// Parameters:
//
//	params - a ListingQueryParams struct that includes:
//	  - City: filter for listings in a specific city.
//	  - MinPrice: minimum price filter for listings.
//	  - SortBy & SortOrder: field and order to sort the result.
//	  - Cursor: timestamp used for pagination to retrieve listings created before this value.
//	  - PageSize: the maximum number of listings to retrieve.
//
// Returns:
//
//	[]models.Listing      - a slice containing the listings that match the filtering and sorting criteria.
//	PaginationMeta        - metadata about the pagination including total count, a boolean flag indicating if there's a next page,
//	                        and the cursor pointing to the last listing's creation time.
//	error                 - an error value that is non-nil if the operation encounters an issue.
func GetListingsRepo(params ListingQueryParams) ([]Listing, PaginationMeta, error) {
	query := config.DB.Model(&Listing{})

	// Validate and map SortBy field
	_, ok := allowedSortFields[params.SortBy]
	sortBy := params.SortBy
	if !ok {
		sortBy = "created_at" // Default to created_at field
	}

	// Validate and map SortOrder value
	sortOrder := params.SortOrder
	if _, ok := allowedSortOrders[sortOrder]; !ok {
		sortOrder = "asc" // Default to ascending order
	}

	// Sorting
	query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

	// Cursor based pagination
	if params.Cursor != "" {
		query = query.Where("created_at < ?", params.Cursor)
	}

	// Eager loading
	query = query.Preload("Realtor")

	// Execute query
	var listings []Listing
	result := query.Limit(params.PageSize).Find(&listings)

	if result.Error != nil {
		return nil, PaginationMeta{}, result.Error
	}

	var cursor string
	hasNext := false

	if len(listings) > 0 {
		hasNext = len(listings) == params.PageSize
		cursor = listings[len(listings)-1].CreatedAt.Format(time.RFC3339)
	}

	var total int64
	config.DB.Model(&Listing{}).Count(&total)

	return listings, PaginationMeta{
		Total:   total,
		HasNext: hasNext,
		Cursor:  cursor,
	}, nil
}
