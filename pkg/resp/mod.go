package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok 返回成功响应
func Ok(c *gin.Context, data any) {
	c.JSON(
		http.StatusOK,
		gin.H{"data": data},
	)
}

// Err 返回错误响应
func Err(c *gin.Context, message string) {
	c.JSON(
		http.StatusOK,
		gin.H{"error": message},
	)
}
