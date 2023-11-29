package routes

import (
	"github.com/AwespireTech/RiverCare-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.RouterGroup) {
	userController := controllers.UserController{}
	router.GET("/:user/stewardshipTokens", userController.GetStewardshipTokens)
	router.GET("/:user/eventTokens", userController.GetEventTokens)
	router.GET("/:user/hostedEvents", userController.GetHostedEvents)
}
