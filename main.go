package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Run(":8080")
}
