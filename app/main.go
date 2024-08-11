package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"app-ink/app/controller/messages"
	"app-ink/app/core"
	"app-ink/pkg/graceful"
	"app-ink/pkg/rsa256"

	"github.com/gin-gonic/gin"
)

/*
*
TODO:
- [ ] 单元测试
- [ ] 埋点
- [ ] 监控
- [ ] 配置文件
- [ ] 消息队列
- [ ] 分布式
- [ ] 压力测试
- [ ] 性能测试
- [x] Postgres数据库
- [x] 优雅停机
- [x] 错误处理
- [x] 数据校验
- [x] jwt鉴权
- [x] Mongo数据库
- [x] makefile {参考package.json}
- [x] 缓存
- [x] 定时任务
- [x] 依赖注入
- [x] 日志

ENV:
- [x] GIN_MODE
- [x] JWT_SECRET
- [x] POSTGRES_URI
- [x] REDIS_URL
- [ ] RABBIT_MQ_URL
*/

var ctx, cancel = context.WithCancel(context.Background())

func main() {
	core := core.New()

	defer cancel()
	defer core.Dep.Pg.Pool.Close()
	defer core.Dep.Rds.Client.Close()

	initStatic(core.Dep.Env.VolumePath)
	initAsync(core)

	graceful.Listen(
		initEngine(core),
		":8080",
		30*time.Second,
	)
}

func initStatic(VolumePath string) {
	basePem := VolumePath + "/pem"
	baseLog := VolumePath + "/log"
	{
		rsa256.SetBasePath(basePem)
		rsa256.CreateRsaPem()
	}
	{
		err := os.MkdirAll(baseLog, os.ModePerm)
		if err != nil {
			fmt.Println("Create folder fail:", baseLog, err)
			os.Exit(1)
		}
		file, _ := os.Create(baseLog + "/gin.log")
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	}
}

func initAsync(core *core.Core) {
	go messages.New(core).Stream.Listen(ctx)
}

// sudo docker run \
// 	-p 9988:8080 \
// 	-d \
// 	-e JWT_SECRET="asdf" \
// 	-e REDIS_URI="redis://192.168.31.88:6379" \
// 	-e POSTGRES_URI="postgres://postgres:postgres@192.168.31.88:5432/postgres?sslmode=disable" \
// 	app-ink:0
