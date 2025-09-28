package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"golang.org/x/net/context"
	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/ent/user"
	"ppgroup.ppgroup.com/internal/config"
)

func AuthInit(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	session := sessions.Default(c)
	session.Set("oauth_redirect", "http://localhost:3000")
	err := session.Save()
	if err != nil {
		return
	}

	authUrl, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get auth URL",
			"message": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, authUrl)
}

func AuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	req := c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))

	oauthUser, err := gothic.CompleteUserAuth(c.Writer, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth failed", "message": fmt.Sprintf("%s authentication failed: %v", provider, err)})
		return
	}

	if verified, _ := oauthUser.RawData["verified_email"].(bool); !verified && provider == "google" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Email verification required", "message": "Please verify your email with " + provider})
		return
	}

	db := c.MustGet("db").(*config.Database)
	existingUser, err := db.Client.User.Query().Where(user.ProviderEQ(oauthUser.Provider), user.ProviderIDEQ(oauthUser.UserID)).First(c.Request.Context())

	if err != nil {
		if ent.IsNotFound(err) {
			newUser, err := db.Client.User.Create().
				SetAvatar(oauthUser.AvatarURL).
				SetEmail(oauthUser.Email).
				SetFullName(oauthUser.Name).
				SetUsername(oauthUser.Email).
				SetProvider(provider).
				SetProviderID(oauthUser.UserID).
				SetPassword(fmt.Sprintf("oauth-%s-user", provider)).
				Save(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "message": fmt.Sprintf("Could not create %s user: %v", provider, err)})
				return
			}
			existingUser = newUser
		}
	}

	session := sessions.Default(c)
	session.Set("userEmail", existingUser.Email)
	session.Set("authProvider", provider)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Session error", "message": "Failed to save authentication session"})
		return
	}

	// Redirect to the original URL or the user profile page
	if redirectURL := session.Get("oauth_redirect"); redirectURL != nil {
		session.Delete("oauth_redirect")
		err := session.Save()
		if err != nil {
			return
		}
		c.Redirect(http.StatusSeeOther, redirectURL.(string))
	} else {
		c.Redirect(http.StatusFound, "http://localhost:3000")
	}
}
