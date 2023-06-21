package qp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON 解析请求体
// 给定类型，自动解析请求体
func JSON[T any](c *gin.Context) (val *T, err error) {
	err = c.ShouldBindJSON(&val)
	if err != nil {
		Err(c, fmt.Errorf("Invalid request format %v", err).Error())
	}
	return val, err
}

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
