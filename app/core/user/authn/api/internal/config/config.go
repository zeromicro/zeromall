package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// require rpc client config
	SmsRpc zrpc.RpcClientConf // 注册码服务
}
