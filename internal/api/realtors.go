package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ppgroup.m0chi.com/ent"
	"ppgroup.m0chi.com/internal/repositories"
)

// CreateRealtor handles the creation of a new realtor.
// @Summary Create a new realtor
// @Description Create a new realtor with the provided input data
// @Tags realtors
// @Accept json
// @Produce json
// @Param input body repositories.CreateRealtorInput true "Realtor input data"
// @Success 201 {object} gin.H{"status": "OK", "data": "Realtor created!"}
// @Failure 400 {object} gin.H{"error": "Invalid input", "message": "Please provide required fields"}
// @Failure 500 {object} gin.H{"error": "Failed to create realtor", "message": "Error message"}
// @Router /realtors [post]
func CreateRealtor(c *gin.Context) {
	var input *ent.Realtor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"message": "Please provide required fields",
		})
		return
	}
	entClient := c.MustGet("entClient").(*ent.Client)

	err := repositories.CreateRealtorRepo(entClient, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create realtor", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "data": "Realtor created!"})
}

// GetRealtor handles the request to retrieve a realtor by email.
// @Summary Get a realtor by email
// @Description Get details of a realtor using their email address
// @Tags realtors
// @Accept json
// @Produce json
// @Param email path string true "Realtor Email"
// @Success 200 {object} gin.H{"data": models.Realtor}
// @Failure 500 {object} gin.H{"error": string, "message": string}
// @Router /realtors/{email} [get]
func GetRealtor(c *gin.Context) {
	email := c.Param("email")
	entClient := c.MustGet("entClient").(*ent.Client)

	realtor, err := repositories.GetRealtorRepo(entClient, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get realtor",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": realtor})
}

// GetRealtors handles HTTP requests to retrieve all realtors.
// It extracts the ent.Client from the Gin context, queries the repository
// for all realtors, and returns the result as JSON.
// On success, it returns HTTP 200 with the realtors data.
// On failure, it returns HTTP 500 with an error message.
func GetRealtors(c *gin.Context) {
	entClient := c.MustGet("entClient").(*ent.Client)

	realtors, err := repositories.GetRealtorsRepo(entClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get realtors",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   realtors,
	})
}
