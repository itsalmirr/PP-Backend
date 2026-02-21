package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.ppgroup.com/internal/api"
	"ppgroup.ppgroup.com/internal/api/auth"
	"ppgroup.ppgroup.com/internal/config"
	"ppgroup.ppgroup.com/internal/repositories"
	"ppgroup.ppgroup.com/internal/services"
)

func SetupRouter(cfg *config.Config, db *config.Database, imageService *services.ImageService) (*gin.Engine, error) {
	r := gin.Default()

	// Middleware to set database connection and image service in context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("imageService", imageService)
		c.Set("frontendURL", cfg.FrontendURL)
		c.Next()
	})

	// CORS configuration from environment
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{cfg.FrontendURL}
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// Session store
	store, err := config.SessionStorage(cfg)
	if err != nil {
		return nil, fmt.Errorf("setting up session storage: %w", err)
	}
	r.Use(sessions.Sessions("auth-session", store))

	// Database middleware
	r.Use(DatabaseMiddleware())

	// OAuth
	if err := config.InitOAuth(cfg); err != nil {
		return nil, fmt.Errorf("initializing OAuth: %w", err)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Auth routes
	r.GET("/auth/:provider", auth.AuthInit)
	r.GET("/auth/:provider/callback", auth.AuthCallback)
	r.POST("/auth/signout", auth.SignOut)

	// Image upload route
	imageHandler := repositories.NewImageHandler(imageService)
	r.POST("/upload-images", imageHandler.UploadImages)

	public := r.Group("/api/v1")
	{
		userRoutes := public.Group("/users")
		{
			userRoutes.POST("/", api.CreateUser)
			userRoutes.POST("/signin", auth.EmailSignIn)
		}
		realtorRoutes := public.Group("/realtors")
		{
			realtorRoutes.GET("/:email", api.GetRealtor)
			realtorRoutes.GET("/all", api.GetRealtors)
			realtorRoutes.POST("/", api.CreateRealtor)
		}
		listingRoutes := public.Group("/properties")
		{
			listingRoutes.POST("/add", api.CreateListing)
			listingRoutes.POST("/add-json", api.CreateListingJSON)
			listingRoutes.DELETE("/", api.DeleteListing)
			listingRoutes.GET("/buy", api.GetListings)
			listingRoutes.PATCH("/update", api.UpdateListing)
		}
	}

	private := r.Group("/api/v1")
	private.Use(AuthMiddleware())
	{
		userRoutes := private.Group("/users")
		{
			userRoutes.GET("/me", api.Dashboard)
		}
	}

	return r, nil
}
