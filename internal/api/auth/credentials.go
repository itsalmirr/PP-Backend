package auth

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn handles user sign-in requests.
// It expects a JSON payload with "email" and "password" fields.
// If the input is invalid, it returns a 400 status code with an error message.
// If the user does not exist or the password is incorrect, it returns a 401 status code with an error message.
// If there is an internal server error, it returns a 500 status code with an error message.
// On successful sign-in, it creates a session and returns a 200 status code with a success message.
// func SignIn(c *gin.Context) {
// 	provider := c.Param("provider")

// 	var input SignInInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input", "message": "Please provide required fields"})
// 		return
// 	}

// 	// Check if user exists
// 	user, err := repositories.GetUserRepository(input.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user", "message": "User not found"})
// 		return
// 	}

// 	// check if password is correct
// 	if match, err := argon2id.ComparePasswordAndHash(input.Password, user.Password); err != nil || !match {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "message": "Incorrect password"})
// 		return
// 	}

// 	session := sessions.Default(c)
// 	session.Set("userEmail", user.Email)
// 	session.Set("authProvider", provider)
// 	if err := session.Save(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "message": err.Error()})
// 		return
// 	}

// 	c.Redirect(http.StatusSeeOther, "/api/v1/users/me")
// }
