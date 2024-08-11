package core

import (
	"app-ink/app/core/common"
	"app-ink/app/core/dependency"
	"app-ink/app/core/sender"
)

// Core 结构体
// 整个应用核心功能函数
type Core struct {
	Dep    *dependency.Dependency
	Common *common.Common
	Sender *sender.Sender
}

func New() *Core {
	Dep := dependency.New()
	Common := common.New(Dep)
	Sender := sender.New(Dep)

	return &Core{
		Dep,
		Common,
		Sender,
	}
}
