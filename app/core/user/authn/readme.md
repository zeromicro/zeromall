# key:

- 授权码 / 注册码 生成服务

## 服务列表：

- 注册码生成
- 注册码分配
- 注册码管理

## 说明：

- 启动依赖的 docker 中间件:
    - mysql
    - redis
    - kafka
    - etcd

> 服务初始化:

- db 创建, 表生成

```ruby
cd zeromall/

# db 初始化:
task license:key:init:db

```

> 服务代码生成:

```ruby
cd learn-go/

# pb 新增api， 动态生成新的模板代码：
task try:zero:gen:pb

# 生成 model CRUD 代码


```

## 参考:

- https://www.cloudflare.com/zh-cn/learning/access-management/authn-vs-authz/