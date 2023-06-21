package remote

import (
	"app.land.x/app/core/dependency"
)

// Middleware 结构体
// 包含了所有的中间件
type Remote struct {
	dep *dependency.Dependency
}

// New 创建一个新的中间件
func New(dep *dependency.Dependency) *Remote {
	return &Remote{
		dep,
	}
}

func (r *Remote) RunCommand() string {
	// return r.dep.Common.GetRemoteIP()
	return "hello"
}
