package rpc

import (
	"github.com/tal-tech/go-zero/zrpc"

	"mall/app/basic/queue/proto/config"
)

type Dao struct {
	Inner zrpc.Client // own rpc client
	Grace zrpc.Client // other rpc client
}

func New(cfg config.ClientUnit, isRpcServer bool) *Dao {
	d := &Dao{
		// TODO: add other rpc client
	}

	// check own rpc client:
	if !isRpcServer {
		d.Inner = zrpc.MustNewClient(
			cfg.Inner.RpcClientConf,
		)
	}
	return d
}

func (m *Dao) Close() {
	if m.Inner != nil {
		defer m.Inner.Conn().Close()
	}
	//defer m.Grace.Conn().Close()
}
