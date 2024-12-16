package api

import (
	"net/http"

	"backend.com/go-backend/src/repositories"
	"github.com/gin-gonic/gin"
)

func CreateRealtor(c *gin.Context) {
	var input repositories.CreateRealtorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}

	err := repositories.CreateRealtorRepository(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "data": "Realtor created!"})
}
