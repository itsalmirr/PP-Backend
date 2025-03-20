package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ppgroup.i0sys.com/internal/api"
	"ppgroup.i0sys.com/internal/api/auth"
	"ppgroup.i0sys.com/internal/config"
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
	cors_config := cors.DefaultConfig()
	cors_config.AllowOrigins = []string{"http://localhost:3000"}
	cors_config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	cors_config.AllowCredentials = true
	r.Use(cors.New(cors_config))

	// Set redis session store
	r.Use(sessions.Sessions("auth-session", config.SessionStorage(keys)), DatabaseMiddleware())
	config.InitOAuth(keys)

	// Public routes
	r.GET("/auth/:provider", auth.AuthInit)
	r.GET("/auth/:provider/callback", auth.AuthCallback)
	r.GET("/auth/signout", auth.SignOut)
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

// func SetupGRPCServer(db *config.Database) {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()

// 	proto.RegisterProfilingServer(s, &services.)
// }
