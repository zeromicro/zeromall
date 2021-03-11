
# event queue:

- 通用队列管理: 事件响应

## 应用场景:

- 部分后端服务: 不需要在服务中开 MQ 的, 可以通过简单的 http api 调用
- 前端: 直接通过 HTTP API 实现异步能力.
- 类似 graphQL 把 DB 能力开放给前端, 这里把 MQ 能力, 开放给前端.
- 签名机制: 防攻击
    - MQ 里的消息, 增加校验, 签名无效的消息, 丢掉(防止伪造)
    - 签名规则: api secret, timestamp, nonce, msg_uuid, data. 防重放攻击(消费者需要 cache msg_uuid, 去重)
- API 认证方式:
    - http header: token 认证

- auth 认证:
    - 用户认证: user auth
    - 服务认证: service auth
- 权限控制:
    - 用户权限: user permission
    - 服务权限: service permission


## 运行:

- run server:

```bash

make run.auto.reload


```

- http api test:

```bash

make curl.pub.msg

```

## API List:

- `app/basic/queue/internal/router/router.go`

- 队列写入消息: http://localhost:8080/queue/publish/
    - POST
    - Makefile 有 curl 示例


## 功能:

### 队列管理:

- [ ] 队列 创建
- [ ] 队列 auth-key 生成/鉴权
- [ ] 队列 销毁

### 生产者:

- [ ] 队列 写: HTTP API/ GRPC
- [ ] 队列 读: HTTP API 轮训 / websocket 监听
- [ ] 把 `task` 和 `task handler function` 都当成参数, 传给 server.
-

### 消费者: (worker)

- [ ] 队列 写: HTTP API/ GRPC
- [ ] 队列 读: HTTP API 轮训 / websocket 监听
- [ ] task handler

### task handler function:

- 通用的 task handler 方法
- 归纳常用的一些 task 处理方法:
    - 异步 记录数据到 db, 并返回处理成功
    - 异步 send mobile msg
    - 异步 数据分析: 生成报表
    - 异步 ...
- 将标准化的方法, 当成参数, 对外提供
- 类似 `AWS Lambda`: https://aws.amazon.com/cn/lambda/

