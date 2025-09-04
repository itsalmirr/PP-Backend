package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"ppgroup.m0chi.com/ent"
	"ppgroup.m0chi.com/ent/listing"
)

// ListingQueryParams holds parameters for querying listings.
type ListingQueryParams struct {
	PageSize  int             `form:"page_size" binding:"omitempty,min=1,max=100"`
	Cursor    string          `form:"cursor"`
	SortBy    string          `form:"sort_by" binding:"omitempty,oneof=created_at price sqft"`
	SortOrder string          `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	City      string          `form:"city"`
	MinPrice  decimal.Decimal `form:"min_price" binding:"omitempty,min=0"`
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

func CreateListingRepo(entClient *ent.Client, data *ent.Listing) error {
	ctx := context.Background()

	exists, err := entClient.Listing.Query().Where(listing.Or(listing.TitleEQ(data.Title), listing.AddressEQ(data.Address))).Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("listing with the given title or address already exists")
	}

	// Start a transaction
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	// Create a new listing
	_, err = tx.Listing.Create().
		SetAddress(data.Address).
		SetTitle(data.Title).
		SetCity(data.City).
		SetState(data.State).
		SetZipCode(data.ZipCode).
		SetDescription(data.Description).
		SetPrice(data.Price).
		SetBedroom(data.Bedroom).
		SetBathroom(data.Bathroom).
		SetGarage(data.Garage).
		SetSqft(data.Sqft).
		SetTypeOfProperty(data.TypeOfProperty).
		SetLotSize(data.LotSize).
		SetPool(data.Pool).
		SetYearBuilt(data.YearBuilt).
		SetMedia(data.Media).
		SetStatus(data.Status).
		SetRealtorID(data.RealtorID).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return errors.New("failed to create listing" + err.Error())
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}

// GetListingsRepo retrieves a paginated list of listings with optional filtering and sorting.
//
// It fetches listings from the database according to the provided query parameters, supporting
// filtering by city and minimum price. The function implements cursor-based pagination and
// can sort results by price, city, or creation time.
//
// Parameters:
//   - entClient: Ent client for database operations
//   - params: ListingQueryParams containing filtering, sorting, and pagination options
//
// Returns:
//   - []*ent.Listing: Array of listing entities matching the query parameters
//   - PaginationMeta: Metadata about the result set (total count, cursor for next page, etc.)
//   - error: Any error that occurred during the query execution
//
// The function handles the following pagination logic:
//   - Default page size is 10 if not specified
//   - Provides a cursor for fetching the next page of results
//   - Returns total count of matching records regardless of pagination
func GetListingsRepo(entClient *ent.Client, params ListingQueryParams) ([]*ent.Listing, PaginationMeta, error) {
	ctx := context.Background()

	query := entClient.Listing.Query()
	query = query.WithRealtor()

	if params.City != "" {
		query = query.Where(listing.CityEQ(params.City))
	}
	if params.MinPrice.GreaterThan(decimal.NewFromInt(0)) {
		query = query.Where(listing.PriceGTE(params.MinPrice))
	}

	// Get total count
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, PaginationMeta{}, err
	}

	// Sorting
	if params.SortBy != "" && allowedSortFields[params.SortBy] {
		order := params.SortOrder
		if !allowedSortOrders[order] {
			order = "asc"
		}
		switch params.SortBy {
		case "price":
			if order == "asc" {
				query = query.Order(ent.Asc(listing.FieldPrice))
			} else {
				query = query.Order(ent.Desc(listing.FieldPrice))
			}
		case "city":
			if order == "asc" {
				query = query.Order(ent.Asc(listing.FieldCity))
			} else {
				query = query.Order(ent.Desc(listing.FieldCity))
			}
		case "created_at":
			if order == "asc" {
				query = query.Order(ent.Asc(listing.FieldCreateTime))
			} else {
				query = query.Order(ent.Desc(listing.FieldCreateTime))
			}
		}
	} else {
		query = query.Order(ent.Asc(listing.FieldCreateTime))
	}

	// Cursor based pagination
	if params.Cursor != "" {
		cursorID, err := uuid.Parse(params.Cursor)
		if err != nil {
			return nil, PaginationMeta{}, errors.New("invalid cursor")
		}

		if params.SortOrder == "asc" {
			query = query.Where(listing.IDLT(cursorID))
		} else {
			query = query.Where(listing.IDGT(cursorID))
		}
	}

	pageSize := params.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	query = query.Limit(pageSize + 1)

	listings, err := query.All(ctx)
	if err != nil {
		return nil, PaginationMeta{}, err
	}

	hasNext := len(listings) > pageSize
	if hasNext {
		listings = listings[:pageSize]
	}

	var nextCursor string
	if hasNext {
		lastListing := listings[len(listings)-1]
		nextCursor = lastListing.ID.String()
	}

	meta := PaginationMeta{
		Total:   int64(total),
		HasNext: hasNext,
		Cursor:  nextCursor,
	}

	return listings, meta, nil
}

// DeleteListing deletes a listing from the database using the provided ent.Client and Listing data.
// It takes an ent.Client instance and a Listing entity as input parameters.
// If the listing with the specified ID does not exist, it returns an error indicating "listing not found".
// If any other error occurs during the deletion process, it wraps and returns the error.
// On successful deletion, it returns nil.
//
// Parameters:
// - entClient: The ent.Client instance used to interact with the database.
// - data: The Listing entity containing the ID of the listing to be deleted.
//
// Returns:
// - error: An error if the deletion fails or the listing is not found, otherwise nil.
func DeleteListing(entClient *ent.Client, idStr string) error {
	ctx := context.Background()
	ID, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid ID format")
	}

	err = entClient.Listing.DeleteOneID(ID).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("listing not found")
		}
		return fmt.Errorf("failed to delete listing: %w", err)
	}

	return nil
}

func UpdateListingRepo(entClient *ent.Client, data *ent.Listing) error {
	ctx := context.Background()

	// Fetch the current listing from the database
	current, err := entClient.Listing.Get(ctx, data.ID)
	if err != nil {
		return err
	}

	// Check for duplicate title or address if they are changed
	if data.Title != current.Title || data.Address != current.Address {
		duplicate, err := entClient.Listing.Query().
			Where(
				listing.Or(
					listing.TitleEQ(data.Title),
					listing.AddressEQ(data.Address),
				),
				listing.IDNEQ(data.ID),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if duplicate {
			return errors.New("another listing with the given title or address already exists")
		}
	}

	// Begin building the update, only setting fields that have changed
	updater := entClient.Listing.UpdateOneID(data.ID)

	if data.Title != current.Title {
		updater = updater.SetTitle(data.Title)
	}
	if data.Address != current.Address {
		updater = updater.SetAddress(data.Address)
	}
	if data.City != current.City {
		updater = updater.SetCity(data.City)
	}
	if data.State != current.State {
		updater = updater.SetState(data.State)
	}
	if data.ZipCode != current.ZipCode {
		updater = updater.SetZipCode(data.ZipCode)
	}
	if data.Description != current.Description {
		updater = updater.SetDescription(data.Description)
	}
	if !data.Price.Equal(current.Price) {
		updater = updater.SetPrice(data.Price)
	}
	if data.Bedroom != current.Bedroom {
		updater = updater.SetBedroom(data.Bedroom)
	}
	if data.Bathroom != current.Bathroom {
		updater = updater.SetBathroom(data.Bathroom)
	}
	if data.Garage != current.Garage {
		updater = updater.SetGarage(data.Garage)
	}
	if data.Sqft != current.Sqft {
		updater = updater.SetSqft(data.Sqft)
	}
	if data.TypeOfProperty != current.TypeOfProperty {
		updater = updater.SetTypeOfProperty(data.TypeOfProperty)
	}
	if data.LotSize != current.LotSize {
		updater = updater.SetLotSize(data.LotSize)
	}
	if data.Pool != current.Pool {
		updater = updater.SetPool(data.Pool)
	}
	if data.YearBuilt != current.YearBuilt {
		updater = updater.SetYearBuilt(data.YearBuilt)
	}
	if data.Status != current.Status {
		updater = updater.SetStatus(data.Status)
	}
	if data.RealtorID != current.RealtorID {
		updater = updater.SetRealtorID(data.RealtorID)
	}
	if data.Media != nil {
		updater = updater.SetMedia(data.Media)
	}

	_, err = updater.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
