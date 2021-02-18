package service

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	svc "mall/app/basic/demo/internal/dao"
	"mall/app/basic/demo/internal/domain/demo"
	types "mall/app/basic/demo/proto/api"
)

func DemoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := demo.NewDemoLogic(r.Context(), ctx)
		resp, err := l.Demo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
