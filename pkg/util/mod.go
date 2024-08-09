package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON 解析请求体
// 给定类型，自动解析请求体
func JSON[T any](c *gin.Context) (val T, err error) {
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
		data,
	)
}

// Err 返回错误响应
func Err(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		message,
	)
}

func Bad(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		message,
	)
}

func NotFound(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		message,
	)
}

func UnAuth(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		message,
	)
}
