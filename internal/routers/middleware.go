package routers

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("userEmail")

		if strings.HasPrefix(c.Request.URL.Path, "/auth/") {
			provider := strings.TrimPrefix(c.Request.URL.Path, "/auth/")
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
