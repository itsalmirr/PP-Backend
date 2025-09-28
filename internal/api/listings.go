package api

import (
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
)

type ListingQueryParams = repositories.ListingQueryParams

// CreateListing handles the creation of a new listing with image upload.
// @Summary Create a new listing with images
// @Description Create a new listing with multipart form data including images
// @Tags listings
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Listing title (10-120 chars)"
// @Param address formData string true "Property address (unique)"
// @Param city formData string true "City"
// @Param state formData string true "State (2 letters, e.g., CA)"
// @Param zip_code formData string true "ZIP code (5 digits)"
// @Param description formData string false "Property description"
// @Param price formData number true "Property price"
// @Param bedroom formData int true "Number of bedrooms"
// @Param bathroom formData number true "Number of bathrooms"
// @Param garage formData int false "Number of garage spaces"
// @Param sqft formData int true "Square footage"
// @Param type_of_property formData string true "Type of property" Enums(house, apartment, condo, townhouse)
// @Param lot_size formData int false "Lot size"
// @Param pool formData bool false "Has pool"
// @Param year_built formData int true "Year built"
// @Param realtor_id formData string true "Realtor UUID"
// @Param images formData file false "Property images (multiple files allowed, formats: jpg, jpeg, png, gif, webp)"
// @Success 201 {object} gin.H{"status": "OK", "message": "Listing created!", "data": object}
// @Failure 400 {object} gin.H{"error": "Invalid input", "message": string}
// @Failure 500 {object} gin.H{"error": "Failed to create listing", "message": string}
// @Router /api/v1/properties/add [post]
func CreateListing(c *gin.Context) {
	// Parse multipart form
	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB max
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form", "message": err.Error()})
		return
	}

	// Get form values
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

	// Validate required fields
	if title == "" || address == "" || city == "" || state == "" || zipCode == "" ||
		priceStr == "" || bedroomStr == "" || bathroomStr == "" || sqftStr == "" ||
		typeOfPropertyStr == "" || yearBuiltStr == "" || realtorIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Parse numeric fields
	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
		return
	}

	bedroom, err := strconv.Atoi(bedroomStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bedroom format"})
		return
	}

	bathroom, err := strconv.ParseFloat(bathroomStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bathroom format"})
		return
	}

	sqft, err := strconv.Atoi(sqftStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sqft format"})
		return
	}

	yearBuilt, err := strconv.Atoi(yearBuiltStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year_built format"})
		return
	}

	realtorID, err := uuid.Parse(realtorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid realtor_id format"})
		return
	}

	// Parse optional fields
	var garage int
	if garageStr != "" {
		garage, err = strconv.Atoi(garageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid garage format"})
			return
		}
	}

	var lotSize int
	if lotSizeStr != "" {
		lotSize, err = strconv.Atoi(lotSizeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lot_size format"})
			return
		}
	}

	var pool bool
	if poolStr != "" {
		pool, err = strconv.ParseBool(poolStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pool format"})
			return
		}
	}

	// Validate and convert type_of_property
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type_of_property. Must be one of: house, apartment, condo, townhouse"})
		return
	}

	// Handle image uploads
	var mediaItems []schema.Media
	imageService := c.MustGet("imageService").(*services.ImageService)

	// Get uploaded files
	files := c.Request.MultipartForm.File["images"]

	// If no files with "images" key, try singular "image"
	if len(files) == 0 {
		files = c.Request.MultipartForm.File["image"]
	}

	// Only proceed with image processing if we have actual files
	if len(files) > 0 {
		for i, fileHeader := range files {
			// Validate file type
			if !isValidImageType(fileHeader.Filename) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid file type: " + fileHeader.Filename,
				})
				return
			}

			// Open file
			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file: " + fileHeader.Filename})
				return
			}
			defer file.Close()

			// Upload to Cloudinary
			url, err := imageService.UploadImage(c.Request.Context(), file, fileHeader.Filename)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to Cloudinary: " + err.Error()})
				return
			}

			// Verify URL is not empty
			if url == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Image upload returned empty URL"})
				return
			}

			// Create media item
			mediaItem := schema.Media{
				URL:       url,
				Type:      "image",
				Caption:   "",     // You can add caption support later
				IsPrimary: i == 0, // First image is primary
			}
			mediaItems = append(mediaItems, mediaItem)
		}
	}

	// Create listing entity
	listing := &ent.Listing{
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
		Media:          mediaItems,
		RealtorID:      realtorID,
		Status:         listing.StatusDRAFT, // Default status
	}

	// Save to database
	entClient := c.MustGet("entClient").(*ent.Client)
	err = repositories.CreateListingRepo(entClient, listing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Listing created successfully!",
		"data": gin.H{
			"title":           listing.Title,
			"address":         listing.Address,
			"uploaded_images": len(mediaItems),
		},
	})
}

