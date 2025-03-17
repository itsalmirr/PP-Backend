package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ppgroup.i0sys.com/ent"
	"ppgroup.i0sys.com/internal/repositories"
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
	var input *ent.Listing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields" + err.Error()})
		return
	}
	// Create listing
	entClient := c.MustGet("entClient").(*ent.Client)

	err := repositories.CreateListingRepo(entClient, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing, please check your input", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "message": "Listing created!"})
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
