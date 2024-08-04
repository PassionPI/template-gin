package core

import (
	"app_ink/app/core/dependency"
)

// Core 结构体
// 整个应用核心功能函数
type Core struct {
	Dep *dependency.Dependency
}

func New() *Core {
	Dep := dependency.New()

	return &Core{
		Dep: Dep,
	}
}
