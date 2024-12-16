package api

import (
	"net/http"

	"backend.com/go-backend/src/repositories"
	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user.
// @Summary Create a new user
// @Description This endpoint creates a new user with the provided input data.
// @Tags users
// @Accept json
// @Produce json
// @Param user body repositories.CreateUserInput true "User input data"
// @Success 200 {object} map[string]interface{} "status: OK, data: User created!"
// @Failure 400 {object} map[string]interface{} "error: Invalid input, message: Please provide required fields"
// @Failure 500 {object} map[string]interface{} "error: Failed to create user, message: error message"
// @Router /users [post]
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

// GetUser handles the HTTP request to retrieve a user by email.
// @Summary Get user by email
// @Description Get user by email from the repository
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} gin.H{"data": interface{}}
// @Failure 500 {object} gin.H{"error": string, "message": string}
// @Router /users/{email} [get]
func GetUser(c *gin.Context) {
	email := c.Param("email")

	user, err := repositories.GetUserRepository(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user", "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
