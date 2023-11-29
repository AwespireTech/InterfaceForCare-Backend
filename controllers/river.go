package controllers

import (
	"net/http"

	"github.com/AwespireTech/RiverCare-Backend/database"
	"github.com/AwespireTech/RiverCare-Backend/models"
	"github.com/gin-gonic/gin"
)

type RiverController struct{}

func (rc RiverController) GetAllRivers(c *gin.Context) {
	params := models.RiversParams{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		if err.Error() == "EOF" {
			params.Ascending = true
			params.SortBy = "createdTime"
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	rivers, err := database.GetAllRivers(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := models.RiversResponse{Rivers: rivers, Count: len(rivers)}
	c.JSON(http.StatusOK, resp)

}
func (riverController RiverController) GetRiverById(c *gin.Context) {
	rid := c.Param("id")
	river, err := database.GetRiverById(rid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, river)
}
func (rc RiverController) GetAllProposals(c *gin.Context) {
	rid := c.Param("id")
	river, err := database.GetRiverById(rid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, river.ProposalData)
}
