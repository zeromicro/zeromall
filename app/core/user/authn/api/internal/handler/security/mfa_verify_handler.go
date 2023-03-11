package security

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/authn/api/internal/logic/security"
	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"
)

func MfaVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MfaVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := security.NewMfaVerifyLogic(r.Context(), svcCtx)
		resp, err := l.MfaVerify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
