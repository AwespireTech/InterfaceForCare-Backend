package controllers

import (
	"github.com/AwespireTech/InterfaceForCare-Backend/database"
	"github.com/gin-gonic/gin"
)

type EventController struct{}

func (ec EventController) GetEventById(c *gin.Context) {
	id := c.Param("id")
	event, err := database.GetEventById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, event)

}
