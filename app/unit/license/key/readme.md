# key:

- 授权码 / 注册码 生成服务

## 服务列表：

- 注册码生成
- 注册码分配
- 注册码管理

## 说明：

- 此示例 repo， 可以当作模板
- 修复了 goctl 生成的一些不好的点， rpc 示例， 可以当模板

```ruby
cd learn-go/

# pb 新增api， 动态生成新的模板代码：
task try:zero:gen:pb

```

## 目录结构说明：

```ruby

➤ lt -L 2
 .
├──  api           // api 服务， 基于 proto/api/*.api 自动生成的代码
│  ├──  api.go
│  ├──  etc
│  └──  internal
├──  go.mod
├──  go.sum
├──  proto          // 微服务定义部分， 包含各种服务定义源文件（api, rpc, sql)
│  ├──  api         // api 定义
│  ├──  model       // 基于 sql， 生成的 model CRUD 代码
│  ├──  rpc         // rpc 服务定义
│  └──  sql         // raw sql， 用于生成 model
├──  readme.md
├──  rpc            // rpc 服务， 基于 proto/rpc/*.proto 自动生成的代码
│  ├──  etc
│  ├──  internal
│  ├──  main.go
│  ├──  pb
│  └──  service
└──  Taskfile.yml   // 服务启动脚本


```

## 示例服务：

- run:

```ruby

cd zeromall/

# go mod tidy:
task try:zero:tidy 

# run api:
task try:zero:run:api

# run rpc:
task try:zero:run:rpc

```

- api 验证：

```ruby

cd zeromall/

# api test:
task try:zero:api:test

```

### api 示例：

- new：

```ruby

goctl api new hello

```

### rpc 示例：

- new：

```ruby

goctl rpc new rpc
```

- 依赖服务发现机制， 默认基于 etcd 注册 / 发现服务。

## 参考：

> 文档：

- https://go-zero.dev/cn/docs/goctl/other
- https://go-zero.dev/cn/docs/goctl/api
- https://go-zero.dev/cn/docs/goctl/zrpc
- https://go-zero.dev/cn/docs/goctl/model

> 目录结构设计：

- https://go-zero.dev/cn/docs/advance/service-design

> 参考示例：

- https://go-zero.dev/cn/docs/eco/showcase
- https://github.com/Mikaelemmmm/go-zero-looklook
