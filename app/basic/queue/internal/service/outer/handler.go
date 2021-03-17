package outer

import (
	"context"
	"net/http"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/rest/httpx"

	"mall/app/basic/queue/proto/api"
	"mall/app/basic/queue/proto/rpc"
)

////////////////////////////////////////////////////////////////////////////////////////////////

// 示例 API:
func (m *Service) DemoHandler2(r *http.Request) (*api.Response, error) {
	var req api.HelloReq
	if err := httpx.Parse(r, &req); err != nil {
		return nil, err
	}
	return m.d.Hello.Demo(req)
}

// 消息队列写入:
func (m *Service) PublishMessage(r *http.Request) (*api.MessageResp, error) {
	var req api.MessageReq
	if err := httpx.Parse(r, &req); err != nil {
		return nil, err
	}
	return m.d.Hello.Publish(req)
}

// 解析块:
func (m *Service) ParseBlock(r *http.Request) (*api.Response, error) {
	m.d.Block.BlockParse(context.Background())
	return nil, nil
}

// rpc api:
func (m *Service) Graceful(r *http.Request) (resp *rpc.Response, err error) {
	resp, err = m.d.Hello.RPC.Graceful()
	log.Infof("rpc2 call api: resp=%+v, err=%v", resp, err)
	return
}

// rpc api:
func (m *Service) RestCall(r *http.Request) (resp *rpc.RestResp, err error) {
	// TODO: can not use proto req directly, need fix
	var req struct {
		From string `json:"from"`
	}
	if err := httpx.Parse(r, &req); err != nil {
		log.Infof("req parse err =%v", err)
		return nil, err
	}

	resp, err = m.d.Hello.RPC.RestCall(&rpc.RestReq{
		From: req.From,
	})
	log.Infof("RestCall: resp=%+v, err=%v", resp, err)
	return
}
