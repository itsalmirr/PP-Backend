package api

import (
	"net/http"

	"backend.com/go-backend/app/repositories"
	"github.com/gin-gonic/gin"
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
	var input repositories.CreateRealtorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}

	err := repositories.CreateRealtorRepository(input)
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

	realtor, err := repositories.GetRealtorRepository(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get realtor", "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": realtor})
}
