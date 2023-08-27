package remote

import (
	"app.land.x/app/core/dependency"
)

// Remote 结构体
// 包含了所有的 Remote机器 操作
type Remote struct {
	dep *dependency.Dependency
}

// New 创建一个新的中间件
func New(dep *dependency.Dependency) *Remote {
	return &Remote{
		dep,
	}
}
