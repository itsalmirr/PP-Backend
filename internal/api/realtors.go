package api

import (
	"log/slog"
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

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	err := repositories.CreateRealtorRepo(c.Request.Context(), entClient, input)
	if err != nil {
		slog.Error("failed to create realtor", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create realtor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "data": "Realtor created!"})
}

func GetRealtor(c *gin.Context) {
	email := c.Param("email")

	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	realtor, err := repositories.GetRealtorRepo(c.Request.Context(), entClient, email)
	if err != nil {
		slog.Error("failed to get realtor", "email", email, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get realtor"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": realtor})
}

func GetRealtors(c *gin.Context) {
	entClient, ok := getEntClient(c)
	if !ok {
		return
	}

	realtors, err := repositories.GetRealtorsRepo(c.Request.Context(), entClient)
	if err != nil {
		slog.Error("failed to get realtors", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get realtors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   realtors,
	})
}
