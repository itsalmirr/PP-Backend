package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"golang.org/x/net/context"
	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/ent/user"
	"ppgroup.ppgroup.com/internal/config"
)

// getFrontendURL reads the frontend URL from context (set by middleware).
func getFrontendURL(c *gin.Context) string {
	if val, exists := c.Get("frontendURL"); exists {
		if url, ok := val.(string); ok {
			return url
		}
	}
	return "http://localhost:3000"
}

// generateOAuthPassword creates a random password for OAuth users
// instead of a predictable placeholder.
func generateOAuthPassword() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		// Fallback — still better than a static string
		return fmt.Sprintf("oauth-%d", b[0])
	}
	return hex.EncodeToString(b)
}

func AuthInit(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	frontendURL := getFrontendURL(c)

	session := sessions.Default(c)
	session.Set("oauth_redirect", frontendURL)
	if err := session.Save(); err != nil {
		slog.Error("failed to save oauth redirect session", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize authentication"})
		return
	}

	authURL, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		slog.Error("failed to get auth URL", "provider", provider, "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get auth URL"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func AuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	req := c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))

	oauthUser, err := gothic.CompleteUserAuth(c.Writer, req)
	if err != nil {
		slog.Error("oauth authentication failed", "provider", provider, "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OAuth failed", "message": fmt.Sprintf("%s authentication failed", provider)})
		return
	}

	if verified, _ := oauthUser.RawData["verified_email"].(bool); !verified && provider == "google" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Email verification required", "message": "Please verify your email with " + provider})
		return
	}

	val, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	db := val.(*config.Database)

	existingUser, err := db.Client.User.Query().
		Where(user.ProviderEQ(oauthUser.Provider), user.ProviderIDEQ(oauthUser.UserID)).
		First(c.Request.Context())

	if err != nil {
		if ent.IsNotFound(err) {
			newUser, createErr := db.Client.User.Create().
				SetAvatar(oauthUser.AvatarURL).
				SetEmail(oauthUser.Email).
				SetFullName(oauthUser.Name).
				SetUsername(oauthUser.Email).
				SetProvider(provider).
				SetProviderID(oauthUser.UserID).
				SetPassword(generateOAuthPassword()).
				Save(c.Request.Context())
			if createErr != nil {
				slog.Error("failed to create oauth user", "provider", provider, "error", createErr)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}
			existingUser = newUser
		} else {
			slog.Error("failed to query user", "provider", provider, "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to look up user"})
			return
		}
	}

	session := sessions.Default(c)
	session.Set("userEmail", existingUser.Email)
	session.Set("authProvider", provider)
	if err := session.Save(); err != nil {
		slog.Error("failed to save session after oauth", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save authentication session"})
		return
	}

	frontendURL := getFrontendURL(c)

	if redirectURL := session.Get("oauth_redirect"); redirectURL != nil {
		session.Delete("oauth_redirect")
		if err := session.Save(); err != nil {
			slog.Error("failed to clear oauth redirect", "error", err)
		}
		c.Redirect(http.StatusSeeOther, redirectURL.(string))
	} else {
		c.Redirect(http.StatusFound, frontendURL)
	}
}
