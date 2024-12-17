package api

import (
	"net/http"

	"backend.com/go-backend/src/repositories"
	"github.com/alexedwards/argon2id"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	type SignInInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
		return
	}

	// Check if user exists
	user, err := repositories.GetUserRepository(input.Email)
	println("User:", user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user", "message": err.Error()})
		return
	}

	// check if password is correct
	if match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password); err != nil || !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "message": "Incorrect password"})
		return
	}

	session := sessions.Default(c)
	session.Set("email", user.Email)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": "User signed in!"})
}

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": "Welcome to the dashboard!", "user": "you"})
}
