package controller

import "app_ink/app/core"

// Controller 结构体
// 包含了所有的控制器
type Controller struct {
	core *core.Core
}

// New 创建一个新的中间件
func New(core *core.Core) *Controller {
	return &Controller{
		core,
	}
}
