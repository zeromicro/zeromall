package server

import "mall/std/proto/server"

// 服务启动器:
func NewServerSelector() (server.Server, server.Server, server.Server, server.Server) {
	return &InnerServer{}, // gRpc server
		&OuterServer{}, // http server
		&JobServer{}, // job server
		&AdminServer{} // admin server
}
