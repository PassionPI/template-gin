package req

import (
	"fmt"

	"app.land.x/pkg/resp"

	"github.com/gin-gonic/gin"
)

// JSON 解析请求体
// 给定类型，自动解析请求体
func JSON[T any](c *gin.Context) (val *T, err error) {
	err = c.ShouldBindJSON(&val)
	if err != nil {
		resp.Err(c, fmt.Errorf("Invalid request format %v", err).Error())
	}
	return val, err
}
