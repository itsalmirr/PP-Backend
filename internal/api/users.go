package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.i0sys.com/ent"
	"ppgroup.i0sys.com/internal/repositories"
)

// CreateUser handles the creation of a new user.
// @Summary Create a new user
// @Description This endpoint creates a new user with the provided input data.
// @Tags users
// @Accept json
// @Produce json
// @Param user body repositories.CreateUserInput true "User input data"
// @Success 200 {object} map[string]interface{} "status: OK, data: User created!"
// @Failure 400 {object} map[string]interface{} "error: Invalid input, message: Please provide required fields"
// @Failure 500 {object} map[string]interface{} "error: Failed to create user, message: error message"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input *ent.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"message": "Please provide required fields",
		})
		return
	}

	entClient := c.MustGet("entClient").(*ent.Client)
	err := repositories.CreateUserRepo(entClient, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   "User created!",
	})
}

// Dashboard handles requests to the dashboard endpoint.
// It retrieves the user email from the session, ensuring that the user is authenticated.
// If the email is missing or empty, it responds with an "Unauthorized" JSON error.
// It then fetches the user details from the repository using the email. If an error occurs while fetching user data,
// it responds with an Internal Server Error. Upon a successful fetch, it clears the user's password field
// before returning the user data as a JSON response.
func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	email, ok := session.Get("userEmail").(string)
	if !ok || email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Please sign in"})
		return
	}
	entClient := c.MustGet("entClient").(*ent.Client)
	user, err := repositories.GetUserRepo(entClient, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user", "message": err.Error()})
		return
	}

	// Remove the password hash from the user object
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"data": user})
}
