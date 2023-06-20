# template-gin

### 项目说明

简易模版, 集成了`优雅停机`、`限流`、`登录`、`jwt`、`mongo`、`redis`、`定时任务`

### 目录说明

private - 私有内容, 用于存放 公钥 私钥 日志等等

pkg - 纯函数包

app - 项目应用

app/common - 公共常量定义

app/controller - 路由控制器、用于处理副作用

app/core - 应用核心依赖、方法

app/db - 数据库及其封装方法

app/job - 定时任务

app/middleware - 路由中间件

app/model - 应用数据模型

app/view - 视图数据模型
