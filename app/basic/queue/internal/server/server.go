package server

import (
	"github.com/better-go/pkg/x/go-zero/option"
)

// 服务启动器:
func NewServerSelector() (option.Server, option.Server, option.Server, option.Server) {
	return &InnerServer{}, // gRpc option
		&OuterServer{}, // http option
		&JobServer{}, // job option
		&AdminServer{} // admin option
}
