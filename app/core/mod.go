package core

import (
	"os"

	"app.land.x/app/api/mgo"
	"app.land.x/app/api/rds"
	"app.land.x/app/common"
	"app.land.x/app/core/dependency"
	"app.land.x/app/core/remote"
	"app.land.x/app/tasks"
	"app.land.x/pkg/token"
)

// Core 结构体
// 整个应用核心功能函数
type Core struct {
	Dep    *dependency.Dependency
	Remote *remote.Remote
}

func createDependency() *dependency.Dependency {
	redisURI := os.Getenv("REDIS_URI")   // "redis://localhost:6379/0"
	mongoURI := os.Getenv("MONGODB_URI") // "mongodb://localhost:27017"
	jwtSecret := os.Getenv("JWT_SECRET") // "Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ"

	Rds := rds.New(redisURI)
	Mongo := mgo.New(mongoURI)
	Token := token.New(jwtSecret)
	Common := common.New()
	Scheduler := tasks.New(Rds.Client)

	return &dependency.Dependency{
		Rds:       Rds,
		Mongo:     Mongo,
		Token:     Token,
		Common:    Common,
		Scheduler: Scheduler,
	}
}

func New() *Core {
	Dep := createDependency()

	Remote := remote.New(Dep)

	return &Core{
		Dep:    Dep,
		Remote: Remote,
	}
}
