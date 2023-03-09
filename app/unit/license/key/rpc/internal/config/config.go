package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	// db url:
	DataSource string

	// cache:
	Cache cache.CacheConf

	// 服务注册/发现:
	Consul consul.Conf
}
