package routers

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/internal/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("userEmail")

		if after, ok := strings.CutPrefix(c.Request.URL.Path, "/auth/"); ok {
			provider := after
			provider = strings.Split(provider, "/")[0]

			q := c.Request.URL.Query()
			q.Add("provider", provider)
			c.Request.URL.RawQuery = q.Encode()

			c.Next()
			return
		}

		if email == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "session expired or not logged in",
			})
			return
		}

		c.Set("userEmail", email)
		c.Next()
	}
}

func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, ok := c.Get("db")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Database client not found",
				"message": "Unable to process request due to missing database connection",
			})
			return
		}
		dbTyped, ok := db.(*config.Database)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Invalid database client type",
				"message": "Database client is not of the expected type",
			})
			return
		}
		// Set the *ent.Client directly in the context
		c.Set("entClient", dbTyped.Client)
		c.Next()
	}
}
