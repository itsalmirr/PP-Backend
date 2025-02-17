package routers

import (
	"os"

	"backend.com/go-backend/app/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetupRouter(session_store redis.Store) *gin.Engine {
	r := gin.Default()
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:8080/auth/google/callback",
		),
	)

	r.Use(sessions.Sessions("auth-session", session_store))
	// Public routes
	public := r.Group("/api/v1")
	{
		// Group of user routes
		userRoutes := public.Group("/users")
		{
			userRoutes.POST("/", api.CreateUser)
			// userRoutes.GET("/:email", api.GetUser)
			userRoutes.POST("/signin", api.SignIn)
		}
		// Group of realtor routes
		realtorRoutes := public.Group("/realtors")
		{
			realtorRoutes.GET("/:email", api.GetRealtor)
		}
		// Group of listings routes
		listingRoutes := public.Group("/listings")
		{
			listingRoutes.GET("/", api.GetListings)
		}
	}

	private := r.Group("/api/v1")
	private.Use(AuthMiddleware())
	{
		// Group of user routes
		userRoutes := private.Group("/users")
		{
			userRoutes.GET("/me", api.Dashboard)
		}
		// Group of realtor routes
		realtorRoutes := private.Group("/realtors")
		{
			realtorRoutes.POST("/", api.CreateRealtor)
		}
		// Group of listing routes
		listingRoutes := private.Group("/listings")
		{
			listingRoutes.POST("/", api.CreateListing)
		}
	}

	return r
}
