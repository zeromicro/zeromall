package outer

import (
	"context"
	"demo/proto/config"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"demo/internal/dao"
	"demo/internal/domain/demo"
	types "demo/proto/api"
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
		d: demo.NewScope(cfg, context.TODO(), ctx),
	}
}

// 示例 API:
func (m *Server) DemoHandler(cfg *config.Config, ctx *dao.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := demo.NewScope(cfg, r.Context(), ctx)

		resp, err := l.Hello.Demo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
