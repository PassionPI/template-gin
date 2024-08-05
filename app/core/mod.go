package core

import (
	"app_ink/app/core/common"
	"app_ink/app/core/dependency"
)

// Core 结构体
// 整个应用核心功能函数
type Core struct {
	Dep    *dependency.Dependency
	Common *common.Common
}

func New() *Core {
	Dep := dependency.New()
	Common := common.New()

	return &Core{
		Dep,
		Common,
	}
}
