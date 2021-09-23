

# 部署相关:


# 环境:


## local:

- 本地开发环境:
- 启动 docker 容器:
    - 自动启动: mysql/redis/etcd/rabbitmq

```bash
cd zeromall/
task local:init

# or:

cd zeromall/deploy/local
task init


```


- 停止中间件:

```bash

cd zeromall/
task local:clean

```


### 配置:

- kafka
- RabbitMQ: https://github.com/service-mesh/devops/blob/master/deploy/compose/local/common-mq-rabbitmq.yml



## staging:

- 预发布环境:


## production:

- 生产环境:



# ref:

- https://github.com/service-mesh/devops


