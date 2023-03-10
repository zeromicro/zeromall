# zero mall

- [x] 基于 `go` + `go-zero` + `gRPC`, 以`微服务`方式实现的`B2B2C`通用电商平台.
- [x] 目标:
    - [x] 综合社区`最佳实践`, 力求`产品级`项目, 而不是 `demo 玩具`.
    - [x] 通用业务单元: 可以自由组合使用.

## WIP:

- 🚀 架构公式: [biz](app/biz) = [core](app/core) + [unit](app/unit)
    - ✅ [biz](app/biz): 产品构成
    - ✅ [core](app/core): 基础服务, 公共服务
    - ✅ [unit](app/unit): 业务单元, 通用业务单元

### 产品:

| [产品(app/biz) ](app/biz)           | 说明         |
|:----------------------------------|------------|
| 💖 [app/biz/mall](app/biz/mall)   | B2B2C 电商平台 |
| 💖 [app/biz/eshop](app/biz/eshop) | 虚拟商品售卖商城   |

### 核心服务:

| [核心服务(app/core)](app/core)                 | 服务           |
|:-------------------------------------------|--------------|
| [app/core/user/authn](app/core/user/authn) | 用户服务: 注册/登录  |
| [app/core/user/authz](app/core/user/authz) | 用户服务: 访问权限控制 |
| [app/core/queue](app/core/queue)           | 基础队列服务       |

### 服务矩阵:

| [业务单元(app/unit)](app/unit)                              | 说明                 |
|:--------------------------------------------------------|--------------------|
| ☀️ [app/unit/license/key](app/unit/license/key)         | 注册码服务: 售卖 app `密钥` |
| ☀️ [app/unit/cart](app/unit/cart)                       | 购物车系统              |
| ☀️ [app/unit/order](app/unit/order)                     | 订单系统               |
| ☀️ [app/unit/finance/payment](app/unit/finance/payment) | 支付系统               |

## Development:

- ✅ [Taskfile.yml](./Taskfile.yml)
- ✅ 所有`微服务`启动脚本, 依赖`中间件`启动, 都在此.

### Server:

- [x] [Setup Server Development Environment](./docs/dev/setup.server.md)

### Client:

- [x] [Setup Client Development Environment](./docs/dev/setup.client.md)

## 示例服务:

- ✅ [app/unit/license](app/unit/license) : 授权码服务(app license key)
    - [app/unit/license/key/Taskfile.yml](app/unit/license/key/Taskfile.yml)
- ✅ 以此服务为例, 可以快速熟悉整个项目的架构.
- ✅ 每个服务内都有 [Taskfile.yml](app/unit/license/key/Taskfile.yml), 包含服务所有启动脚本.







