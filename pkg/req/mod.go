package req

import (
	"fmt"

	"app_land_x/pkg/resp"

	"github.com/gin-gonic/gin"
)

func JSON[T any](c *gin.Context) (val *T, err error) {
	err = c.ShouldBindJSON(&val)
	if err != nil {
		resp.Err(c, fmt.Errorf("Invalid request format %v", err).Error())
	}
	return val, err
}
