package controllers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/AwespireTech/RiverCare-Backend/config"
	"github.com/gin-gonic/gin"
)

type IPFSController struct{}

func (ipc IPFSController) Forward(c *gin.Context) {
	target, err := url.Parse(config.IPFS_URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	target.User = url.UserPassword(config.IPFS_USER, config.IPFS_PASSWORD)
	proxy := httputil.ReverseProxy{Director: func(req *http.Request) {
		req.URL = target
		req.Host = target.Host
	}}
	proxy.ServeHTTP(c.Writer, c.Request)

}
