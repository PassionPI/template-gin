package core

import (
	"os"

	"app.land.x/app/common"
	"app.land.x/app/db/mgo"
	"app.land.x/app/db/rds"
	"app.land.x/app/job"
	"app.land.x/pkg/token"
)

// Core 结构体
// 整个应用所有依赖
// 整个应用核心功能函数
type Core struct {
	Rds       *rds.Rds
	Mongo     *mgo.Mongo
	Token     *token.Token
	Common    *common.Common
	Scheduler *job.Job
}

func New() *Core {
	redisURI := os.Getenv("REDIS_URI")   // "redis://localhost:6379/0"
	mongoURI := os.Getenv("MONGODB_URI") // "mongodb://localhost:27017"
	jwtSecret := os.Getenv("JWT_SECRET") // "Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ"

	Rds := rds.New(redisURI)
	Mongo := mgo.New(mongoURI)
	Token := token.New(jwtSecret)
	Common := common.New()
	Scheduler := job.New(Rds.Client)

	return &Core{
		Rds:       Rds,
		Mongo:     Mongo,
		Token:     Token,
		Common:    Common,
		Scheduler: Scheduler,
	}
}
