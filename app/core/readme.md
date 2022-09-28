# core:

- 基础公共服务: 通用服务, 如 用户注册/登录/鉴权等, 短信, 邮件, 推送, 通知等

## 服务列表:

> [User](user): 用户中心

- [authn](user/authn): 注册/登录
- [authz](user/authz): rabc
- [identity](user/identity): 客户信息
- [system](user/system): 平台内部用户管理

> [Notification](notification): 消息推送中心

- [sms](notification/sms)
- [email](notification/email)
- [push](notification/push)

> [Security](security)

- [auth](security/auth): 平台内部 RPC 服务, 访问权限控制

> 基础中间件:

- [id-allocator](id-allocator): 发号器


