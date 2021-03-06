package outer

import (
	"context"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"

	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/internal/domain/demo"
	"mall/app/basic/queue/proto/api"
	"mall/app/basic/queue/proto/config"
)

/*
	对外暴露的 API server: HTTP/WebSocket/GraphQL

*/
type Server struct {
	d *demo.Domain // 引入业务单元
}

func NewServer(cfg *config.Config, ctx *dao.ServiceContext) *Server {
	return &Server{
		// todo: need fix
		d: demo.NewDomain(cfg, context.TODO(), ctx),
	}
}

// 示例 API:
func (m *Server) DemoHandler(cfg *config.Config, ctx *dao.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		resp, err := m.d.Hello.Demo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

// 示例 API:
func (m *Server) DemoHandler2(r *http.Request) (*api.Response, error) {
	var req api.HelloReq
	if err := httpx.Parse(r, &req); err != nil {
		return nil, err
	}
	return m.d.Hello.Demo(req)
}

// 消息队列写入:
func (m *Server) PublishMessage(r *http.Request) (*api.MessageResp, error) {
	var req api.MessageReq
	if err := httpx.Parse(r, &req); err != nil {
		return nil, err
	}
	return m.d.Hello.Publish(req)
}

// 解析块:
func (m *Server) ParseBlock(r *http.Request) (*api.Response, error) {
	m.d.Graph.BlockParse(context.Background())
	return nil, nil
}
