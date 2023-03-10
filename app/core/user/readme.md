# user:

- ✅ 注意服务顺序性:
    - member -> authn -> authz
- ✅ 用户中心:
    - 账户系统
    - 登录鉴权
    - 用户角色
    - 用户权限

## 服务列表:

- ✅ 服务的端口号 = 服务编号 + 服务类型编码

| 服务                     | 编号   | 说明      | 依赖服务             | 服务等级 |
|------------------------|------|---------|------------------|------|
| [member](member)       | 1010 | 用户基础信息  | 无                | L0   |
| [authn](authn)         | 1011 | 注册/登录   | [member](member) | L1   |
| [authz](authz)         | 1012 | 权限控制    |                  | L1   |
| [oauth](oauth)         | 1013 | 第三方授权登录 |                  | L1   |
| [system](system)       | 1014 | 系统用户    |                  | L1   |
| [developer](developer) | 1015 | 开发者用户   |                  | L2   |

### [member](member):

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

