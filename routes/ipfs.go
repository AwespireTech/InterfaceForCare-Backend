package routes

import (
	"github.com/AwespireTech/RiverCare-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetIPFSRoutes(router *gin.RouterGroup) {
	ipfsController := controllers.IPFSController{}
	router.GET("/ipfs/*path", ipfsController.Forward)
}
