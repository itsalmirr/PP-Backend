package routers

import (
	"backend.com/go-backend/app/api"
	"backend.com/go-backend/app/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SetupRouter(session_store redis.Store) *gin.Engine {
	r := gin.Default()
	r.Use(sessions.Sessions("auth-session", session_store))

	config.InitOAuth()
	// Public routes
	r.GET("/auth/:provider", api.AuthInit)
	r.GET("/auth/:provider/callback", api.AuthCallback)
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
		listingRoutes := public.Group("/properties")
		{
			listingRoutes.GET("/buy", api.GetListings)
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
