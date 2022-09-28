package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"license/key/api/internal/config"
	"license/key/rpc/service"
)

type ServiceContext struct {
	Config config.Config

	// rpc clients:
	KeyRpc service.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// 注册 rpc client:
		KeyRpc: service.NewService(zrpc.MustNewClient(c.LicenseKeyRPC)),
	}
}
