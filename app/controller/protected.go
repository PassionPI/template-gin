package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Protected(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(
			http.StatusForbidden,
			gin.H{"message": "Welcome to the protected route!"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"message":  "Welcome to the protected route!",
			"username": username,
		},
	)
}
