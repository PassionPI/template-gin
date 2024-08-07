package main

import (
	"context"
	"time"

	"app_ink/app/core"
	"app_ink/pkg/graceful"
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
- [x] MONGODB_URI
- [x] POSTGRES_URI
- [x] REDIS_URL
- [ ] RABBIT_MQ_URL
*/
func main() {
	background := context.Background()
	core := core.New()

	defer core.Dep.Pg.Conn.Close(background)
	defer core.Dep.Rds.Client.Close()
	// defer core.Dep.Mongo.Client.Disconnect(context.TODO())

	initialize(core.Dep.Env.VolumePath)

	graceful.Listen(
		createEngine(core),
		":8080",
		30*time.Second,
	)
}

// sudo docker run \
// 	-p 9988:8080 \
// 	-d \
// 	-e JWT_SECRET="asdf" \
// 	-e REDIS_URI="redis://192.168.31.88:6379" \
// 	-e POSTGRES_URI="postgres://postgres:postgres@192.168.31.88:5432/postgres?sslmode=disable" \
// 	app_ink:0
