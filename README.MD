# zero mall

- [x] 基于 `go` + `go-zero` + `gRPC`, 以`微服务`方式实现的`B2B2C`通用电商平台.
- [x] 目标:
    - [x] 综合社区`最佳实践`, 力求`产品级`项目, 而不是 `demo 玩具`.
    - [x] 通用业务单元: 可以自由组合使用.

## WIP:

- 🚀 架构公式: [biz](app/biz) = [core](app/core) + [unit](app/unit)

> 产品列表:

- 💖 [app/biz/mall](app/biz/mall): B2B2C 电商平台
- 💖 [app/biz/eshop](app/biz/eshop): 虚拟商品售卖商城

> 服务列表:

- ☀️ [app/unit/license/key](./app/unit/license/key): 注册码服务
    - 售卖 app `密钥/注册码`服务

> 基础服务列表:

- 🔧 [app/core/queue](app/core/queue): 队列服务

## Wiki:

- [ZeroMall - wiki](https://github.com/zeromicro/zeromall/wiki)
- 项目文档

## Client:

- [client](./client)
- https://github.com/zeromicro/zeromall/wiki/client
- 基于 flutter 实现客户端(Web/iOS/Android/Windows/MacOS/Linux)

### Requirements:

- [x] flutter 2.5.0+
- [x] GetX: 主要框架

## Server:

- 基于 go + go-zero 实现
- [app](./app)

### Requirements:

- [x] go 1.15+
- [x] go-zero: 全功能的 go 微服务框架
- [x] gRPC: RPC 领域事实上的标准
- [x] etcd: 服务注册/服务发现
- [x] mysql: 存储
- [x] redis: 缓存
- [x] kafka: 队列
- [x] rabbitmq: 队列
- [x] docker + docker-compose

> 这也是掌握本项目, 需要熟知的先修技能栈.

#### Development Environment:

- [x] [reference - wiki](https://github.com/zeromicro/zeromall/wiki/reference)
- [x] [搭建本地开发环境 - wiki](https://github.com/zeromicro/zeromall/wiki/setup-development)
    - 请务必仔细查看文档, 有详细说明
- [deploy/docs](./deploy/readme.md)


- install:

```

# clone repo:
git clone git@github.com:zeromicro/zeromall.git

# enter project root:
cd zeromall/


# run mysql/redis/rabbitmq/etcd:
task local:init

# install:
make install
make init

# install go pkg:
task tidy


# run:
task run


```

- 创建新服务目录:
    - [go-zero 微服务创建脚手架: cookiecutter-go](https://github.com/better-go/cookiecutter-go)

```bash
cd zeromall/

# 在 app/biz 下创建一个新的业务服务目录:
make app.gen.biz

# 在 app/basic 下创建一个新的基础服务目录:
make app.gen.basic
```

#### 示例服务:

- [x] [示例服务 app/basic/demo - wiki](https://github.com/zeromicro/zeromall/wiki/demo-service)
- [x] 服务启动说明, 辅助理解代码结构

### Structure:

- [项目目录结构详细说明 - wiki](https://github.com/zeromicro/zeromall/wiki/project-structure)
- 非常详细的目录说明, 快速理解项目

## RoadMap:

- https://github.com/zeromicro/zeromall/issues/3
- https://github.com/zeromicro/zeromall/projects/1


