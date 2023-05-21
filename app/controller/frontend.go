package controller

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Frontend(router *gin.Engine) {
	router.Static("/assets", "./frontend/assets")
	router.StaticFile("/", "./frontend/index.html")
	router.StaticFile("/favicon.svg", "./frontend/favicon.svg")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
}
