package core

import (
	"app.ai_painter/app/core/dependency"
	"app.ai_painter/app/core/user"
)

// Core 结构体
// 整个应用核心功能函数
type Core struct {
	Dep  *dependency.Dependency
	User *user.User
}

func New() *Core {
	Dep := dependency.New()

	User := user.New(Dep)

	return &Core{
		Dep:  Dep,
		User: User,
	}
}
