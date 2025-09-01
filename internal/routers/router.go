package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.m0chi.com/internal/api"
	"ppgroup.m0chi.com/internal/api/auth"
	"ppgroup.m0chi.com/internal/config"
)

func SetupRouter(keys *config.Config, db *config.Database) *gin.Engine {
	r := gin.Default()
	// middleware to set database connection in the context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Configure CORS to allow cross-origin requests from client source
	// with the listed HTTP methods and credentials
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// Set redis session store
	r.Use(sessions.Sessions("auth-session", config.SessionStorage(keys)))
	// Apply database middleware to make entClient available in all routes
	r.Use(DatabaseMiddleware())
	config.InitOAuth(keys)

	// Public routes
	r.GET("/auth/:provider", auth.AuthInit)
	r.GET("/auth/:provider/callback", auth.AuthCallback)
	r.POST("/auth/signout", auth.SignOut)
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
			realtorRoutes.GET("/all", api.GetRealtors)
			realtorRoutes.POST("/", api.CreateRealtor)
		}
		// Group of listings routes
		listingRoutes := public.Group("/properties")
		{
			listingRoutes.POST("/add", api.CreateListing)
			listingRoutes.DELETE("/", api.DeleteListing)
			listingRoutes.GET("/buy", api.GetListings)
			listingRoutes.PATCH("/update", api.UpdateListing)
		}
	}

	// Private routes
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
		{
			// realtorRoutes.POST("/", api.CreateRealtor)
		}
		// // Group of listing routes
		// listingRoutes := private.Group("/properties")
		// {
		// 	listingRoutes.POST("/", api.CreateListing)
		// }
	}

	return r
}
