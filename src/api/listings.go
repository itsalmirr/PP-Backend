package api

import (
	"net/http"

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

	err := repositories.CreateListingRepository(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create listing", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "data": "Listing created!"})
}
