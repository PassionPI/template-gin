package main

import (
	"time"

	"app_land_x/pkg/graceful"
)

/*
*
TODO:
- [x] 优雅停机
- [x] 错误处理
- [x] 数据校验
- [x] jwt鉴权
- [x] Mongo数据库
- [x] makefile {参考package.json}
- [ ] 缓存
- [ ] 消息队列
- [ ] 定时任务
- [ ] Postgres数据库
- [ ] 依赖注入
- [ ] 分布式
- [ ] 配置文件
- [ ] 监控
- [ ] 日志
- [ ] 埋点
- [ ] 单元测试
- [ ] 性能测试
- [ ] 压力测试

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
