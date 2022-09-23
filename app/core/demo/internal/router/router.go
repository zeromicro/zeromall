package router

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"demo/internal/dao"
	"demo/internal/service"
	"demo/proto/config"
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
		},
	)
}
