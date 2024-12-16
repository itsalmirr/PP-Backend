package api

import (
	"net/http"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
	"backend.com/go-backend/src/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": "User created!"})
}

// GetUser godoc
// @Summary Get a user by username
// @Description Get a user by username
func GetUser(c *gin.Context) {
	username := c.Param("username")
	var user models.User

	if err := config.DB.Select("id", "avatar", "email", "username", "full_name", "start_date", "is_staff", "is_active").Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found!",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
