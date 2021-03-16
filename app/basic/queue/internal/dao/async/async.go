package async

import (
	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/sync/async"
)

const (
	defaultWorkerNum = 4
	defaultBufSize   = 1024
)

/*

Go Func() 方式: Goroutines 异步并发调度器
	- 按资源类型, 预分配 goroutines worker pool
	- 根据具体业务情况, 调整 workerNum 和 taskBufferSize

用途:
	- 禁止使用 go func(), 随意开 Goroutine, 导致写出糟糕难以维护的代码.
	- 统一使用 async.TaskDispatcher 利用 pool 的方式, 开 go func()

示例场景:
	- 在 一个 api 中, 需要批量刷缓存, 这时候可以异步处理刷缓存动作. 代码如下:

```
		m.CACHE.Dispatch(context.Background(), func(ctx context.Context) {
			doTaskfunc()
		})
```
	- 在 db 写入时, 用户登录场景, 有多张表更新,登录日志表, 用户登录积分表, 等表的更新
	- 这些表不 care 事务一致性和写失败. 可以异步处理, 加速请求返回. 写法类上例.

*/
type Dao struct {
	DB    *async.TaskDispatcher // 存储层 - 异步
	CACHE *async.TaskDispatcher // 缓存层 - 异步
	MQ    *async.TaskDispatcher // 队列层 - 异步
	HTTP  *async.TaskDispatcher // HTTP 层 - 异步
	RPC   *async.TaskDispatcher // RPC 层 - 异步
}

func New() *Dao {
	return &Dao{
		DB:    async.New("async-db-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
		CACHE: async.New("async-cache-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
		MQ:    async.New("async-mq-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
		HTTP:  async.New("async-http-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
		RPC:   async.New("async-rpc-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
	}
}

// need sync here:
func (m *Dao) Close() {
	if err := m.DB.Close(); err != nil {
		log.Errorf("async db task close error, err=%v", err)
	}
	if err := m.CACHE.Close(); err != nil {
		log.Errorf("async cache task close error, err=%v", err)
	}
	if err := m.MQ.Close(); err != nil {
		log.Errorf("async mq task close error, err=%v", err)
	}
	if err := m.HTTP.Close(); err != nil {
		log.Errorf("async http task close error, err=%v", err)
	}
	if err := m.RPC.Close(); err != nil {
		log.Errorf("async rpc task close error, err=%v", err)
	}
	return
}
