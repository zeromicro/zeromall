package quick

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/authn/api/internal/logic/quick"
	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"
)

func SendSmsCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSmsCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := quick.NewSendSmsCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendSmsCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
