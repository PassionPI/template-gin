package user

import (
	"app.ai_painter/app/core/dependency"
)

// Middleware 结构体
// 包含了所有的中间件
type User struct {
	dep *dependency.Dependency
}

// New 创建一个新的中间件
func New(dep *dependency.Dependency) *User {
	return &User{
		dep,
	}
}
