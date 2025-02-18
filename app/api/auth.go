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
	"golang.org/x/net/context"
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
	session.Set("userEmail", user.Email)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/api/v1/users/me")
}

func AuthInit(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	session := sessions.Default(c)
	session.Set("oauth_redirect", "/api/v1/users/me")
	session.Save()

	authUrl, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get auth URL", "message": err.Error()})
		return
	}
	fmt.Printf("URL: %s\n", authUrl)
	c.Redirect(http.StatusTemporaryRedirect, authUrl)
}

// check if user email is verified
// rawData := user.RawData
// emailVerified, _ := rawData["email_verified"].(bool)
// if !emailVerified {
// 	c.JSON(http.StatusForbidden, gin.H{"error": "OAuth failed", "message": "Email not verified"})
// 	return
// }

func AuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	req := c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(c.Writer, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth failed", "message": fmt.Sprintf("%s authentication failed: %v", provider, err)})
		return
	}

	if verified, _ := user.RawData["email_verified"].(bool); !verified && provider == "google" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Email verification required", "message": "Please verify your email with " + provider})
		return
	}

	var existingUser models.User
	result := config.DB.Where("provider = ? AND provider_id = ?", "google", user.UserID).First(&existingUser)
	if result.Error != nil {
		newUser := models.User{
			Avatar:     user.AvatarURL,
			Email:      user.Email,
			FullName:   user.Name,
			Username:   user.Email,
			Provider:   provider,
			ProviderID: user.UserID,
			Password:   fmt.Sprintf("oauth-%s-user", provider),
		}
		if err := config.DB.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "message": fmt.Sprintf("Could not create %s user: %v", provider, err)})
			return
		}
		existingUser = newUser
	}

	session := sessions.Default(c)
	session.Set("userEmail", existingUser.Email)
	session.Set("authProvider", provider)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Session error", "message": "Failed to save authentication session"})
		return
	}

	if redirectURL := session.Get("oauth_redirect"); redirectURL != nil {
		session.Delete("oauth_redirect")
		session.Save()
		c.Redirect(http.StatusSeeOther, redirectURL.(string))
	} else {
		c.Redirect(http.StatusFound, "/api/v1/users/me")
	}
}
