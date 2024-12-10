package routers

import (
	"backend.com/go-backend/src/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginoauth2 "github.com/zalando/gin-oauth2"

	"github.com/zalando/gin-oauth2/google"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"os"
)

func SetupRouter() *gin.Engine {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}

	// OAuth2 configuration
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/callback",
	}

	router := gin.Default()

	// Public routes
	public := router.Group("/api/v1")
	{
		// Group of user routes
		userRoutes := public.Group("/users")
		{
			userRoutes.POST("/", controllers.CreateUser)
			userRoutes.GET("/:username", controllers.GetUser)
		}
	}

	// Protected routes with OAuth2 middleware
	protected := router.Group("/api/v1")
	protected.Use(ginoauth2.Auth(conf))
	{
		protected.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized"})
		})
	}

	return router
}