// CreateListingJSON handles the creation of a new listing with JSON input (for pre-existing image URLs).
// @Summary Create a new listing with JSON
// @Description Create a new listing using JSON format with existing image URLs
// @Tags listings
// @Accept json
// @Produce json
// @Param input body ent.Listing true "Listing input data"
// @Success 201 {object} gin.H{"status": "OK", "message": "Listing created!", "data": object}
// @Failure 400 {object} gin.H{"error": "Invalid input", "message": string}
// @Failure 500 {object} gin.H{"error": "Failed to create listing", "message": string}
// @Router /properties/add-json [post]
func CreateListingJSON(c *gin.Context) {
	var input *ent.Listing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields: " + err.Error()})
		return
	}

	// Validate required fields
	if input.Title == "" || input.Address == "" || input.City == "" || input.State == "" ||
		input.ZipCode == "" || input.Price.IsZero() || input.Bedroom == 0 ||
		input.Bathroom == 0 || input.Sqft == 0 || input.YearBuilt == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Set default status if not provided
	if input.Status == "" {
		input.Status = listing.StatusDRAFT
	}

	// Create listing
	entClient := c.MustGet("entClient").(*ent.Client)
	err := repositories.CreateListingRepo(entClient, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing", "message": err.Error()})
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

// isValidImageType checks if the file has a valid image extension
func isValidImageType(filename string) bool {
	validTypes := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	filename = strings.ToLower(filename)

	for _, ext := range validTypes {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

// GetListings handles the retrieval of paginated property listings.
// @Summary Get paginated listings
// @Description Retrieves a list of property listings with pagination support
// @Tags listings
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1, min: 1)"
// @Param limit query int false "Number of items per page (default: 10, min: 1)"
//
//	@Success 200 {object} gin.H{
//	    "status": string,
//	    "data": []repositories.Listing,
//	    "total": int64,
//	    "current_page": int,
//	    "total_page": int,
//	    "per_page": int
//	}
//
// @Failure 500 {object} gin.H{"error": string, "message": string}
// @Router /listings [get]
func GetListings(c *gin.Context) {
	var params ListingQueryParams

	// Bind query parameters to ListingQueryParams struct
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid query parameters",
			"details": err.Error(),
		})
		return
	}

	entClient := c.MustGet("entClient").(*ent.Client)
	// Get listings from repo
	listings, meta, err := repositories.GetListingsRepo(entClient, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrive listings",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"status": "OK",
		"data":   listings,
		"pagination": gin.H{
			"total":       meta.Total,
			"has_next":    meta.HasNext,
			"next_cursor": meta.Cursor,
			"page_size":   params.PageSize,
		},
	}
	c.JSON(http.StatusOK, response)
}

// DeleteListing handles the deletion of a listing based on the provided ID query parameter.
//
// @param c *gin.Context - The Gin context containing the HTTP request and response.
//
// The function performs the following steps:
//  1. Retrieves the "ID" query parameter from the request. If the parameter is missing, it responds with
//     a 400 Bad Request status and an error message.
//  2. Retrieves the ent.Client instance from the context.
//  3. Calls the repositories.DeleteListing function to delete the listing with the specified ID.
//     If an error occurs during deletion, it responds with a 500 Internal Server Error status and the error message.
//  4. If the deletion is successful, it responds with a 200 OK status and a success message.
func DeleteListing(c *gin.Context) {
	ID := c.Query("ID")
	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing ID query parameter",
		})
		return
	}

	entClient := c.MustGet("entClient").(*ent.Client)
	err := repositories.DeleteListing(entClient, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete listing",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Successfully deleted listing",
	})
}

// UpdateListing handles the updating of an existing listing.
// @Summary Update an existing listing
// @Description Update an existing listing with the provided input data
// @Tags listings
// @Accept json
// @Produce json
// @Param input body ent.Listing true "Listing update data"
// @Success 200 {object} gin.H{"status": "OK", "message": "Listing updated!"}
// @Failure 400 {object} gin.H{"error": "Invalid input", "message": "Please provide required fields"}
// @Failure 500 {object} gin.H{"error": "Failed to update listing", "message": "Error message"}
// @Router /listings [put]
func UpdateListing(c *gin.Context) {
	var input *ent.Listing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields: " + err.Error()})
		return
	}

	entClient := c.MustGet("entClient").(*ent.Client)

	err := repositories.UpdateListingRepo(entClient, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update listing", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Listing updated!"})
}
