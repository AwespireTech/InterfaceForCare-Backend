package routes

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetRiverRoutes(router *gin.RouterGroup) {
	riverController := controllers.RiverController{}
	router.GET("/rivers", riverController.GetAllRivers)
	router.GET("/rivers/:id", riverController.GetRiverById)
	router.GET("/river/:id/proposals", riverController.GetAllProposals)
}
