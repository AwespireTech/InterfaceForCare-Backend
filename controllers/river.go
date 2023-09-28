package controllers

import (
	"strconv"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/gin-gonic/gin"
)

type RiverController struct{}

func (riverController RiverController) GetAllRivers(c *gin.Context) {
	river := models.River{}
	c.JSON(200, river)
}
func (riverController RiverController) GetRiverById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid River Id"})
		return
	}
	river := models.River{ID: id}
	c.JSON(200, river)
}
