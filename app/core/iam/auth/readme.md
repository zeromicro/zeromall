# security/auth:

- 内网 rpc 服务, 互相访问权限控制

> 方案1:

- 主动管理服务互相访问的密钥对 (key, secret)
- 颁发调用证书, key

> 方案2:

- rpc 开启 TLS

## 服务:

- 服务证书颁发:
    - appKey, appSecret
    - appPermissions
- 权限验证
    - 性能瓶颈, 服务调用的单点
