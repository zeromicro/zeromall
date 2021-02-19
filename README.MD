
# zero mall

- 基于 `go-zero` 实现的 通用电商平台.


## Requirements:

- go 1.15+
- go-zero
- gRPC
- mysql
- redis
- kafka


## QuickStart:

- https://www.yuque.com/tal-tech/go-zero/yaoehb


### Development Environment:

- install:

```

# clone repo:
git clone git@github.com:zeromicro/zeromall.git

# enter project root:
cd zeromall/

# install require pkg:
make go.mod.tidy

```


- run 示例服务: `app/basic/demo`

```
cd ./app/basic/demo

# 启动 API 服务:
make run.demo

# 或者使用: 基于 air, watch 该服务目录, 修改后, 自动 hot reload, 无需重新编译
make run.auto.reload


# http echo test:
make curl.api


#
# will print:
#
-> % make curl.api
curl -i http://localhost:8888/from/you
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 18 Feb 2021 10:48:23 GMT
Content-Length: 25

{"message":"hello world"}%


```


## Structure:

- 目录结构简略说明: 对整体目录的 quick look.


```html


-> % tree . -L 3
.
├── app
│   ├── basic            // 基础服务: 通常是业务无关的
│   │   ├── demo         //     示例:
│   │   └── user         //     用户管理:
│   ├── biz              // 业务服务: 根据业务切分
│   │   └── member       //     会员服务:
│   │   └── order        //     订单服务:
│   │   └── cart         //     购物车服务:
│   └── std              // 业务内统一定义:
│       ├── proto        //     统一定义业务状态码
├── deploy               // 部署脚本, 服务编排
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── infra                // 基础设施
│   └── tool
└── pkg                  // 项目积累的 utility 代码库, 与业务无关


```


- 代码调用链路说明: 以 `app/basic/demo` 服务为例(自顶向下调用)
    - 启动入口: `app/basic/demo/cmd/main.go`
    - Service 入口: `app/basic/demo/internal/service/service.go`
        - 对外 API 网关: `app/basic/demo/internal/service/outer/outer.go`
        - 对内 RPC 网关: `app/basic/demo/internal/service/inner/inner.go`
        - 对内 Admin API 网关: `app/basic/demo/internal/service/admin/admin.go`
        - 对内 Job 网关: `app/basic/demo/internal/service/job/job.go`
    - 数据定义:
        - API 层: `app/basic/demo/proto/api/api.proto`
        - Config 层: `app/basic/demo/proto/config/config.proto`
        - Model 层: `app/basic/demo/proto/model/model.proto`
    - 路由注册: `app/basic/demo/internal/router/router.go`
    - 业务领域 Domain:
        - 某业务单元: `app/basic/demo/internal/domain/demo/demo.go`
    - Dao 数据层操作:
        - Meta 聚合(收敛 Dao 层资源): `app/basic/demo/internal/dao/meta.go`
        - DB 层: `app/basic/demo/internal/dao/db/db.go`
        - Cache 层: `app/basic/demo/internal/dao/cache/cache.go`


- 目录结构详细说明:



```html



-> % tree . -L 6
.
├── LICENSE
├── Makefile
├── README.MD
├── app
│   ├── basic                                       // 基础服务单元1 : 用户管理/推送/短信/等业务无关的通用服务
│   │   ├── demo                                    // 单个微服务示例:
│   │   │   ├── cmd                                 //      单服务启动入口: 包含 多个启动方式:
│   │   │   │   └── main.go                         // api server, grpc server, job server, admin server 启动
│   │   │   ├── configs                             //      外部中间件配置项: db,缓存,mq 等
│   │   │   │   └── configs.toml
│   │   │   ├── docs                                // 单服务自身文档
│   │   │   ├── internal                            //      业务逻辑(不对外暴露)
│   │   │   │   ├── dao                             //      数据层 read/write
│   │   │   │   │   ├── cache                       //      缓存 r/w
│   │   │   │   │   ├── db                          //      db crud
│   │   │   │   │   ├── http                        //      调用本服务之外的 http api
│   │   │   │   │   ├── meta.go                     //      dao 资源收敛
│   │   │   │   │   ├── mq                          //      mq r/w
│   │   │   │   │   └── rpc                         //      调用本服务之外的 rpc(gRPC) api
│   │   │   │   ├── domain                          // 服务内业务拆分:
│   │   │   │   │   └── demo                        //      业务单元1 / 业务单元2
│   │   │   │   └── service                         // API 收敛层(对外收敛内部逻辑, 暴露 API: grpc/http/job/admin)
│   │   │   │       └── service.go
│   │   │   └── proto                               // 数据定义层: (可对外暴露)
│   │   │       ├── api                             //      grpc + http api 定义
│   │   │       │   └── api.proto
│   │   │       ├── config                          //      config toml 映射 model
│   │   │       │   └── config.proto
│   │   │       └── model                           //      内部 model
│   │   │           └── model.proto
│   │   └── user                                    // 基础服务2:
│   │       ├── identity                            //   服务内
│   │       │   ├── Makefile
│   │       │   ├── cmd
│   │       │   │   └── main.go
│   │       │   ├── configs
│   │       │   │   └── configs.toml
│   │       │   ├── docs
│   │       │   ├── internal
│   │       │   │   ├── dao
│   │       │   │   ├── domain
│   │       │   │   └── service
│   │       │   ├── proto
│   │       │   │   ├── api
│   │       │   │   ├── config
│   │       │   │   └── model
│   │       │   └── readme.md
│   │       └── readme.md
│   ├── biz                                         // 具体业务单元:
│   │   │── member                                  // 会员服务
│   │   │   └── cmd
│   │   │       └── main.go
│   │   └── order                                   // 订单服务
│   │       └── cmd
│   │           └── main.go
│   └── std                                        // 项目业务内统一定义(业务状态码, 业务出错 msg 编号)
│       ├── Makefile
│       ├── proto
│       │   ├── config
│       │   │   └── config.proto
│       │   └── error
│       │       └── code.proto                     // 业务状态码
│       └── readme.md
├── deploy                                         // 部署相关脚本: dockerfile 等
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── go.mod
├── go.sum
├── infra                                          // 插件依赖服务: cli, 中间件等
│   └── tool
└── pkg                                            // 项目中逐步积累的 utility 代码库

46 directories, 24 files


```


## Reference:


- https://github.com/macrozheng/mall




