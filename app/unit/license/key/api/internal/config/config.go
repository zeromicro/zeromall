package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// jwt auth: https://go-zero.dev/cn/docs/advance/jwt
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 注册码服务:
	LicenseKeyRPC zrpc.RpcClientConf
}
