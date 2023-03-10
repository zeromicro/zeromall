# client:

- 电商平台客户端:
- admin: 平台运营系统
    - web: PC web 客户端
    - app(iOS/Android): 移动端
    - desktop: 桌面版
- to C: 消费者客户端
    - web: PC web 客户端
    - app(iOS/Android): 移动端
    - desktop: 桌面版
- to B: 商户管理端
    - web: PC web 客户端
    - app(iOS/Android): 移动端
    - desktop: 桌面版

## 架构方案:

- flutter
- 当前 flutter 支持 iOS/Android/Windows/macos/Linux/web 全平台
- 目标是最大化复用前端代码
- 单个 flutter 项目, 多个 main 入口, 根据需求, 条件编译不同页面
- 比较复杂的权限系统

## admin:

- 平台运营系统

## to-C:

- 用户侧

## to-B:

- 商户侧




