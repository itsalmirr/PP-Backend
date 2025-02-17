package routers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")

		// Enhanced session validation
		if email == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Session expired or invalid",
			})
			return
		}

		// Add user context for downstream handlers
		c.Set("userEmail", email)
		c.Next()
	}
}
