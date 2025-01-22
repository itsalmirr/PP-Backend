package api

import (
	"math"
	"net/http"
	"strconv"

	"backend.com/go-backend/src/repositories"
	"github.com/gin-gonic/gin"
)

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
	var input repositories.CreateListingInput
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
	page := 1
	limit := 10

	if pageQuery := c.Query("page"); pageQuery != "" {
		if pageNum, err := strconv.Atoi(pageQuery); err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	if limitQuery := c.Query("limit"); limitQuery != "" {
		if limitNum, err := strconv.Atoi(limitQuery); err == nil && limitNum > 0 {
			limit = limitNum
		}
	}

	listings, total, err := repositories.GetListingsRepo(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get listings", "message": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"data":         listings,
		"total":        total,
		"current_page": page,
		"total_page":   totalPages,
		"per_page":     limit,
	})

}
