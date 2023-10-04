package main

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := createRouter()

	router.Run(":8080")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	river := router.Group("api")
	{
		routes.SetRiverRoutes(river)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to InterfaceForCare API",
		})
	})
	return router
}
