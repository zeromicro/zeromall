# id-allocator:

- 分布式发号器， ID 生成器
- Distributed ID Generate Service
- Allocator is a distributed ID allocator backed by a KVstore.

## 使用场景：

- 数据库场景： 表唯一主键， 自增， 有序， 排序友好
- 交易场景： 订单号
- 队列场景：mq 消息流水号

## 参考：

- https://github.com/topics/distributed-id-generator

> 项目：

- https://github.com/rpcxio/did
- https://github.com/bwmarrin/snowflake
- https://github.com/edwingeng/wuid
- https://github.com/godruoyi/go-snowflake
- https://github.com/derry6/gleafd
    - 参考美团实现
- https://github.com/LeechanX/ekko-idgenerator
- https://github.com/EarlyZhao/id_generator
- https://github.com/cilium/cilium/blob/master/pkg/allocator/allocator.go#L88
    - 一个设计示例
    - Allocator is a distributed ID allocator backed by a KVstore.
- https://github.com/Meituan-Dianping/Leaf
    - 💖💖💖💖
    - 美团点评开源
    - Distributed ID Generate Service