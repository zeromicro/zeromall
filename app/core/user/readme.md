# user:

- union
- 注意服务顺序性:
    - identity -> authn -> authz
- 用户中心:
    - 账户系统
    - 登录鉴权
    - 用户角色
    - 用户权限

## 服务列表:

- ✅ [identity](identity): 用户基础信息
- ✅ [system](system) : 系统(平台)/管理员用户
- ✅ [developer](developer) : 开发者用户
- ✅ [authn](authn)
- ✅ [authz](authz)
- ✅ [oauth](oauth)


### [identity](identity):

- 客户管理
- 用户基础信息: 查询, 更新
- 用户设置更改

### [authn](authn):

- 外部用户, 身份识别

### [oauth](oauth):

- oauth 第三方授权登录:
- providers:
    - google
    - twitter
    - github
    - weixin
    - weibo

### [authz](authz):

- 内网服务, 访问权限控制
- rabc

### [system](system):

- 系统平台用户: adminUser, superUser, systemUser.
- 单独的用户系统, 隔离开.
- 避免出现安全问题.

