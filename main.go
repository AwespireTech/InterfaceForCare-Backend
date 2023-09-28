package main

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	river := router.Group("api")
	{
		routes.SetRiverRoutes(river)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to InterfaceForCare API",
		})
	})

	router.Run(":8080")
}
