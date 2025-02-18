package api

import (
	"fmt"
	"net/http"

	"backend.com/go-backend/app/config"
	"backend.com/go-backend/app/models"
	"backend.com/go-backend/app/repositories"
	"github.com/alexedwards/argon2id"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

// SignIn handles user sign-in requests.
// It expects a JSON payload with "email" and "password" fields.
// If the input is invalid, it returns a 400 status code with an error message.
// If the user does not exist or the password is incorrect, it returns a 401 status code with an error message.
// If there is an internal server error, it returns a 500 status code with an error message.
// On successful sign-in, it creates a session and returns a 200 status code with a success message.
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

func GoogleAuthInit(c *gin.Context) {
	provider := "google"
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	session := sessions.Default(c)
	session.Set("oauth_redirect", "/dashboard")
	session.Save()

	// gothic.BeginAuthHandler(c.Writer, c.Request)
	url, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Printf("URL: %s\n", url)
	c.Redirect(http.StatusFound, url)
}

func GoogleAuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request.WithContext(c))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth failed", "message": err.Error()})
		return
	}

	var existingUser models.User
	result := config.DB.Where("provider = ? AND provider_id = ?", "google", user.UserID).First(&existingUser)

	if result.Error != nil {
		newUser := models.User{
			Email:      user.Email,
			FullName:   user.Name,
			Provider:   "google",
			ProviderID: user.UserID,
			Password:   "oauth-user",
		}
		if err := config.DB.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "message": err.Error()})
			return
		}
		existingUser = newUser
	}

	session := sessions.Default(c)
	session.Set("email", existingUser.Email)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}

	redirectURL := session.Get("oauth_redirect")
	if redirectURL != nil {
		c.Redirect(http.StatusSeeOther, redirectURL.(string))
		return
	} else {
		c.Redirect(http.StatusFound, "/dashboard")
	}

}
