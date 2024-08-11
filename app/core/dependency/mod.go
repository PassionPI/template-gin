package dependency

import (
	"os"

	"app-ink/app/core/corn"
	"app-ink/app/service/pg"
	"app-ink/app/service/rds"
	"app-ink/pkg/token"
)

type env struct {
	AppName     string
	RedisURI    string
	JwtSecret   string
	PostgresURI string
	VolumePath  string
}

func getEnv() *env {
	AppName := os.Getenv("IMAGE")            // "app-ink"
	RedisURI := os.Getenv("REDIS_URI")       // "redis://localhost:6379/0"
	JwtSecret := os.Getenv("JWT_SECRET")     // "Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ"
	PostgresURI := os.Getenv("POSTGRES_URI") // "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	return &env{
		AppName:     AppName,
		RedisURI:    RedisURI,
		JwtSecret:   JwtSecret,
		PostgresURI: PostgresURI,
		VolumePath:  "./private",
	}
}

// Dependency 结构体
// 整个应用所有依赖
type Dependency struct {
	Env   *env
	Rds   *rds.Rds
	Pg    *pg.Pg
	Cron  *corn.Cron
	Token *token.Token
}

// New Dependency 结构体
func New() *Dependency {
	Env := getEnv()
	Rds := rds.New(Env.RedisURI)
	Pg := pg.New(Env.PostgresURI)
	Cron := corn.New(Rds.Client)
	Token := token.New(Env.JwtSecret)

	return &Dependency{
		Env:   Env,
		Rds:   Rds,
		Pg:    Pg,
		Cron:  Cron,
		Token: Token,
	}
}
