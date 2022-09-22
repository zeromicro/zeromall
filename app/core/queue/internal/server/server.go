package server

import (
	server "github.com/better-go/pkg/x/go-zero/option"
)

// 服务启动器:
func NewServerSelector() (server.Server, server.Server, server.Server, server.Server) {
	return &InnerServer{}, // gRpc server
		&OuterServer{}, // http server
		&JobServer{}, // job server
		&AdminServer{} // admin server
}
