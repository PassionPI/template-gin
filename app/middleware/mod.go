package middleware

import (
	"app.ai_painter/app/core"

	"github.com/gin-gonic/gin"
)

// Middleware 结构体
// 包含了所有的中间件
type Middleware struct {
	core *core.Core
}

// New 创建一个新的中间件
func New(core *core.Core) *Middleware {
	return &Middleware{
		core,
	}
}

// Empty 空中间件
func (m *Middleware) Empty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
