package controllers

import (
	"net/url"

	"github.com/AwespireTech/RiverCare-Backend/config"
	"github.com/gin-gonic/gin"
)

type IPFSController struct{}

func (ipc IPFSController) Forward(c *gin.Context) {
	target, _ := url.Parse(config.IPFS_URL)
	target.User = url.UserPassword("ipfs", "ipfs")
}
