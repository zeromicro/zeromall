# user/identity:

- 账号系统, 支持各种第三方登录方式

## 功能:

- 用户基本的 CRUD.
- 用户设置更改:
    - 密码修改
    - 手机号修改
    - 邮箱修改
    - 用户基础信息修改
    - 其他设置

## 参考:

- https://github.com/casdoor/casdoor/blob/master/object/user.go
- https://github.com/casdoor/casdoor/blob/master/object/token.go

> 权限控制:

- casbin:
- https://casdoor.org/zh/docs/permission/overview

> 图片验证码:

- https://casdoor.org/zh/docs/provider/captcha/overview
- https://developers.google.com/recaptcha

> oauth:

- https://casdoor.org/zh/docs/provider/overview
- https://github.com/casdoor/casdoor/blob/master/idp/bilibili.go
    - oauth providers
- https://github.com/pennersr/django-allauth
    - https://django-allauth.readthedocs.io/en/latest/installation.html