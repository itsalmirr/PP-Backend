package routers

import (
	"os"

	"backend.com/go-backend/src/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
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
		Endpoint:     googleOAuth.Endpoint,
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
	return router
}
