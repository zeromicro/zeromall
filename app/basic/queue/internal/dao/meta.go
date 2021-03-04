package dao

import (
	"mall/app/basic/queue/internal/dao/cache"
	"mall/app/basic/queue/internal/dao/db"
	"mall/app/basic/queue/internal/dao/mq"
	"mall/app/basic/queue/proto/config"
)

// 数据层资源收敛入口:
type MetaResource struct {
	//Async *async.Dao // async task handler
	DB    *db.Dao    // db layer
	Cache *cache.Dao // cache layer
	MQ    *mq.Dao    // mq layer
	//HTTP  *http.Dao  // http layer
	//RPC   *rpc.Dao   // rpc layer
}

func NewMetaResource(cfg *config.Config) *MetaResource {
	return &MetaResource{
		//Async: async.New(),
		DB:    db.New(cfg.DB),
		Cache: cache.New(cfg.Cache),
		MQ:    mq.NewDao(cfg.MQ),
		//HTTP:  http.New(cfg.HTTP),
		//RPC:   rpc.New(cfg.RPC),
	}
}

func (m *MetaResource) Close() {
	//m.Async.Close()
	m.DB.Close()
	m.Cache.Close()
	m.MQ.Close()
	m.DB.Close()
	//m.HTTP.Close()
	//m.RPC.Close()
}

/*

todo: need refactor
*/
type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
