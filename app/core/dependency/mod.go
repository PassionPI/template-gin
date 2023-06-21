package dependency

import (
	"os"

	"app.land.x/app/api/mgo"
	"app.land.x/app/api/rds"
	"app.land.x/app/common"
	"app.land.x/app/tasks"
	"app.land.x/pkg/token"
)

// Dependency 结构体
// 整个应用所有依赖
type Dependency struct {
	Rds       *rds.Rds
	Mongo     *mgo.Mongo
	Token     *token.Token
	Common    *common.Common
	Scheduler *tasks.Tasks
}

func New() *Dependency {
	redisURI := os.Getenv("REDIS_URI")   // "redis://localhost:6379/0"
	mongoURI := os.Getenv("MONGODB_URI") // "mongodb://localhost:27017"
	jwtSecret := os.Getenv("JWT_SECRET") // "Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ"

	Rds := rds.New(redisURI)
	Mongo := mgo.New(mongoURI)
	Token := token.New(jwtSecret)
	Common := common.New()
	Scheduler := tasks.New(Rds.Client)

	return &Dependency{
		Rds:       Rds,
		Mongo:     Mongo,
		Token:     Token,
		Common:    Common,
		Scheduler: Scheduler,
	}
}
