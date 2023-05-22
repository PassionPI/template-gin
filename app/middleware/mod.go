package middleware

import (
	"app.land.x/app/controller"

	"github.com/gin-gonic/gin"
)

// Middleware 结构体
// 包含了所有的中间件
type Middleware struct {
	ctrl *controller.Controller
}

// New 创建一个新的中间件
func New(ctrl *controller.Controller) *Middleware {
	return &Middleware{
		ctrl,
	}
}

// Empty 空中间件
func (m *Middleware) Empty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
