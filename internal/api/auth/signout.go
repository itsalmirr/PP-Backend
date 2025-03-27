package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	c.SetCookie(session.ID(), "", -1, "/", "", false, true)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save session",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
