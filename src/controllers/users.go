package controllers

import (
	"net/http"

	"backend.com/go-backend/src/cmd"
	"backend.com/go-backend/src/models"
	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsStaff  bool   `json:"is_staff"`
}

func CreateUser(c *gin.Context) {
	// Create a new user
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := argon2id.CreateHash(input.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
		return
	}

	input.Password = hashedPassword

	user := models.User{
		Avatar:   input.Avatar,
		Email:    input.Email,
		Username: input.Username,
		FullName: input.FullName,
		Password: input.Password,
		IsStaff:  input.IsStaff,
	}

	cmd.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
