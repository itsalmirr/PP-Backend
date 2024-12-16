package routers

import (
	"backend.com/go-backend/src/cmd/api"
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
			userRoutes.GET("/:username", api.GetUser)
		}
	}
	return router
}
