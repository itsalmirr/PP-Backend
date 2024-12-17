package routers

import (
	"backend.com/go-backend/src/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SetupRouter(session_store redis.Store) *gin.Engine {
	r := gin.Default()
	r.Use(sessions.Sessions("session0x", session_store))
	// Public routes
	public := r.Group("/api/v1")
	{
		// Group of user routes
		userRoutes := public.Group("/users")
		{
			userRoutes.POST("/", api.CreateUser)
			userRoutes.GET("/:email", api.GetUser)
			userRoutes.POST("/signin", api.SignIn)
			userRoutes.GET("/me", api.Dashboard)
		}
		// Group of realtor routes
		realtorRoutes := public.Group("/realtors")
		{
			realtorRoutes.GET("/:email", api.GetRealtor)
		}
	}

	private := r.Group("/api/v1")
	{
		// Group of realtor routes
		realtorRoutes := private.Group("/realtors")
		{
			realtorRoutes.POST("/", api.CreateRealtor)
		}
	}

	return r
}
