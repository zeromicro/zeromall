package router

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"mall/app/core/queue/internal/service/outer"
)

// api server routes:
func RegisterApiRoutes(engine *rest.Server, outer *outer.Service) {
	// register:
	engine.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodGet,
				Path:   "/hello/:name",
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return outer.DemoHandler2(r)
				}),
				//Handler: outer.Outer.DemoHandler(cfg, serverCtx),
			},
			{
				Method: http.MethodPost,
				Path:   "/queue/publish", // 队列写入
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return outer.PublishMessage(r)
				}),
			},
			{
				Method: http.MethodGet,
				Path:   "/graphql/do", // query block
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return outer.ParseBlock(r)
				}),
			},

			// rpc api test:
			{
				Method: http.MethodGet,
				Path:   "/api/graceful", // query block
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return outer.Graceful(r)
				}),
			},

			// rpc api test:
			{
				Method: http.MethodPost,
				Path:   "/api/rpc/test", // query block
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return outer.RestCall(r)
				}),
			},
		},
	)
}
