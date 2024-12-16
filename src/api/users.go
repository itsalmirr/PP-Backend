package api

import (
	"net/http"

	"backend.com/go-backend/src/repositories"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param input body CreateUserInput true "User input"
// @Endpoint /users [post]
func CreateUser(c *gin.Context) {
	// Create a new user
	var input repositories.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}

	err := repositories.CreateUserRepository(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": "User created!"})
}

// GetUser godoc
// @Summary Get a user by username
// @Description Get a user by username
func GetUser(c *gin.Context) {
	email := c.Param("email")

	user, err := repositories.GetUserRepository(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user", "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
