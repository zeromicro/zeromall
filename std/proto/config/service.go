package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

// 本服务定义: internal total
type ServerUnit struct {
	Inner RpcServerConf  // rpc server
	Outer RestServerConf // http api server
	Job   RestServerConf // job server
	Admin RestServerConf // admin api server
}

// 依赖服务定义:
type ClientUnit struct {
	// 本服务:
	Inner RpcClientConf

	// 其他依赖服务:
	// TODO: required other services, add here
}

////////////////////////////////////////////////////////////////////////////////////

// http restful server:
type RestServerConf struct {
	rest.RestConf
	Name    string
	Version string
}

// gRPC server:
type RpcServerConf struct {
	zrpc.RpcServerConf
	Version string
}

// gRPC client:
type RpcClientConf struct {
	zrpc.RpcClientConf
	Name    string
	Version string
}
