package router

import (
	"net/http"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"

	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/internal/service"
	"mall/app/basic/queue/proto/config"
)

func RegisterHandlers(engine *rest.Server, cfg *config.Config, serverCtx *dao.ServiceContext) {
	// new:
	s := service.NewService(cfg, serverCtx)

	// register:
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: s.Outer.DemoHandler(cfg, serverCtx),
			},
			{
				Method: http.MethodGet,
				Path:   "/hello/:name",
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return s.Outer.DemoHandler2(r)
				}),
				//Handler: s.Outer.DemoHandler(cfg, serverCtx),
			},
			{
				Method: http.MethodPost,
				Path:   "/queue/publish", // 队列写入
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return s.Outer.PublishMessage(r)
				}),
			},
		},
	)
}

type ServiceFunc func(r *http.Request) (resp interface{}, err error)

func HandlerWrap(fn ServiceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("http req=%v, err=%v", r)

		// biz logic handle:
		resp, err := fn(r)
		log.Infof("http resp=%v, err=%v", resp, err)

		// resp:
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
