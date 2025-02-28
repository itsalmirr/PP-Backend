package routers

import (
	"backend.com/go-backend/internal/api"
	"backend.com/go-backend/internal/api/auth"
	"backend.com/go-backend/internal/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetupRouter(keys *config.Config, db *config.Database) *gin.Engine {
	r := gin.Default()
	// middleware to set database connection in the context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Use(sessions.Sessions("auth-session", config.SessionStorage(keys)), DatabaseMiddleware())
	config.InitOAuth(keys)
	// Public routes
	r.GET("/auth/:provider", auth.AuthInit)
	r.GET("/auth/:provider/callback", auth.AuthCallback)
	public := r.Group("/api/v1")
	{
		// Group of user routes
		userRoutes := public.Group("/users")
		{
			userRoutes.POST("/", api.CreateUser)
			userRoutes.POST("/signin", auth.EmailSignIn)
		}
		// Group of realtor routes
		realtorRoutes := public.Group("/realtors")
		{
			realtorRoutes.GET("/:email", api.GetRealtor)
			realtorRoutes.POST("/", api.CreateRealtor)
		}
		// Group of listings routes
		listingRoutes := public.Group("/properties")
		{
			listingRoutes.GET("/buy", api.GetListings)
			listingRoutes.POST("/", api.CreateListing)
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
		// realtorRoutes := private.Group("/realtors")
		// {
		// 	realtorRoutes.POST("/", api.CreateRealtor)
		// }
		// // Group of listing routes
		// listingRoutes := private.Group("/properties")
		// {
		// 	listingRoutes.POST("/", api.CreateListing)
		// }
	}

	return r
}
