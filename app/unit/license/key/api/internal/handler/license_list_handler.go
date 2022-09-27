package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"license/key/api/internal/logic"
	"license/key/api/internal/svc"
	"license/key/api/internal/types"
)

func LicenseListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLicenseListLogic(r.Context(), svcCtx)
		resp, err := l.LicenseList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
