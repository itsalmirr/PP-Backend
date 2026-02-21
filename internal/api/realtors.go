package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/internal/repositories"
)

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

	err := repositories.CreateRealtorRepo(c.Request.Context(), entClient, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create realtor", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "data": "Realtor created!"})
}

func GetRealtor(c *gin.Context) {
	email := c.Param("email")
	entClient := c.MustGet("entClient").(*ent.Client)

	realtor, err := repositories.GetRealtorRepo(c.Request.Context(), entClient, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get realtor",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": realtor})
}

func GetRealtors(c *gin.Context) {
	entClient := c.MustGet("entClient").(*ent.Client)

	realtors, err := repositories.GetRealtorsRepo(c.Request.Context(), entClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get realtors",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   realtors,
	})
}
