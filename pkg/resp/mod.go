package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, data *gin.H) {
	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

func Err(c *gin.Context, message string) {
	c.JSON(
		http.StatusOK,
		gin.H{"error": message},
	)
}
