package routes

import (
	"github.com/AwespireTech/RiverCare-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetEventRoutes(router *gin.RouterGroup) {
	eventController := controllers.EventController{}
	router.GET("/events/:id", eventController.GetEventById)
}
