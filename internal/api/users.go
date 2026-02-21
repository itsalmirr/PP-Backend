package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/internal/repositories"
)

func CreateUser(c *gin.Context) {
	var input repositories.CreateUserInput
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

	err := repositories.CreateUserRepo(c.Request.Context(), entClient, &input)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   "User created!",
	})
}

func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	email, ok := session.Get("userEmail").(string)
	if !ok || email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Please sign in"})
		return
	}

	entClient, clientOk := getEntClient(c)
	if !clientOk {
		return
	}

	user, err := repositories.GetUserRepo(c.Request.Context(), entClient, email)
	if err != nil {
		slog.Error("failed to get user for dashboard", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"data": user})
}
