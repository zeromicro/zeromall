package register

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/authn/api/internal/logic/register"
	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"
)

func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterCheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := register.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
