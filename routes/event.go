package routes

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetEventRoutes(router *gin.RouterGroup) {
	eventController := controllers.EventController{}
	router.GET("/events/:id", eventController.GetEventById)
}
