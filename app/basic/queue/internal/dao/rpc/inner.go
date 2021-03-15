package rpc

import (
	"context"

	"mall/app/basic/queue/proto/rpc"
)

//
func (m *Dao) Graceful() (resp *rpc.Response, err error) {
	// rpc conn:
	conn := m.Inner.Conn()
	cli := rpc.NewQueueServiceClient(conn)

	return cli.Grace(context.Background(), &rpc.Request{From: "test"})
}

func (m *Dao) RestCall(req *rpc.RestReq) (resp *rpc.RestResp, err error) {
	// rpc conn:
	conn := m.Inner.Conn()
	cli := rpc.NewQueueServiceClient(conn)
	return cli.RestCall(context.Background(), req)
}

func (m *Dao) AdminCall(req *rpc.AdminReq) (resp *rpc.AdminResp, err error) {
	// rpc conn:
	conn := m.Inner.Conn()
	cli := rpc.NewQueueServiceClient(conn)
	return cli.AdminCall(context.Background(), req)
}
