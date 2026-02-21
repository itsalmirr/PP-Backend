package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/ent/listing"
	"ppgroup.ppgroup.com/ent/schema"
	"ppgroup.ppgroup.com/internal/repositories"
	"ppgroup.ppgroup.com/internal/services"
	"ppgroup.ppgroup.com/internal/util"
)

type ListingQueryParams = repositories.ListingQueryParams

// getEntClient safely retrieves the ent.Client from the Gin context.
func getEntClient(c *gin.Context) (*ent.Client, bool) {
	val, exists := c.Get("entClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return nil, false
	}
	client, ok := val.(*ent.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return nil, false
	}
	return client, true
}

// getImageService safely retrieves the ImageService from the Gin context.
func getImageService(c *gin.Context) (services.ImageUploader, bool) {
	val, exists := c.Get("imageService")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return nil, false
	}
	svc, ok := val.(services.ImageUploader)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return nil, false
	}
	return svc, true
}

// parseListingForm extracts and validates all listing fields from a multipart form.
func parseListingForm(c *gin.Context) (*ent.Listing, error) {
	title := c.PostForm("title")
	address := c.PostForm("address")
	city := c.PostForm("city")
	state := c.PostForm("state")
	zipCode := c.PostForm("zip_code")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	bedroomStr := c.PostForm("bedroom")
	bathroomStr := c.PostForm("bathroom")
	garageStr := c.PostForm("garage")
	sqftStr := c.PostForm("sqft")
	typeOfPropertyStr := c.PostForm("type_of_property")
	lotSizeStr := c.PostForm("lot_size")
	poolStr := c.PostForm("pool")
	yearBuiltStr := c.PostForm("year_built")
	realtorIDStr := c.PostForm("realtor_id")

	if title == "" || address == "" || city == "" || state == "" || zipCode == "" ||
		priceStr == "" || bedroomStr == "" || bathroomStr == "" || sqftStr == "" ||
		typeOfPropertyStr == "" || yearBuiltStr == "" || realtorIDStr == "" {
		return nil, fmt.Errorf("missing required fields")
	}

	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		return nil, fmt.Errorf("invalid price format")
	}

	bedroom, err := strconv.Atoi(bedroomStr)
	if err != nil {
		return nil, fmt.Errorf("invalid bedroom format")
	}

	bathroom, err := strconv.ParseFloat(bathroomStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid bathroom format")
	}

	sqft, err := strconv.Atoi(sqftStr)
	if err != nil {
		return nil, fmt.Errorf("invalid sqft format")
	}

	yearBuilt, err := strconv.Atoi(yearBuiltStr)
	if err != nil {
		return nil, fmt.Errorf("invalid year_built format")
	}

	realtorID, err := uuid.Parse(realtorIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid realtor_id format")
	}

	var garage int
	if garageStr != "" {
		garage, err = strconv.Atoi(garageStr)
		if err != nil {
			return nil, fmt.Errorf("invalid garage format")
		}
	}

	var lotSize int
	if lotSizeStr != "" {
		lotSize, err = strconv.Atoi(lotSizeStr)
		if err != nil {
			return nil, fmt.Errorf("invalid lot_size format")
		}
	}

	var pool bool
	if poolStr != "" {
		pool, err = strconv.ParseBool(poolStr)
		if err != nil {
			return nil, fmt.Errorf("invalid pool format")
		}
	}

	var typeOfProperty listing.TypeOfProperty
	switch strings.ToLower(typeOfPropertyStr) {
	case "house":
		typeOfProperty = listing.TypeOfPropertyHouse
	case "apartment":
		typeOfProperty = listing.TypeOfPropertyApartment
	case "condo":
		typeOfProperty = listing.TypeOfPropertyCondo
	case "townhouse":
		typeOfProperty = listing.TypeOfPropertyTownhouse
	default:
		return nil, fmt.Errorf("invalid type_of_property, must be one of: house, apartment, condo, townhouse")
	}

	return &ent.Listing{
		Title:          title,
		Address:        address,
		City:           city,
		State:          state,
		ZipCode:        zipCode,
		Description:    description,
		Price:          price,
		Bedroom:        bedroom,
		Bathroom:       bathroom,
		Garage:         garage,
		Sqft:           sqft,
		TypeOfProperty: typeOfProperty,
		LotSize:        lotSize,
		Pool:           pool,
		YearBuilt:      yearBuilt,
		RealtorID:      realtorID,
		Status:         listing.StatusDRAFT,
	}, nil
}

func CreateListing(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	listingData, err := parseListingForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle image uploads
	var mediaItems []schema.Media
	imageService, ok := getImageService(c)
	if !ok {
		return
	}

	files := c.Request.MultipartForm.File["images"]
	if len(files) == 0 {
		files = c.Request.MultipartForm.File["image"]
	}

	if len(files) > 0 {
		for i, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
				return
			}
			defer file.Close()

			// Validate by MIME type (magic bytes), not extension
			valid, err := util.ValidateImageFile(file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate file"})
				return
			}
			if !valid {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("Invalid file type: %s. Accepted: JPEG, PNG, GIF, WebP", fileHeader.Filename),
				})
				return
			}

			url, err := imageService.UploadImage(c.Request.Context(), file, fileHeader.Filename)
			if err != nil {
				slog.Error("image upload failed", "filename", fileHeader.Filename, "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
				return
			}

			if url == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Image upload returned empty URL"})
				return
			}

			mediaItems = append(mediaItems, schema.Media{
				URL:       url,
				Type:      "image",
				IsPrimary: i == 0,
			})
		}
	}

	listingData.Media = mediaItems

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	err = repositories.CreateListingRepo(c.Request.Context(), entClient, listingData)
	if err != nil {
		slog.Error("failed to create listing", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Listing created successfully!",
		"data": gin.H{
			"title":           listingData.Title,
			"address":         listingData.Address,
			"uploaded_images": len(mediaItems),
		},
	})
}

func CreateListingJSON(c *gin.Context) {
	var input *ent.Listing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}

	if input.Title == "" || input.Address == "" || input.City == "" || input.State == "" ||
		input.ZipCode == "" || input.Price.IsZero() || input.Bedroom == 0 ||
		input.Bathroom == 0 || input.Sqft == 0 || input.YearBuilt == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	if input.Status == "" {
		input.Status = listing.StatusDRAFT
	}

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	err := repositories.CreateListingRepo(c.Request.Context(), entClient, input)
	if err != nil {
		slog.Error("failed to create listing", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Listing created successfully!",
		"data": gin.H{
			"title":       input.Title,
			"address":     input.Address,
			"media_count": len(input.Media),
		},
	})
}

func GetListings(c *gin.Context) {
	var params ListingQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	listings, meta, err := repositories.GetListingsRepo(c.Request.Context(), entClient, params)
	if err != nil {
		slog.Error("failed to retrieve listings", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve listings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   listings,
		"pagination": gin.H{
			"total":       meta.Total,
			"has_next":    meta.HasNext,
			"next_cursor": meta.Cursor,
			"page_size":   params.PageSize,
		},
	})
}

func DeleteListing(c *gin.Context) {
	ID := c.Query("ID")
	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID query parameter"})
		return
	}

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	err := repositories.DeleteListing(c.Request.Context(), entClient, ID)
	if err != nil {
		slog.Error("failed to delete listing", "id", ID, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete listing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Successfully deleted listing",
	})
}

func UpdateListing(c *gin.Context) {
	var input *ent.Listing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	err := repositories.UpdateListingRepo(c.Request.Context(), entClient, input)
	if err != nil {
		slog.Error("failed to update listing", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update listing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Listing updated!"})
}
