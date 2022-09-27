package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"license/key/api/internal/logic"
	"license/key/api/internal/svc"
	"license/key/api/internal/types"
)

func LicenseGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LicenseGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLicenseGetLogic(r.Context(), svcCtx)
		resp, err := l.LicenseGet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
