package main

import (
	"github.com/AwespireTech/RiverCare-Backend/config"
	"github.com/AwespireTech/RiverCare-Backend/database"
	"github.com/AwespireTech/RiverCare-Backend/routes"
	"github.com/AwespireTech/RiverCare-Backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.PrintConfig()
	err := database.Init(config.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	router := createRouter()
	router.Run()
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	river := router.Group("api")
	{
		routes.SetIPFSRoutes(river)
		routes.SetRiverRoutes(river)
		routes.SetUserRoutes(river)
		routes.SetEventRoutes(river)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to RiverCare API",
		})
	})
	return router
}
