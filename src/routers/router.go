package routers

import (
	"backend.com/go-backend/src/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Public routes
	public := router.Group("/api/v1")
	{
		// Group of user routes
		userRoutes := public.Group("/users")
		{
			userRoutes.POST("/", api.CreateUser)
			userRoutes.GET("/:email", api.GetUser)
		}
		// Group of realtor routes
		realtorRoutes := public.Group("/realtors")
		{
			realtorRoutes.POST("/", api.CreateRealtor)
			realtorRoutes.GET(":email", api.GetRealtor)
		}
	}
	return router
}
