package auth

import (
	"log/slog"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/internal/repositories"
)

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func EmailSignIn(c *gin.Context) {
	var input SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"message": "Please provide required fields",
		})
		return
	}

	val, exists := c.Get("entClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	entClient, ok := val.(*ent.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	user, err := repositories.GetUserRepo(c.Request.Context(), entClient, input.Email)
	if err != nil {
		slog.Warn("sign-in attempt failed: user lookup", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password)
	if err != nil {
		slog.Error("password comparison error", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	if !match {
		slog.Warn("sign-in attempt failed: password mismatch")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("userEmail", user.Email)
	if err := session.Save(); err != nil {
		slog.Error("failed to save session", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	slog.Info("login successful", "user_id", user.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Login successful",
	})
}
