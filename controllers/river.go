package controllers

import (
	"strconv"

	"github.com/AwespireTech/InterfaceForCare-Backend/database"
	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/gin-gonic/gin"
)

type RiverController struct{}

func (rc RiverController) GetAllRivers(c *gin.Context) {
	params := models.RiversParams{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	rivers, err := database.GetAllRivers(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := models.RiversResponse{Rivers: rivers, Count: len(rivers)}
	c.JSON(200, resp)

}
func (riverController RiverController) GetRiverById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid River Id"})
		return
	}
	river, err := database.GetRiverById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, river)
}
func (rc RiverController) GetAllProposals(c *gin.Context) {
	rid := c.Param("id")
	riverId, err := strconv.Atoi(rid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid River Id"})
		return
	}
	river, err := database.GetRiverById(riverId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, river.Proposals)
}
