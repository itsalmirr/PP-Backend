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
