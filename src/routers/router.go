package routers

import (
	"backend.com/go-backend/src/controllers"
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
			userRoutes.POST("/", controllers.CreateUser)
			userRoutes.GET("/:username", controllers.GetUser)
		}
	}

	return router
}
