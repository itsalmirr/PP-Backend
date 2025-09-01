package auth

import (
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.m0chi.com/ent"
	"ppgroup.m0chi.com/internal/repositories"
)

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func EmailSignIn(c *gin.Context) {
	var input SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"message": "Please provide required fields",
		})
		return
	}

	log.Printf("Attempting sign-in for email: %s", input.Email)

	// Check if user exists
	entClient := c.MustGet("entClient").(*ent.Client)
	user, err := repositories.GetUserRepo(entClient, input.Email)
	if err != nil {
		log.Printf("Error finding user '%s': %v", input.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}
	log.Printf("User found: %s, Password hash length: %d", user.Email, len(user.Password))

	// Check if password is correct
	log.Printf("Comparing password for user '%s'", user.Email)
	match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password)
	if err != nil {
		log.Printf("argon2id comparison error for user '%s': %v", user.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	log.Printf("Password match result for user '%s': %t", user.Email, match)

	if !match {
		log.Printf("Password mismatch for user '%s'.", user.Email)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("userEmail", user.Email)
	if err := session.Save(); err != nil {
		log.Printf("Session save error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}

	log.Printf("Login successful for user: %s", user.Email)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Login successful",
	})
}
