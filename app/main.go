package main

import (
	"time"

	"app.ai_painter/pkg/graceful"
)

/*
*
TODO:
- [ ] 单元测试
- [ ] 埋点
- [ ] 监控
- [ ] 配置文件
- [ ] 消息队列
- [ ] Postgres数据库
- [ ] 分布式
- [ ] 压力测试
- [ ] 性能测试
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
- [ ] REDIS_URL
- [ ] RABBIT_MQ_URL
*/
func main() {
	graceful.Listen(
		createEngine(),
		":8080",
		30*time.Second,
	)
}
