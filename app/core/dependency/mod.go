package dependency

import (
	"os"

	"app_ink/app/common"
	"app_ink/app/service/mgo"
	"app_ink/app/service/rds"
	"app_ink/app/tasks"
	"app_ink/pkg/token"
)

type env struct {
	RedisURI   string
	MongoURI   string
	JwtSecret  string
	AppName    string
	VolumePath string
}

func getEnv() *env {
	AppName := os.Getenv("IMAGE")        // "app_ink"
	RedisURI := os.Getenv("REDIS_URI")   // "redis://localhost:6379/0"
	MongoURI := os.Getenv("MONGODB_URI") // "mongodb://localhost:27017"
	JwtSecret := os.Getenv("JWT_SECRET") // "Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ"

	return &env{
		RedisURI:   RedisURI,
		MongoURI:   MongoURI,
		JwtSecret:  JwtSecret,
		AppName:    AppName,
		VolumePath: "./private",
	}
}

// Dependency 结构体
// 整个应用所有依赖
type Dependency struct {
	Env       *env
	Rds       *rds.Rds
	Mongo     *mgo.Mongo
	Token     *token.Token
	Common    *common.Common
	Scheduler *tasks.Tasks
}

// New Dependency 结构体
func New() *Dependency {
	Env := getEnv()

	Rds := rds.New(Env.RedisURI)
	Mongo := mgo.New(Env.MongoURI, Env.AppName)
	Token := token.New(Env.JwtSecret)
	Common := common.New()
	Scheduler := tasks.New(Rds.Client)

	return &Dependency{
		Env:       Env,
		Rds:       Rds,
		Mongo:     Mongo,
		Token:     Token,
		Common:    Common,
		Scheduler: Scheduler,
	}
}
