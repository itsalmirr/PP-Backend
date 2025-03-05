package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userEmail")
	session.Delete("authProvider")
	session.Delete("auth-session")

	c.Redirect(http.StatusForbidden, "/")
}
