package router

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest"

	svc "mall/app/basic/demo/internal/dao"
	"mall/app/basic/demo/internal/service"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: service.DemoHandler(serverCtx),
			},
		},
	)
}
