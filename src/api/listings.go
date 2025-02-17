package api

import (
	"net/http"

	"backend.com/go-backend/src/models"
	"backend.com/go-backend/src/repositories"
	"github.com/gin-gonic/gin"
)

type ListingQueryParams = repositories.ListingQueryParams

// CreateListing handles the creation of a new listing.
// @Summary Create a new listing
// @Description Create a new listing with the provided input data
// @Tags listings
// @Accept json
// @Produce json
// @Param input body repositories.CreateListingInput true "Listing input data"
// @Success 201 {object} gin.H{"status": "OK", "data": "Listing created!"}
// @Failure 400 {object} gin.H{"error": "Invalid input", "message": "Please provide required fields"}
// @Failure 500 {object} gin.H{"error": "Failed to create listing", "message": "Error message"}
// @Router /listings [post]
func CreateListing(c *gin.Context) {
	var input models.Listing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}
	// Create listing
	err := repositories.CreateListingRepo(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing, please check your input", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "data": "Listing created!"})
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

	// Default values if not provided
	if params.PageSize <= 0 || params.PageSize > 100 {
		params.PageSize = 10 // Default page size
	}

	// Default sort field and order
	if params.SortBy == "" {
		params.SortBy = "created_at" // Default sort field
	}

	// Default sort order
	if params.SortOrder == "" {
		params.SortOrder = "desc" // Default sort order
	}

	// Get listings from repo
	listings, meta, err := repositories.GetListingsRepo(params)
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
