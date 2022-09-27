// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"license/key/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/license/key",
				Handler: LicenseCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/license/key",
				Handler: LicenseGetHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/license/keys",
				Handler: LicenseListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/license/key/disable",
				Handler: LicenseDisableHandler(serverCtx),
			},
		},
	)
}
