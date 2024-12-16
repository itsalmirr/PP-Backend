package api

import (
	"net/http"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsStaff  bool   `json:"is_staff"`
}

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
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ? OR username = ?", input.Email, input.Username).
		First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := argon2id.CreateHash(input.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		return
	}

	user := models.User{
		Avatar:   input.Avatar,
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
		Password: hashedPassword,
		IsStaff:  input.IsStaff,
	}
	tx := config.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user!"})
		return
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"data": user})
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
