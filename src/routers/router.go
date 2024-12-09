package routers

import (
	"net/http"

	"backend.com/go-backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api/v1")
	{
		public.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})

		public.POST("/users", controllers.CreateUser)
	}

	return router
}
