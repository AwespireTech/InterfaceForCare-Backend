package controllers

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/database"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc UserController) GetStewardshipTokens(c *gin.Context) {
	user := c.Param("user")
	hist, err := database.GetStewardshipTokensByUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, hist)
}
func (uc UserController) GetEventTokens(c *gin.Context) {
	user := c.Param("user")
	hist, err := database.GetEventTokensByUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, hist)

}
func (uc UserController) GetHostedEvents(c *gin.Context) {
	user := c.Param("user")
	events, err := database.GetHostedEvents(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, events)

}
