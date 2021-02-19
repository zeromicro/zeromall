package outer

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"

	svc "mall/app/basic/demo/internal/dao"
	"mall/app/basic/demo/internal/domain/demo"
	types "mall/app/basic/demo/proto/api"
)

/*
	对外暴露的 API server: HTTP/WebSocket/GraphQL

*/
type Server struct {
}

func NewServer() *Server {
	return &Server{}
}


// 示例 API:
func (m *Server) DemoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := demo.NewScope(r.Context(), ctx)
		resp, err := l.Demo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}